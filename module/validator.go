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

type validatorModule struct {
	*pgs.ModuleBase
	initGoContext func(ctx pgs.Parameters) pgsgo.Context
	ctx           pgsgo.Context
}

func (m *validatorModule) buildValidateFieldsCase(buf *strings.Builder, imports importMap, tabCount uint, subs string, f pgs.Field) error {
	m.Push(f.Name().String())
	defer m.Pop()

	buildIndented(buf, tabCount, fmt.Sprintf(`case "%s":`, f.Name()))

	goType := m.ctx.Type(f)

	fieldPath := "msg"
	if f.InOneOf() {
		fieldPath = fmt.Sprintf("%s.Get%s()", fieldPath, m.ctx.Name(f))
	} else {
		name := m.ctx.Name(f).String()
		if name == "" {
			name = goType.Value().String()
			if i := strings.LastIndex(name, "."); i > 0 {
				name = name[i+1:]
			}
		}
		fieldPath = fmt.Sprintf("%s.%s", fieldPath, name)
	}

	ft := f.Type()

	// TODO: Validate Embed field (https://github.com/TheThingsIndustries/protoc-gen-fieldmask/issues/21)

	if ft.IsEmbed() && ft.Embed().BuildTarget() {
		buildIndented(buf, tabCount+1, fmt.Sprintf(`if err := %s.ValidateFields(%s...); err != nil {
			return err
		}`,
			fieldPath, subs,
		))
		return nil
	}

	if err := imports.Add("fmt", "fmt"); err != nil {
		return err
	}

	buildIndented(buf, tabCount+1, fmt.Sprintf(`if len(%s) > 0 {
	return fmt.Errorf("'%s' has no subfields, but %%s were specified", %s)
}
_ = %s // TODO: Validate (https://github.com/TheThingsIndustries/protoc-gen-fieldmask/issues/21)
if false {
	return fmt.Errorf("'%s' is invalid")
}`,
		subs,
		f.Name(), subs,
		fieldPath,
		f.Name(),
	))

	return nil
}

func (m *validatorModule) buildValidateFields(buf *strings.Builder, imports importMap, msg pgs.Message) error {
	m.Push(msg.FullyQualifiedName())
	defer m.Pop()

	if err := imports.Add("fmt", "fmt"); err != nil {
		return err
	}

	mType := m.ctx.Name(msg)
	if len(msg.Fields()) == 0 {
		fmt.Fprintf(buf, `
func (msg *%s) ValidateFields(paths ...string) error {
	if len(paths) != 0 {
		return fmt.Errorf("message %s has no fields, but paths %%s were specified", paths)
	}
	return nil
}`,
			mType,
			mType,
		)
		return nil
	}

	fmt.Fprintf(buf, `
func (msg *%s) ValidateFields(paths ...string) error {
	if msg == nil {
		return nil
	}
	if len(paths) == 0 {
		paths = %sFieldPathsNested
	}
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
`,
		mType,
		mType,
	)

	for _, f := range msg.NonOneOfFields() {
		if err := m.buildValidateFieldsCase(buf, imports, 2, "subs", f); err != nil {
			return err
		}
	}

	for _, o := range msg.OneOfs() {
		fs := o.Fields()

		subPaths := make([]string, 0, len(fs))
		for _, f := range fs {
			subPaths = append(subPaths, fmt.Sprintf("%s.%s", o.Name(), f.Name()))
		}
		sort.Strings(subPaths)

		buildIndented(buf, 2, fmt.Sprintf(`case "%s":
	if len(subs) == 0 {
		subs = []string{
			%s
		}
	}`,
			o.Name(),
			`"`+strings.Join(subPaths, `",
			"`)+`",`,
		))

		fmt.Fprintln(buf)

		buildIndented(buf, 3, `subPathMap := _processPaths(subs)
for oneofName, oneofSubs := range subPathMap {
	switch oneofName {`)

		for _, f := range fs {
			if err := m.buildValidateFieldsCase(buf, imports, 4, "oneofSubs", f); err != nil {
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

func (m *validatorModule) Name() string { return "validator" }

func (m *validatorModule) InitContext(ctx pgs.BuildContext) {
	m.ModuleBase.InitContext(ctx)
	m.ctx = m.initGoContext(ctx.Parameters())
}

func (m *validatorModule) Execute(files map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
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
			if err := m.buildValidateFields(mBuf, imports, msg); err != nil {
				m.AddError(fmt.Errorf("failed to build ValidateFields for %s: %s", msg.Name(), err).Error())
				return m.Artifacts()
			}
			mBufs = append(mBufs, mBuf)

			for _, mBuf := range mBufs {
				fmt.Fprintf(buf, `
%s`,
					mBuf.String())
			}
		}

		m.AddGeneratorTemplateFile(m.ctx.OutputPath(f).SetExt(".validators.fm.go").String(), template.Must(template.New("validators").Parse(`package {{ .Package }}{{ .ImportString }}

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

// Validator generates ValidateFields method on messages.
// It depends on code generated by Pather module.
func Validator(initGoContext func(ctx pgs.Parameters) pgsgo.Context) pgs.Module {
	return &validatorModule{
		ModuleBase:    &pgs.ModuleBase{},
		initGoContext: initGoContext,
	}
}
