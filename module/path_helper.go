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

type messagePaths struct {
	NestedPaths   []string
	TopLevelPaths []string
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

func (m *pathHelperModule) renderMessageFieldPaths(buf *strings.Builder, paths map[string]messagePaths) {
	tmpl := `
var {{ .MessageType }}FieldPathsNested = []string{
	{{- range $path := .MessagePaths.NestedPaths }}
	"{{ $path }}",
	{{- end }}
}
{{printf "\n"}}
var {{ .MessageType }}FieldPathsTopLevel = []string{
	{{- range $path := .MessagePaths.TopLevelPaths }}
	"{{ $path }}",
	{{- end }}
}
{{printf "\n"}}
`
	// Sort the message types to ensure deterministic output.
	keys := make([]string, 0, len(paths))
	for k := range paths {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, t := range keys {
		p := paths[t]

		data := struct {
			MessageType  string
			MessagePaths messagePaths
		}{
			MessageType:  t,
			MessagePaths: p,
		}

		err := template.Must(template.New("paths").Parse(tmpl)).Execute(buf, data)
		if err != nil {
			m.AddError(fmt.Errorf("failed to execute paths template: %w", err).Error())
			return
		}
	}
}

func (m *pathHelperModule) renderPathUtil() string {
	content := `
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
}`

	return content
}

func (m *pathHelperModule) generatePaths(files map[string]pgs.File) error {
	dirs := map[pgs.FilePath]pgs.Name{}

	for _, f := range files {
		m.Push(f.Name().String())

		if len(f.Messages()) == 0 {
			m.Pop()
			continue
		}

		paths := map[string]messagePaths{}
		for _, msg := range f.AllMessages() {
			messageType := m.ctx.Name(msg).String()

			nestedPaths, topLevelPaths, err := m.buildPaths(msg)
			if err != nil {
				return fmt.Errorf("failed to build paths for %s: %s", msg.Name(), err)
			}

			paths[messageType] = messagePaths{
				NestedPaths:   nestedPaths,
				TopLevelPaths: topLevelPaths,
			}
		}

		content := &strings.Builder{}
		m.renderMessageFieldPaths(content, paths)
		m.writeFile(
			m.ctx.OutputPath(f).SetExt(".paths.fm.go").String(),
			m.ctx.PackageName(f).String(),
			content.String(),
		)

		dirs[m.ctx.OutputPath(f).Dir()] = m.ctx.PackageName(f)
		m.Pop()
	}

	for dir, pkg := range dirs {
		content := m.renderPathUtil()
		m.writeFile(
			dir.Push(pkg.LowerCamelCase().String()).SetExt(".pb.util.fm.go").String(),
			pkg.String(),
			content,
		)
	}

	return nil
}

func (m *pathHelperModule) renderRPCFieldMaskDefinitions(buf *strings.Builder, rpcFieldMaskPaths map[string][]string) {
	const t = `
	var {{ .ID }}AllowedFieldMaskPaths = []string{
		{{- range $path := .Paths }}
		"{{ $path }}",
		{{- end }}
	}
	{{printf "\n"}}
`

	// Sort the methodIDs and paths to ensure deterministic output.
	methodIDs := make([]string, 0, len(rpcFieldMaskPaths))
	for id := range rpcFieldMaskPaths {
		methodIDs = append(methodIDs, id)
	}
	sort.Strings(methodIDs)

	for methodID, paths := range rpcFieldMaskPaths {
		sort.Strings(paths)
		rpcFieldMaskPaths[methodID] = paths
	}

	for _, id := range methodIDs {
		rpcFieldMaskPaths[id] = rpcFieldMaskPaths[id]
		data := struct {
			ID    string
			Paths []string
		}{
			ID:    id,
			Paths: rpcFieldMaskPaths[id],
		}

		err := template.Must(template.New("AllowedRPCFieldMaskPaths").Parse(t)).Execute(buf, data)
		if err != nil {
			m.AddError(fmt.Errorf("failed to execute RPCFieldMaskPaths template: %w", err).Error())
			return
		}
	}
}

func (m *pathHelperModule) generateRPCFieldMaskPaths(files map[string]pgs.File) error {
	rpcFieldMaskPaths := map[string][]string{}
	for _, f := range files {
		for _, svc := range f.Services() {
			for _, method := range svc.Methods() {
				options := method.Descriptor().GetOptions()
				if !proto.HasExtension(options, annotations.E_Allow) {
					continue
				}

				methodID := fmt.Sprintf("%s%s", svc.Name().String(), method.Name().String())
				ext, ok := proto.GetExtension(options, annotations.E_Allow).(*annotations.RPCMaskAllowOption)
				if !ok {
					return fmt.Errorf("failed to get method extension")
				}

				if ext.GetFieldMask() == nil {
					// This should never happen as the extension is set.
					return fmt.Errorf("the extension is set, but the field mask is nil")
				}

				if ext.GetFieldMask().GetPaths() == nil {
					// This should never happen as the field_mask is required.
					return fmt.Errorf("field mask paths are nil")
				}

				rpcFieldMaskPaths[methodID] = ext.GetFieldMask().GetPaths()
			}
		}

		if len(rpcFieldMaskPaths) == 0 {
			continue
		}

		content := &strings.Builder{}
		m.renderRPCFieldMaskDefinitions(content, rpcFieldMaskPaths)
		m.writeFile(
			m.ctx.OutputPath(f).SetExt(".allowed.fm.go").String(),
			m.ctx.PackageName(f).String(),
			content.String(),
		)
	}

	return nil
}

func (m *pathHelperModule) writeFile(filepath, pkg, content string) {
	tmpl := `
package {{ .Package }}
{{ printf "\n" }}
{{ .Content }}
	`

	data := struct {
		Package string
		Content string
	}{
		Package: pkg,
		Content: content,
	}

	m.AddGeneratorTemplateFile(filepath, template.Must(template.New("file_template").Parse(tmpl)), data)
}

func (m *pathHelperModule) Name() string { return "paths" }

func (m *pathHelperModule) InitContext(ctx pgs.BuildContext) {
	m.ModuleBase.InitContext(ctx)
	m.ctx = m.initGoContext(ctx.Parameters())
}

func (m *pathHelperModule) Execute(files map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	if err := m.generatePaths(files); err != nil {
		m.AddError(err.Error())
		return m.Artifacts()
	}

	if err := m.generateRPCFieldMaskPaths(files); err != nil {
		m.AddError(err.Error())
		return m.Artifacts()
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
