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

type plugin struct{}

func (p *plugin) WriteMessage(w io.Writer, md *protokit.Descriptor, mds map[string]*protokit.Descriptor) {
	var allPaths []string

	fmt.Fprintf(w,
		`func (m *%s) SetFields(src *%s, mask *types.FieldMask) {
	if len(mask.GetPaths()) == 0 {
		mask = &types.FieldMask{Paths: m.FieldMaskPaths()}
	}

	for _, path := range mask.Paths {
		switch path {`,
		md.GetName(),
		md.GetName())

	for _, mfd := range md.GetMessageFields() {
		allPaths = append(allPaths, mfd.GetName())

		fmt.Fprintf(w, `
		case "%s":`, mfd.GetName())

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
			m.%s = &%s{%s: src.Get%s()}`, oneofName, oneofContainerName, name, name)
			continue
		}

		switch mfd.GetType() {
		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
			if embed, ok := mfd.OptionExtensions["gogoproto.embed"].(*bool); ok && *embed {
				name = mds[strings.TrimPrefix(mfd.GetTypeName(), ".")].GetName()
			}
			switch {
			case strings.HasPrefix(mfd.GetTypeName(), ".google.protobuf."), mfd.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED:
				fmt.Fprintf(w, `
			m.%s = src.%s`, name, name)

			default:
				if nullable, ok := mfd.OptionExtensions["gogoproto.nullable"].(*bool); !ok || *nullable {
					typeName := mds[strings.TrimPrefix(mfd.GetTypeName(), ".")].GetName()

					fmt.Fprintf(w, `
			if src.%s == nil {
				m.%s = nil
			} else {
				m.%s = &%s{}
				m.%s.SetFields(src.%s, FieldMaskWithoutPrefix(mask, "%s"))
			}`,
						name,
						name,
						name, typeName,
						name, name, mfd.GetName(),
					)
				} else {
					fmt.Fprintf(w, `
			m.%s.SetFields(&src.%s, FieldMaskWithoutPrefix(mask, "%s"))`, name, name, mfd.GetName())
				}
			}

		default:
			fmt.Fprintf(w, `
			m.%s = src.%s`, name, name)
		}
	}

	fmt.Fprintf(w, `
		}
	}
}

func (m *%s) FieldMaskPaths() []string {`,
		md.GetName())

	if len(allPaths) == 0 {
		fmt.Fprintf(w, `
	return nil`)
	} else {
		fmt.Fprintf(w, `
	return []string{"%s"}`, strings.Join(allPaths, "\", \""))
	}
	fmt.Fprintf(w, `
}`)
}

func (p *plugin) registerMessage(mds map[string]*protokit.Descriptor, md *protokit.Descriptor) {
	for _, sub := range md.GetMessages() {
		if strings.HasSuffix(sub.GetName(), "Entry") {
			continue
		}
		sub.Name = proto.String(fmt.Sprintf("%s_%s", md.GetName(), sub.GetName()))
		p.registerMessage(mds, sub)
	}
	mds[md.FullName] = md
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
	fds := protokit.ParseCodeGenRequest(in)

	files := make(map[string][]*protokit.Descriptor)
	for _, fd := range fds {
		dirName := filepath.Dir(fd.GetName())
		fileName := strings.TrimSuffix(filepath.Base(fd.GetName()), filepath.Ext(fd.GetName())) + ".pb.fm.go"
		if goPackage := fd.Options.GetGoPackage(); goPackage != "" {
			dirName = fd.Options.GetGoPackage()
		}
		fullName := filepath.Join(dirName, fileName)
		for _, md := range fd.GetMessages() {
			files[fullName] = p.appendDescriptor(files[fullName], md)
		}
	}

	mdMap := make(map[string]*protokit.Descriptor)
	for _, fd := range fds {
		for _, md := range fd.GetMessages() {
			p.registerMessage(mdMap, md)
		}
	}

	resp := &plugin_go.CodeGeneratorResponse{}

	dirs := make(map[string]struct{})
	for fileName, mds := range files {
		if len(mds) == 0 {
			continue
		}
		dirName := filepath.Dir(fileName)
		dirs[dirName] = struct{}{}

		packageName := filepath.Base(dirName)

		buf := &bytes.Buffer{}
		fmt.Fprintf(buf, `%s

package %s

import "github.com/gogo/protobuf/types"`,
			util.FileHeader,
			packageName)

		for _, md := range mds {
			fmt.Fprintln(buf)
			fmt.Fprintln(buf)
			p.WriteMessage(buf, md, mdMap)
		}
		fmt.Fprintln(buf)

		resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
			Name:    proto.String(fileName),
			Content: proto.String(buf.String()),
		})
	}

	for dir := range dirs {
		buf := &bytes.Buffer{}
		_, err := util.Write(buf, filepath.Base(dir))
		if err != nil {
			log.Fatalf("Failed to write utilities to buffer: `%s`", err)
		}
		fmt.Fprintln(buf)

		resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
			Name:    proto.String(filepath.Join(dir, "field_mask_util.pb.fm.go")),
			Content: proto.String(buf.String()),
		})
	}

	return resp, nil
}

func main() {
	if err := protokit.RunPlugin(&plugin{}); err != nil {
		log.Fatalf("Failed to run plugin: %s", err)
	}
}
