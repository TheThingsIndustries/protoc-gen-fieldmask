package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"path/filepath"
	"strings"

	_ "github.com/TheThingsIndustries/protoc-gen-fieldmask/gogoproto"
	"github.com/TheThingsIndustries/protoc-gen-fieldmask/util"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pseudomuto/protokit"
)

func init() {
	log.SetFlags(0)
}

func main() {
	if err := protokit.RunPlugin(&plugin{}); err != nil {
		log.Fatalf("Failed to run plugin: %s", err)
	}
}

type plugin struct{}

func (p *plugin) WriteMessage(w io.Writer, md *protokit.Descriptor, messages map[string]*protokit.Descriptor) {
	var allPaths []string

	fmt.Fprintf(w,
		`func (m *%s) SetFields(src *%s, mask *types.FieldMask) {
\tif len(mask.GetPaths()) == 0 {
\t\tmask = &types.FieldMask{Paths: m.FieldMaskPaths()}
\t}

\tfor _, path := range mask.Paths {
\t\tswitch path {`,
		md.GetName(),
		md.GetName())

	for _, mfd := range md.GetMessageFields() {
		allPaths = append(allPaths, mfd.GetName())

		fmt.Fprintf(w, `
\t\t\tcase "%s":`, mfd.GetName())

		name := generator.CamelCase(mfd.GetName())
		if customName, ok := mfd.OptionExtensions["gogoproto.customname"].(*string); ok {
			name = *customName
		}

		if mfd.OneofIndex != nil {
			oneofName := generator.CamelCase(md.GetOneofDecl()[*mfd.OneofIndex].GetName())
			oneofContainerName := fmt.Sprintf("%s_%s", md.GetName(), name)
			if md.GetMessage(oneofContainerName) != nil {
				oneofContainerName += "_"
			}
			fmt.Fprintf(w, `
\t\t\t\tm.%s = &%s{%s: src.Get%s()}`, oneofName, oneofContainerName, name, name)
			continue
		}

		switch mfd.GetType() {
		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
			if embed, ok := mfd.OptionExtensions["gogoproto.embed"].(*bool); ok && *embed {
				name = messages[strings.TrimPrefix(mfd.GetTypeName(), ".")].GetName()
			}
			switch {
			case strings.HasPrefix(mfd.GetTypeName(), ".google.protobuf."), mfd.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED:
				fmt.Fprintf(w, `
\t\t\t\tm.%s = src.%s`, name, name)

			default:
				if nullable, ok := mfd.OptionExtensions["gogoproto.nullable"].(*bool); !ok || *nullable {
					typeName := messages[strings.TrimPrefix(mfd.GetTypeName(), ".")].GetName()

					fmt.Fprintf(w, `
\t\t\t\tif src.%s == nil {
\t\t\t\t\tm.%s = nil
\t\t\t\t} else {
\t\t\t\t\tm.%s = &%s{}
\t\t\t\t\tm.%s.SetFields(src.%s, FieldMaskWithoutPrefix(mask, "%s"))
\t\t\t\t}`,
						name,
						name,
						name, typeName,
						name, name, mfd.GetName(),
					)
				} else {
					fmt.Fprintf(w, `
\t\t\t\tm.%s.SetFields(&src.%s, FieldMaskWithoutPrefix(mask, "%s"))`, name, name, mfd.GetName())
				}
			}

		default:
			fmt.Fprintf(w, `
\t\t\t\tm.%s = src.%s`, name, name)
		}
	}

	fmt.Fprintf(w, `
\t\t\t}
\t\t}
\t}
}

func (m *%s) FieldMaskPaths() []string {`,
		md.GetName())

	wrappedPaths := make([]string, len(allPaths))
	for i, path := range allPaths {
		wrappedPaths[i] = `"` + path + `"`
	}

	fmt.Fprintf(w, `
\treturn []string{%s}
}`,
		strings.Join(wrappedPaths, ", "))
}

func (p *plugin) registerMessage(messages map[string]*protokit.Descriptor, md *protokit.Descriptor) {
	for _, sub := range md.GetMessages() {
		if strings.HasSuffix(sub.GetName(), "Entry") {
			continue
		}
		sub.Name = proto.String(fmt.Sprintf("%s_%s", md.GetName(), sub.GetName()))
		p.registerMessage(messages, sub)
	}
	messages[md.FullName] = md
}

func (p *plugin) appendDescriptor(s []*protokit.Descriptor, md *protokit.Descriptor) []*protokit.Descriptor {
	for _, mfd := range md.GetMessageFields() {
		if mfd.GetTypeName() == ".google.protobuf.FieldMask" {
			return s
		}
	}
	out := append(s, md)
	for _, sub := range md.GetMessages() {
		if strings.HasSuffix(sub.GetName(), "Entry") {
			continue
		}
		out = p.appendDescriptor(out, sub)
	}
	return out
}

func (p *plugin) Generate(in *plugin_go.CodeGeneratorRequest) (*plugin_go.CodeGeneratorResponse, error) {
	descriptors := protokit.ParseCodeGenRequest(in)

	messages := make(map[string]*protokit.Descriptor)
	for _, fd := range descriptors {
		for _, md := range fd.GetMessages() {
			p.registerMessage(messages, md)
		}
	}

	files := make(map[string][]*protokit.Descriptor)

	for _, fd := range descriptors {
		dirname := filepath.Dir(fd.GetName())
		filename := strings.TrimSuffix(filepath.Base(fd.GetName()), filepath.Ext(fd.GetName())) + ".pb.fm.go"
		if goPackage := fd.Options.GetGoPackage(); goPackage != "" {
			dirname = fd.Options.GetGoPackage()
		}
		fullName := filepath.Join(dirname, filename)
		for _, md := range fd.GetMessages() {
			files[fullName] = p.appendDescriptor(files[fullName], md)
		}
	}

	resp := &plugin_go.CodeGeneratorResponse{}

	dirs := make(map[string]struct{})
	for filename, mds := range files {
		if len(mds) == 0 {
			continue
		}
		dirname := filepath.Dir(filename)
		dirs[dirname] = struct{}{}

		packageName := filepath.Base(dirname)

		buf := &bytes.Buffer{}

		fmt.Fprintf(buf,
			`%s

package %s

import "github.com/gogo/protobuf/types"`,
			util.FileHeader,
			packageName)

		for _, md := range mds {
			p.WriteMessage(buf, md, messages)
		}

		resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
			Name:    proto.String(filename),
			Content: proto.String(buf.String()),
		})
	}

	for dir := range dirs {
		buf := &bytes.Buffer{}
		util.Generate(buf, filepath.Base(dir))
		resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
			Name:    proto.String(filepath.Join(dir, "field_mask_util.pb.fm.go")),
			Content: proto.String(buf.String()),
		})
	}

	return resp, nil
}
