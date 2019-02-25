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

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
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

func (m *pathHelperModule) buildPaths(buf *strings.Builder, msg pgs.Message) error {
	m.Push(msg.FullyQualifiedName())
	defer m.Pop()

	mType := m.ctx.Name(msg)
	if len(msg.Fields()) == 0 {
		fmt.Fprintf(buf, `var %sFieldPathsNested []string
var %sFieldPathsTopLevel []string`,
			mType,
			mType,
		)
		return nil
	}

	nestedPaths, err := m.appendPaths(m.ctx, make([]string, 0, len(msg.Fields())), "", msg, nil)
	if err != nil {
		return err
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

	fmt.Fprintf(buf, `var %sFieldPathsNested = []string{
	%s
}

var %sFieldPathsTopLevel = []string{
	%s
}`,
		mType, `"`+strings.Join(nestedPaths, `",
	"`)+`",`,
		mType, `"`+strings.Join(topLevelPaths, `",
	"`)+`",`,
	)
	return nil
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

			mBuf := &strings.Builder{}
			if err := m.buildPaths(mBuf, msg); err != nil {
				m.AddError(fmt.Errorf("failed to build paths for %s: %s", msg.Name(), err).Error())
				return m.Artifacts()
			}
			mBufs = append(mBufs, mBuf)

			for _, mBuf := range mBufs {
				fmt.Fprintf(buf, `
%s`,
					mBuf.String())
			}
		}

		dirs[m.ctx.OutputPath(f).Dir()] = m.ctx.PackageName(f)

		m.AddGeneratorTemplateFile(m.ctx.OutputPath(f).SetExt(".paths.fm.go").String(), template.Must(template.New("paths").Parse(`package {{ .Package }}

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
		m.AddGeneratorTemplateFile(dir.Push(dir.Base()).SetExt(".pb.util.fm.go").String(), template.Must(template.New("util").Parse(`package {{ .Package }}

import (
	"sort"
	"strings"
)

// _processPaths returns paths as a pathMap.
func _processPaths(paths []string) map[string][]string {
	sort.Strings(paths)

	topLevel := make(map[string]struct{}, len(paths))
	_pathMap := make(map[string]map[string]struct{}, len(paths))
	for _, p := range paths {
		if !strings.Contains(p, ".") {
			topLevel[p] = struct{}{}
			continue
		}
		parts := strings.SplitN(p, ".", 2)
		h, t := parts[0], parts[1]
		if _pathMap[h] == nil {
			_pathMap[h] = map[string]struct{}{t: {}}
		} else {
			_pathMap[h][t] = struct{}{}
		}
	}

	for f := range topLevel {
		_pathMap[f] = nil
	}

	pathMap := make(map[string][]string, len(_pathMap))
	for top, subs := range _pathMap {
		pathMap[top] = make([]string, 0, len(subs))
		for sub := range subs {
			pathMap[top] = append(pathMap[top], sub)
		}
	}
	return pathMap
}`)), struct {
			Package pgs.Name
		}{
			Package: pkg,
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
