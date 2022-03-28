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
	"strings"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type setterModule struct {
	*pgs.ModuleBase
	initGoContext func(ctx pgs.Parameters) pgsgo.Context
	ctx           pgsgo.Context
}

func (m *setterModule) buildSetFieldsCase(buf *strings.Builder, imports importMap, tabCount uint, subs string, f pgs.Field) error {
	m.Push(f.Name().String())
	defer m.Pop()

	buildIndented(buf, tabCount, fmt.Sprintf(`case "%s":`, f.Name()))

	if f.InOneOf() {
		buildIndented(buf, tabCount+1, fmt.Sprintf(`_, srcOk := src.%s.(*%s)`,
			m.ctx.Name(f.OneOf()), m.ctx.OneofOption(f),
		))
	}

	goType := m.ctx.Type(f)

	fName := m.ctx.Name(f).String()
	if fName == "" {
		fName = goType.Value().String()
		if i := strings.LastIndex(fName, "."); i > 0 {
			fName = fName[i+1:]
		}
	}

	ft := f.Type()

	buildFinal := func(tabCount uint) error {
		fPath := fName
		if f.InOneOf() {
			fPath = m.ctx.Name(f.OneOf()).String()
		}

		srcCheck := `src != nil`
		if f.InOneOf() {
			srcCheck = `srcOk`
		}
		buildIndented(buf, tabCount, fmt.Sprintf(`if %s {
			dst.%s = src.%s
		} else {`,
			srcCheck,
			fPath, fPath,
		))

		if goType.IsPointer() {
			buildIndented(buf, tabCount, fmt.Sprintf(`	dst.%s = nil
}`,
				fPath,
			))
			return nil
		}

		if path := m.ctx.FieldTypeImportPath(f); path != "" && path != m.ctx.ImportPath(f.Message()) {
			if err := imports.Add(m.ctx.FieldTypePackageName(f).String(), path.String()); err != nil {
				return err
			}
		}

		if f.InOneOf() {
			buildIndented(buf, tabCount, fmt.Sprintf(`	dst.%s = nil
}`,
				fPath,
			))
			return nil
		}
		buildIndented(buf, tabCount, fmt.Sprintf(`	var zero %s
	dst.%s = zero
}`,
			goType,
			fPath,
		))
		return nil
	}

	if !ft.IsEmbed() || !ft.Embed().BuildTarget() {
		if err := imports.Add("fmt", "fmt"); err != nil {
			return err
		}
		buildIndented(buf, tabCount+1, fmt.Sprintf(`if len(%s) > 0 {
	return fmt.Errorf("'%s' has no subfields, but %%s were specified", %s)
}`,
			subs,
			f.Name(), subs,
		))
		return buildFinal(tabCount + 1)
	}

	buildIndented(buf, tabCount+1, fmt.Sprintf(`if len(%s) > 0 {
	var newDst, newSrc *%s`,
		subs,
		goType.Value(),
	))

	switch {
	case f.InOneOf():
		fPath := fmt.Sprintf("%s.(*%s).%s", m.ctx.Name(f.OneOf()), m.ctx.OneofOption(f), fName)
		buildIndented(buf, tabCount+2, fmt.Sprintf(`if srcOk {
	newSrc = src.%s
}
_, dstOk := dst.%s.(*%s)
if dstOk {
	newDst = dst.%s
} else if srcOk {
	newDst = &%s{}
	dst.%s = &%s{%s: newDst}
} else {
	dst.%s = nil
	continue
}`,
			fPath,
			m.ctx.Name(f.OneOf()), m.ctx.OneofOption(f),
			fPath,
			goType.Value(),
			m.ctx.Name(f.OneOf()), m.ctx.OneofOption(f), fName,
			m.ctx.Name(f.OneOf()),
		))

	case goType.IsPointer():
		buildIndented(buf, tabCount+2, fmt.Sprintf(`if (src == nil || src.%s == nil) && dst.%s == nil {
	continue
}
if src != nil {
	newSrc = src.%s
}
if dst.%s != nil {
	newDst = dst.%s
} else {
	newDst = &%s{}
	dst.%s = newDst
}`,
			fName, fName,
			fName,
			fName,
			fName,
			goType.Value(),
			fName,
		))

	default:
		buildIndented(buf, tabCount+2, fmt.Sprintf(`if src != nil {
	newSrc = &src.%s
}
newDst = &dst.%s`,
			fName,
			fName,
		))
	}

	buildIndented(buf, tabCount+1, fmt.Sprintf(`	if err := newDst.SetFields(newSrc, %s...); err != nil {
		return err
	}
} else {`, subs))
	if err := buildFinal(tabCount + 2); err != nil {
		return err
	}
	buildIndented(buf, tabCount+1, `}`)
	return nil
}

func (m *setterModule) buildSetFields(buf *strings.Builder, imports importMap, msg pgs.Message) error {
	m.Push(msg.FullyQualifiedName())
	defer m.Pop()

	if err := imports.Add("fmt", "fmt"); err != nil {
		return err
	}

	mType := m.ctx.Name(msg)
	if len(msg.Fields()) == 0 {
		fmt.Fprintf(buf, `
func (dst *%s) SetFields(src *%s, paths ...string) error {
	if len(paths) != 0 {
		return fmt.Errorf("message %s has no fields, but paths %%s were specified", paths)
	}
	if src != nil {
		*dst = *src
	}
	return nil
}`,
			mType, mType,
			mType,
		)
		return nil
	}

	fmt.Fprintf(buf, `
func (dst *%s) SetFields(src *%s, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
`,
		mType, mType,
	)

	for _, f := range msg.NonOneOfFields() {
		if err := m.buildSetFieldsCase(buf, imports, 2, "subs", f); err != nil {
			return err
		}
	}

	for _, o := range msg.OneOfs() {
		fmt.Fprintln(buf)

		goName := m.ctx.Name(o)

		buildIndented(buf, 2, fmt.Sprintf(`case "%s":
	if len(subs) == 0 && src == nil {
		dst.%s = nil
		continue
	} else if len(subs) == 0 {
		dst.%s = src.%s
		continue
	}`,
			o.Name(),
			goName,
			goName, goName,
		))

		fmt.Fprintln(buf)

		buildIndented(buf, 3, `subPathMap := _processPaths(subs)
if len(subPathMap) > 1 {
	return fmt.Errorf("more than one field specified for oneof field '%s'", name)
}
for oneofName, oneofSubs := range subPathMap {
	switch oneofName {`)

		for _, f := range o.Fields() {
			if err := m.buildSetFieldsCase(buf, imports, 4, "oneofSubs", f); err != nil {
				return err
			}
		}

		fmt.Fprintln(buf)

		buildIndented(buf, 3, `	default:
		return fmt.Errorf("invalid oneof field: '%s.%s'", name, oneofName)
	}
}`)
	}

	fmt.Fprintf(buf, `
		default:
			return fmt.Errorf("invalid field: '%%s'", name)
		}
	}
	return nil
}`,
	)
	return nil
}

func (m *setterModule) Name() string { return "setter" }

func (m *setterModule) InitContext(ctx pgs.BuildContext) {
	m.ModuleBase.InitContext(ctx)
	m.ctx = m.initGoContext(ctx.Parameters())
}

func (m *setterModule) Execute(files map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	for _, f := range files {
		m.Push(f.Name().String())

		if len(f.Messages()) == 0 {
			m.Pop()
			continue
		}

		imports := importMap{}
		buf := &strings.Builder{}
		for _, msg := range f.AllMessages() {
			var mBufs []*strings.Builder

			mBuf := &strings.Builder{}
			if err := m.buildSetFields(mBuf, imports, msg); err != nil {
				m.AddError(fmt.Errorf("failed to build SetFields for %s: %s", msg.Name(), err).Error())
				return m.Artifacts()
			}
			mBufs = append(mBufs, mBuf)

			for _, mBuf := range mBufs {
				fmt.Fprintf(buf, `
%s`,
					mBuf.String())
			}
		}

		m.AddGeneratorTemplateFile(m.ctx.OutputPath(f).SetExt(".setters.fm.go").String(), template.Must(template.New("setters").Parse(`package {{ .Package }}{{ .ImportString }}

{{ .Content }}`)), struct {
			Package      pgs.Name
			ImportString string
			Content      string
		}{
			Package:      m.ctx.PackageName(f),
			ImportString: imports.GoString(),
			Content:      buf.String(),
		})
		m.Pop()
	}
	return m.Artifacts()
}

// Setter generates SetFields method on messages.
// It depends on code generated by Pather module.
func Setter(initGoContext func(ctx pgs.Parameters) pgsgo.Context) pgs.Module {
	return &setterModule{
		ModuleBase:    &pgs.ModuleBase{},
		initGoContext: initGoContext,
	}
}
