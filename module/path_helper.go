// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package module

import (
	"fmt"
	"sort"
	"strings"
	"text/template"

	"github.com/TheThingsIndustries/protoc-gen-fieldmask/annotations"
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	"google.golang.org/protobuf/proto"
)

type pathHelperModule struct {
	*pgs.ModuleBase
	initGoContext func(ctx pgs.Parameters) pgsgo.Context
	ctx           pgsgo.Context
}

func (m *pathHelperModule) appendPaths(ctx pgsgo.Context, paths []string, prefix string, msg pgs.Message, seen map[string]struct{}) ([]string, error) {
	if seen == nil {
		seen = map[string]struct{}{}
	}

	for _, f := range msg.Fields() {
		fp := f.Name().String()
		if f.InOneOf() {
			fp = fmt.Sprintf("%s.%s", f.OneOf().Name(), fp)
		}
		if prefix != "" {
			fp = fmt.Sprintf("%s.%s", prefix, fp)
		}

		fqn := f.FullyQualifiedName()
		if _, ok := seen[fqn]; ok {
			m.Logf("Field '%s' defined at %s:%d is recursive, stopping traversal", fp, f.File().Name(), f.SourceCodeInfo().Location().Span[0]+1)
			return paths, nil
		}
		seen[fqn] = struct{}{}

		paths = append(paths, fp)

		if f.Type().IsRepeated() || !f.Type().IsEmbed() {
			delete(seen, fqn)
			continue
		}

		sub := f.Type().Embed()
		if !sub.BuildTarget() || len(sub.Fields()) == 0 {
			delete(seen, fqn)
			continue
		}

		var err error
		paths, err = m.appendPaths(ctx, paths, fp, sub, seen)
		if err != nil {
			return nil, err
		}
		delete(seen, fqn)
	}

	for _, o := range msg.OneOfs() {
		fp := o.Name().String()
		if prefix != "" {
			fp = fmt.Sprintf("%s.%s", prefix, fp)
		}
		paths = append(paths, fp)
	}
	return paths, nil
}

func (m *pathHelperModule) buildPaths(msg pgs.Message) ([]string, []string, error) {
	m.Push(msg.FullyQualifiedName())
	defer m.Pop()

	if len(msg.Fields()) == 0 {
		return nil, nil, nil
	}

	nestedPaths, err := m.appendPaths(m.ctx, make([]string, 0, len(msg.Fields())), "", msg, nil)
	if err != nil {
		return nil, nil, err
	}
	sort.Strings(nestedPaths)

	topLevelPaths := make([]string, 0, len(nestedPaths))
	for _, p := range nestedPaths {
		if strings.LastIndex(p, ".") > 0 {
			continue
		}
		topLevelPaths = append(topLevelPaths, p)
	}
	sort.Strings(topLevelPaths)

	return nestedPaths, topLevelPaths, nil
}

func (m *pathHelperModule) writePaths(buf *strings.Builder, msg pgs.Message, nestedPaths, topLevelPaths []string) {
	mType := m.ctx.Name(msg)

	data := struct {
		Type          pgs.Name
		NestedPaths   []string
		TopLevelPaths []string
	}{
		Type:          mType,
		NestedPaths:   nestedPaths,
		TopLevelPaths: topLevelPaths,
	}

	if len(nestedPaths) == 0 && len(topLevelPaths) == 0 {
		emptyPathsTemplate := `
var {{ .Type }}FieldPathsNested []string
var {{ .Type }}FieldPathsTopLevel []string
		`

		err := template.Must(template.New("paths").Parse(emptyPathsTemplate)).Execute(buf, data)
		if err != nil {
			m.AddError(fmt.Errorf("failed to execute paths template: %w", err).Error())
			return
		}

		return
	}

	pathsTemplate := `
var {{ .Type }}FieldPathsNested = []string{
	{{- range $path := .NestedPaths }}
	"{{ $path }}",
	{{- end }}
}

var {{ .Type }}FieldPathsTopLevel = []string{
	{{- range $path := .TopLevelPaths }}
	"{{ $path }}",
	{{- end }}
}
`

	err := template.Must(template.New("paths").Parse(pathsTemplate)).Execute(buf, data)
	if err != nil {
		m.AddError(fmt.Errorf("failed to execute paths template: %w", err).Error())
		return
	}
}

func (m *pathHelperModule) writeRPCFieldMaskPaths(buf *strings.Builder, rpcFieldMaskPaths map[string]RPCFieldMaskPathValue) {
	const structTemplate = `
type RPCFieldMaskPathValue struct {
    All     []string
    Allowed []string
    Set     bool
}`
	fmt.Fprintf(buf, structTemplate)

	// Empty RPCFieldMaskPaths.
	if len(rpcFieldMaskPaths) == 0 {
		fmt.Fprintf(buf, `
// RPCFieldMaskPaths lists the field mask paths for each RPC in this API.
var RPCFieldMaskPaths = map[string]RPCFieldMaskPathValue{}
`)
		return
	}

	const mapTemplate = `

// RPCFieldMaskPaths lists the field mask paths for each RPC in this API.
var RPCFieldMaskPaths = map[string]RPCFieldMaskPathValue{
{{- range $key, $value := . }}
    "{{ $key }}": {
        All:     {{ $value.All }},
        Allowed: []string{
            {{- range $path := $value.Allowed }}
            "{{ $path }}",
            {{- end }}
        },
        Set:     {{ $value.Set }},
    },
{{- end }}
}
`

	keys := make([]string, 0, len(rpcFieldMaskPaths))
	for k := range rpcFieldMaskPaths {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	err := template.Must(template.New("RPCFieldMaskPathsMap").Parse(mapTemplate)).Execute(buf, rpcFieldMaskPaths)
	if err != nil {
		m.AddError(fmt.Errorf("failed to execute RPCFieldMaskPathsMap template: %w", err).Error())
		return
	}
}

func (m *pathHelperModule) Name() string { return "paths" }

func (m *pathHelperModule) InitContext(ctx pgs.BuildContext) {
	m.ModuleBase.InitContext(ctx)
	m.ctx = m.initGoContext(ctx.Parameters())
}

func (m *pathHelperModule) Execute(files map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	dirs := map[pgs.FilePath]pgs.Name{}

	for _, f := range files {
		m.Push(f.Name().String())

		if len(f.Messages()) == 0 {
			m.Pop()
			continue
		}

		buf := &strings.Builder{}
		for _, msg := range f.AllMessages() {
			var mBufs []*strings.Builder

			nestedPaths, topLevelPaths, err := m.buildPaths(msg)
			if err != nil {
				m.AddError(fmt.Errorf("failed to build paths for %s: %s", msg.Name(), err).Error())
				return m.Artifacts()
			}

			mBuf := &strings.Builder{}
			m.writePaths(buf, msg, nestedPaths, topLevelPaths)
			mBufs = append(mBufs, mBuf)

			for _, mBuf := range mBufs {
				fmt.Fprintf(buf, `
%s`,
					mBuf.String())
			}
		}

		dirs[m.ctx.OutputPath(f).Dir()] = m.ctx.PackageName(f)

		m.AddGeneratorTemplateFile(m.ctx.OutputPath(f).SetExt(".paths.fm.go").String(),
			template.Must(template.New("paths").Parse(`package {{ .Package }}

{{ .Content }}`)), struct {
				Package pgs.Name
				Content string
			}{
				Package: m.ctx.PackageName(f),
				Content: buf.String(),
			})
		m.Pop()
	}

	for dir, pkg := range dirs {
		baseName := pkg.LowerCamelCase().String()
		m.AddGeneratorTemplateFile(dir.Push(baseName).SetExt(".pb.util.fm.go").String(),
			template.Must(template.New("util").Parse(`package {{ .Package }}

import (
	"strings"
)

// _processPaths returns paths as a pathMap.
func _processPaths(paths []string) map[string][]string {
	if len(paths) == 0 {
		return nil
	}
	pathMap := make(map[string][]string, len(paths))
	for _, p := range paths {
		if !strings.Contains(p, ".") {
			pathMap[p] = nil
			continue
		}
		parts := strings.SplitN(p, ".", 2)
		h, t := parts[0], parts[1]
		if val, ok := pathMap[h]; ok {
			if val == nil {
				continue
			}
			pathMap[h] = append(pathMap[h], t)
		} else {
			pathMap[h] = []string{t}
		}
	}

	return pathMap
}`)), struct {
				Package pgs.Name
			}{
				Package: pkg,
			})
	}

	// Iterate over all services and their methods to generate RPCFieldMaskPaths
	rpcFieldMaskPaths := map[string]RPCFieldMaskPathValue{}
	for _, f := range files {
		for _, svc := range f.Services() {
			for _, method := range svc.Methods() {
				packageName := svc.Package().ProtoName().String()
				serviceName := svc.Name().String()
				methodName := method.Name().String()
				rpcMethodIdentifier := fmt.Sprintf("/%s.%s/%s", packageName, serviceName, methodName)

				options := method.Descriptor().GetOptions()
				if proto.HasExtension(options, annotations.E_Method) {
					ext, ok := proto.GetExtension(options, annotations.E_Method).(*annotations.MethodOptions)
					if !ok {
						m.AddError("failed to get service extension")
						return m.Artifacts()
					}

					rpcmask := ext.GetRpcmask()
					if rpcmask == nil {
						continue
					}
					rpcFieldMaskPaths[rpcMethodIdentifier] = RPCFieldMaskPathValue{
						All:     fmt.Sprintf("%sFieldPathsNested", rpcmask.GetMessageName()),
						Allowed: rpcmask.FieldMask.GetPaths(),
						Set:     rpcmask.GetSet(),
					}
				}
			}
		}
	}

	// Generate the RPCFieldMaskPaths for every directory and pkg
	for dir, pkg := range dirs {
		m.AddGeneratorTemplateFile(dir.Push("field_mask_validation").SetExt(".go").String(),
			template.Must(template.New("field_mask_validation").Parse(`package {{ .Package }}

{{ .Content }}`)), struct {
				Package pgs.Name
				Content string
			}{
				Package: pkg,
				Content: func() string {
					buf := &strings.Builder{}
					m.writeRPCFieldMaskPaths(buf, rpcFieldMaskPaths)
					return buf.String()
				}(),
			})
	}

	return m.Artifacts()
}

// PathHelper generates various fieldmask-related utilities and variables.
func PathHelper(initGoContext func(ctx pgs.Parameters) pgsgo.Context) pgs.Module {
	return &pathHelperModule{
		ModuleBase:    &pgs.ModuleBase{},
		initGoContext: initGoContext,
	}
}
