// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
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

package main

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"sort"
	"strings"

	_ "github.com/TheThingsIndustries/protoc-gen-fieldmask/internal/extensions"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pseudomuto/protokit"
)

const FileHeader = `// Code generated by protoc-gen-fieldmask. DO NOT EDIT.`

const (
	protoAnyType       = ".google.protobuf.Any"
	protoDurationType  = ".google.protobuf.Duration"
	protoFieldMaskType = ".google.protobuf.FieldMask"
	protoStructType    = ".google.protobuf.Struct"
	protoTimestampType = ".google.protobuf.Timestamp"
)

type unknownTypeError string

func (e unknownTypeError) Error() string {
	return fmt.Sprintf("message of proto type '%s' is unknown", string(e))
}

type unsupportedTypeError string

func (e unsupportedTypeError) Error() string {
	return fmt.Sprintf("fields of proto type '%s' are unsupported", string(e))
}

func walkMessage(md *protokit.Descriptor, f func(md *protokit.Descriptor) error) error {
	if err := f(md); err != nil {
		return err
	}
	for _, smd := range md.GetMessages() {
		if err := walkMessage(smd, f); err != nil {
			return err
		}
	}
	return nil
}

func registerMessages(mdMap map[string]*protokit.Descriptor, mds ...*protokit.Descriptor) error {
	for _, md := range mds {
		if err := walkMessage(md, func(md *protokit.Descriptor) error {
			k := fmt.Sprintf(".%s", md.GetFullName())
			if _, ok := mdMap[k]; ok {
				return fmt.Errorf("message name clash at `%s`", k)
			}
			mdMap[k] = md
			return nil
		}); err != nil {
			return err
		}
	}
	return nil
}

type importMap map[string]string

func (m importMap) Add(name, pkg string) error {
	if name == "" && pkg == "" {
		return nil
	}
	if name == "" {
		return fmt.Errorf("import name must be specified")
	}
	if pkg == "" {
		return fmt.Errorf("package path must be specified")
	}

	if v, ok := m[name]; ok && v != pkg {
		return fmt.Errorf("import name clash at `%s`. Imported `%s` and `%s`", name, pkg, v)
	}
	m[name] = pkg
	return nil
}

func (m importMap) AddMultiple(pairs ...string) error {
	if len(pairs) == 0 || len(pairs)%2 != 0 {
		panic(errors.New("no imports specified"))
	}
	for i := 0; i < len(pairs)-1; i += 2 {
		if err := m.Add(pairs[i], pairs[i+1]); err != nil {
			return err
		}
	}
	return nil
}

func appendPaths(paths []string, prefix string, md *protokit.Descriptor, mdMap map[string]*protokit.Descriptor, seen map[string]struct{}) ([]string, error) {
	if seen == nil {
		seen = map[string]struct{}{}
	}

	for _, fd := range md.GetMessageFields() {
		fp := fd.GetName()
		if fd.OneofIndex != nil {
			fp = fmt.Sprintf("%s.%s", md.GetOneofDecl()[fd.GetOneofIndex()].GetName(), fp)
		}
		if prefix != "" {
			fp = fmt.Sprintf("%s.%s", prefix, fp)
		}

		if _, ok := seen[fd.GetFullName()]; ok {
			log.Printf("Field '%s' defined at %s is recursive, stopping traversal", fp, fd.GetFile().GetName())
			return paths, nil
		}
		seen[fd.GetFullName()] = struct{}{}

		paths = append(paths, fp)

		if fd.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED || fd.GetType() != descriptor.FieldDescriptorProto_TYPE_MESSAGE {
			delete(seen, fd.GetFullName())
			continue
		}

		fmd, ok := mdMap[fd.GetTypeName()]
		if !ok {
			switch fd.GetTypeName() {
			case protoTimestampType, protoFieldMaskType, protoDurationType, protoStructType, protoAnyType:
				delete(seen, fd.GetFullName())
				continue
			}
			return nil, unknownTypeError(fd.GetTypeName())
		}

		if len(fmd.GetMessageFields()) == 0 {
			delete(seen, fd.GetFullName())
			continue
		}

		var err error
		paths, err = appendPaths(paths, fp, fmd, mdMap, seen)
		if err != nil {
			return nil, err
		}
		delete(seen, fd.GetFullName())
	}
	for _, od := range md.GetOneofDecl() {
		paths = append(paths, od.GetName())
	}
	return paths, nil
}

type goType struct {
	Name       string
	Pkg        string
	IsNullable bool
	SetFielder bool
	Elem       *goType
}

var importPathReplacer = strings.NewReplacer(
	".", "_",
	"/", "_",
	"-", "_",
)

func (t goType) PkgAlias() string {
	return importPathReplacer.Replace(t.Pkg)
}

func (t goType) String() string {
	if t.Pkg == "" {
		return t.Name
	}
	return fmt.Sprintf("%s.%s", t.PkgAlias(), t.Name)
}

type goField struct {
	Name      string
	Type      goType
	Anonymous bool
}

func goFieldOf(fd *protokit.FieldDescriptor) goField {
	var typ goType

	switch fd.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		typ.Name = "bool"

	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		typ.Name = "float64"

	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		typ.Name = "float32"

	case descriptor.FieldDescriptorProto_TYPE_INT32, descriptor.FieldDescriptorProto_TYPE_SINT32, descriptor.FieldDescriptorProto_TYPE_SFIXED32:
		typ.Name = "int32"

	case descriptor.FieldDescriptorProto_TYPE_INT64, descriptor.FieldDescriptorProto_TYPE_SINT64, descriptor.FieldDescriptorProto_TYPE_SFIXED64:
		typ.Name = "int64"

	case descriptor.FieldDescriptorProto_TYPE_UINT32, descriptor.FieldDescriptorProto_TYPE_FIXED32:
		typ.Name = "uint32"

	case descriptor.FieldDescriptorProto_TYPE_UINT64, descriptor.FieldDescriptorProto_TYPE_FIXED64:
		typ.Name = "uint64"

	case descriptor.FieldDescriptorProto_TYPE_STRING:
		typ.Name = "string"

	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		typ.Name = "[]byte"

	case descriptor.FieldDescriptorProto_TYPE_GROUP:
		panic(unsupportedTypeError(fd.GetType().String()))

	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		typ.Name = fd.GetTypeName()
		if i := strings.LastIndex(typ.Name, "."); i > 0 {
			typ.Name = typ.Name[i+1:]
		}

		protoType := fd.GetTypeName()[1:]

		for parent := fd.GetMessage(); parent != nil; parent = parent.GetParent() {
			for _, sed := range parent.GetEnums() {
				if protoType == sed.GetFullName() {
					typ.Name = fmt.Sprintf("%s_%s", parent.GetName(), typ.Name)
					if i := strings.LastIndex(typ.Name, "."); i > 0 {
						typ.Name = typ.Name[i+1:]
					}
					protoType = parent.GetFullName()
				}
			}
		}

	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		typ.Name = fd.GetTypeName()
		switch typ.Name {
		case protoTimestampType:
			if v, ok := fd.OptionExtensions["gogoproto.stdtime"].(*bool); ok && *v {
				typ.Pkg = "time"
				typ.Name = "Time"
				break
			}
			panic(unsupportedTypeError(fd.GetTypeName()))

		case protoDurationType:
			if v, ok := fd.OptionExtensions["gogoproto.stdduration"].(*bool); ok && *v {
				typ.Pkg = "time"
				typ.Name = "Duration"
				break
			}
			panic(unsupportedTypeError(fd.GetTypeName()))

		case protoAnyType:
			typ.Pkg = "github.com/gogo/protobuf/types"
			typ.Name = "Any"

		case protoStructType:
			typ.Pkg = "github.com/gogo/protobuf/types"
			typ.Name = "Struct"

		case protoFieldMaskType:
			typ.Pkg = "github.com/gogo/protobuf/types"
			typ.Name = "FieldMask"

		default:
			typ.SetFielder = true
			typ.Name = fd.GetTypeName()
			if i := strings.LastIndex(typ.Name, "."); i > 0 {
				typ.Name = typ.Name[i+1:]
			}

			protoType := fd.GetTypeName()[1:]

			for parent := fd.GetMessage(); parent != nil; parent = parent.GetParent() {
				for _, smd := range parent.GetMessages() {
					if protoType == smd.GetFullName() {
						typ.Name = fmt.Sprintf("%s_%s", parent.GetName(), typ.Name)
						if i := strings.LastIndex(typ.Name, "."); i > 0 {
							typ.Name = typ.Name[i+1:]
						}
						protoType = parent.GetFullName()
					}
				}
			}
		}

	default:
		panic(unsupportedTypeError(fd.GetType().String()))
	}

	var isPtr bool
	if v, ok := fd.OptionExtensions["gogoproto.customtype"].(*string); ok {
		isPtr = true
		typ = goType{}
		if i := strings.LastIndex(*v, "."); i > 0 {
			typ.Pkg = (*v)[:i]
			typ.Name = (*v)[i+1:]
		} else {
			typ.Pkg = ""
			typ.Name = *v
		}
	}

	isPtr = isPtr || fd.GetType() == descriptor.FieldDescriptorProto_TYPE_MESSAGE
	if v, ok := fd.OptionExtensions["gogoproto.nullable"].(*bool); ok {
		isPtr = *v
	}

	if isPtr {
		elemType := typ
		typ = goType{
			IsNullable: true,
			SetFielder: elemType.SetFielder,
			Name:       fmt.Sprintf("*%s", elemType.String()),
			Elem:       &elemType,
		}
	}

	if fd.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED {
		elemType := typ
		if fd.GetType() == descriptor.FieldDescriptorProto_TYPE_MESSAGE &&
			fd.GetMessage().GetMessage(fd.GetTypeName()[1:]) != nil &&
			fd.GetMessage().GetMessage(fd.GetTypeName()[1:]).GetOptions().GetMapEntry() {
			typ = goType{
				IsNullable: true,
				Name:       fmt.Sprintf("map[%s]%s", goFieldOf(fd.GetMessage().GetMessage(fd.GetTypeName()[1:]).GetMessageField("key")).Type.String(), elemType.String()),
				Elem:       &elemType,
			}

		} else {
			typ = goType{
				IsNullable: true,
				Name:       fmt.Sprintf("[]%s", elemType.String()),
				Elem:       &elemType,
			}
		}
	}

	fieldName := generator.CamelCase(fd.GetName())
	if v, ok := fd.OptionExtensions["gogoproto.customname"].(*string); ok {
		fieldName = *v
	}

	if v, ok := fd.OptionExtensions["gogoproto.embed"].(*bool); ok && *v {
		return goField{
			Type:      typ,
			Anonymous: true,
		}
	}
	return goField{
		Name: fieldName,
		Type: typ,
	}
}

func buildIndented(buf *strings.Builder, tabCount uint, s string) {
	for _, l := range strings.Split(s, "\n") {
		fmt.Fprintln(buf, fmt.Sprintf("%s%s", strings.Repeat("	", int(tabCount)), l))
	}
}

func buildSetFieldsCase(buf *strings.Builder, imports importMap, tabCount uint, subs string, fd *protokit.FieldDescriptor) error {
	field := goFieldOf(fd)

	buildIndented(buf, tabCount, fmt.Sprintf(`case "%s":`, fd.GetName()))

	dstPath := "dst"
	srcPath := "src"
	if fd.OneofIndex != nil {
		md := fd.GetMessage()

		oneOfTypeName := fmt.Sprintf("%s_%s", md.GetName(), field.Name)
		if md.GetMessage(field.Name) != nil {
			oneOfTypeName = fmt.Sprintf("%s_", oneOfTypeName)
		}
		oneOfName := generator.CamelCase(md.GetOneofDecl()[fd.GetOneofIndex()].GetName())

		dstPath = fmt.Sprintf("%s.%s", dstPath, oneOfName)

		buildIndented(buf, tabCount+1, fmt.Sprintf(`if _, ok := %s.(*%s); !ok {
	%s = &%s{}
}`,
			dstPath, oneOfTypeName,
			dstPath, oneOfTypeName,
		))

		dstPath = fmt.Sprintf("%s.(*%s).%s", dstPath, oneOfTypeName, field.Name)
		srcPath = fmt.Sprintf("%s.Get%s()", srcPath, field.Name)

	} else {
		name := field.Name
		if field.Anonymous {
			name = field.Type.Name
		}
		dstPath = fmt.Sprintf("%s.%s", dstPath, name)
		srcPath = fmt.Sprintf("%s.%s", srcPath, name)
	}

	buildFinal := func(tabCount uint) error {
		buildIndented(buf, tabCount, fmt.Sprintf(`if src != nil {
	%s = %s
} else {`,
			dstPath, srcPath,
		))

		if field.Type.IsNullable {
			buildIndented(buf, tabCount, fmt.Sprintf(`	%s = nil
}`,
				dstPath,
			))
		} else {
			if err := imports.Add(field.Type.PkgAlias(), field.Type.Pkg); err != nil {
				return err
			}
			buildIndented(buf, tabCount, fmt.Sprintf(`	var zero %s
	%s = zero
}`,
				field.Type,
				dstPath,
			))
		}
		return nil
	}

	if !field.Type.SetFielder {
		if err := imports.Add("fmt", "fmt"); err != nil {
			return err
		}
		buildIndented(buf, tabCount+1, fmt.Sprintf(`if len(%s) > 0 {
	return fmt.Errorf("'%s' has no subfields, but %%s were specified", %s)
}`,
			subs,
			fd.GetName(), subs,
		))
		return buildFinal(tabCount + 1)
	}

	if field.Type.IsNullable {
		buildIndented(buf, tabCount+1, fmt.Sprintf(`if len(%s) > 0 {
	newDst := %s
	if newDst == nil {
		newDst = &%s{}
		%s = newDst
	}
	var newSrc %s
	if src != nil {
		newSrc = %s
	}`,
			subs,
			dstPath,
			field.Type.Elem,
			dstPath,
			field.Type,
			srcPath,
		))
	} else {
		buildIndented(buf, tabCount+1, fmt.Sprintf(`if len(%s) > 0 {
	newDst := &%s
	var newSrc *%s
	if src != nil {
		newSrc = &%s
	}`,
			subs,
			dstPath,
			field.Type,
			srcPath,
		))
	}

	buildIndented(buf, tabCount+1, `	if err := newDst.SetFields(newSrc, subs...); err != nil {
		return err
	}
} else {`)
	if err := buildFinal(tabCount + 2); err != nil {
		return err
	}
	buildIndented(buf, tabCount+1, `}`)
	return nil
}

func buildMethods(buf *strings.Builder, imports importMap, md *protokit.Descriptor, mdMap map[string]*protokit.Descriptor) error {
	if err := imports.Add("fmt", "fmt"); err != nil {
		return err
	}

	mType := md.GetName()
	for parent := md.GetParent(); parent != nil; parent = parent.GetParent() {
		mType = fmt.Sprintf("%s_%s", parent.GetName(), mType)
	}

	if len(md.GetMessageFields()) == 0 {
		fmt.Fprintf(buf, `
func (*%s) FieldMaskPaths(_ bool) []string {
	return nil
}

func (dst *%s) SetFields(src *%s, paths ...string) error {
	if len(paths) != 0 {
		return fmt.Errorf("message %s has no fields, but paths %%s were specified", paths)
	}
	if src != nil {
		*dst = *src
	}
	return nil
}`,
			mType,
			mType, mType,
			mType,
		)
		return nil
	}

	nestedPaths, err := appendPaths(make([]string, 0, len(md.GetMessageFields())), "", md, mdMap, nil)
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

	fmt.Fprintf(buf, `
var _%sFieldPathsNested = [...]string{
	%s
}

var _%sFieldPathsTopLevel = [...]string{
	%s
}

func (*%s) FieldMaskPaths(nested bool) []string {
	paths := _%sFieldPathsTopLevel[:]
	if nested {
		paths = _%sFieldPathsNested[:]
	}
	ret := make([]string, len(paths))
	copy(ret, paths)
	return ret
}

func (dst *%s) SetFields(src *%s, paths ...string) error {`,
		mType, `"`+strings.Join(nestedPaths, `",
	"`)+`",`,
		mType, `"`+strings.Join(topLevelPaths, `",
	"`)+`",`,
		mType,
		mType,
		mType,
		mType, mType,
	)

	fmt.Fprint(buf, `
	for name, subs := range _processPaths(paths) {
		switch name {
`,
	)

	oneOfs := make(map[int32][]*protokit.FieldDescriptor, len(md.GetMessageFields()))
	for _, fd := range md.GetMessageFields() {
		if fd.OneofIndex != nil {
			i := fd.GetOneofIndex()
			oneOfs[i] = append(oneOfs[i], fd)
			continue
		}

		if err := buildSetFieldsCase(buf, imports, 2, "subs", fd); err != nil {
			return err
		}
	}

	for i, fds := range oneOfs {
		declName := md.GetOneofDecl()[i].GetName()
		goName := generator.CamelCase(declName)

		fmt.Fprintln(buf)

		buildIndented(buf, 2, fmt.Sprintf(`case "%s":
	if len(subs) == 0 && src == nil {
		dst.%s = nil
		continue
	} else if len(subs) == 0 {
		dst.%s = src.%s
		continue
	}`,
			declName,
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

		for _, fd := range fds {
			if err := buildSetFieldsCase(buf, imports, 4, "oneofSubs", fd); err != nil {
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

type plugin struct{}

func (p plugin) Generate(in *plugin_go.CodeGeneratorRequest) (*plugin_go.CodeGeneratorResponse, error) {
	resp := &plugin_go.CodeGeneratorResponse{}

	fds := protokit.ParseCodeGenRequest(in)

	mdMap := map[string]*protokit.Descriptor{}
	for _, fd := range fds {
		if err := registerMessages(mdMap, fd.GetMessages()...); err != nil {
			return nil, err
		}
	}

	dirs := map[string]struct{}{}
	for _, fd := range fds {
		if len(fd.GetMessages()) == 0 {
			continue
		}

		dirName := fd.Options.GetGoPackage()
		if dirName == "" {
			dirName = filepath.Dir(fd.GetName())
		}
		fileName := filepath.Join(dirName, fmt.Sprintf("%s.pb.fm.go", strings.TrimSuffix(filepath.Base(fd.GetName()), filepath.Ext(fd.GetName()))))

		imports := importMap{}
		buf := &strings.Builder{}
		for _, md := range fd.GetMessages() {
			if v, ok := md.OptionExtensions["fieldmask.enable"].(*bool); ok && !*v {
				continue
			}

			var mBufs []*strings.Builder
			if err := walkMessage(md, func(md *protokit.Descriptor) error {
				if md.GetOptions().GetMapEntry() {
					return nil
				}

				mBuf := &strings.Builder{}
				if err := buildMethods(mBuf, imports, md, mdMap); err != nil {
					return err
				}

				if mBuf.Len() == 0 {
					return nil
				}
				mBufs = append(mBufs, mBuf)
				return nil
			}); err != nil {
				return nil, err
			}

			for _, mBuf := range mBufs {
				fmt.Fprintf(buf, `
%s`,
					mBuf.String())
			}
		}

		if buf.Len() == 0 {
			continue
		}

		dirs[dirName] = struct{}{}

		var importString string
		switch len(imports) {
		case 0:
		case 1:
			for name, pkg := range imports {
				importString = fmt.Sprintf(`
import %s "%s"`, name, pkg)
			}
		default:
			importLines := make([]string, 0, len(imports))
			for name, pkg := range imports {
				importLines = append(importLines, fmt.Sprintf(`	%s "%s"`, name, pkg))
			}
			sort.Slice(importLines, func(i, j int) bool {
				return strings.Fields(importLines[i])[1] < strings.Fields(importLines[j])[1]
			})
			importString = fmt.Sprintf(`
import (
%s
)`,
				strings.Join(importLines, "\n"))
		}

		resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
			Name: proto.String(fileName),
			Content: proto.String(fmt.Sprintf(`%s

package %s
%s%s
`,
				FileHeader,
				filepath.Base(dirName),
				importString,
				buf.String(),
			)),
		})
	}

	for dirName := range dirs {
		pkgName := filepath.Base(dirName)
		resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
			Name: proto.String(filepath.Join(dirName, fmt.Sprintf("%s.pb.util.fm.go", pkgName))),
			Content: proto.String(fmt.Sprintf(`%s
package %s

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
}
`,
				FileHeader,
				pkgName,
			)),
		})
	}

	return resp, nil
}

func init() {
	log.SetFlags(0)
}

func main() {
	if err := protokit.RunPlugin(plugin{}); err != nil {
		log.Fatalf("Failed to run plugin: %s", err)
	}
}
