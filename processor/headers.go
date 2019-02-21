package processor

import (
	pgs "github.com/lyft/protoc-gen-star"
	"strings"
)

func HeaderPrepender(header string) pgs.PostProcessor { return headerPrepender(header) }

type headerPrepender string

func (p headerPrepender) Match(a pgs.Artifact) bool {
	var n string
	switch a := a.(type) {
	case pgs.GeneratorFile:
		n = a.Name
	case pgs.GeneratorTemplateFile:
		n = a.Name
	case pgs.CustomFile:
		n = a.Name
	case pgs.CustomTemplateFile:
		n = a.Name
	default:
		return false
	}
	return strings.HasSuffix(n, ".go")
}

func (p headerPrepender) Process(in []byte) ([]byte, error) {
	return append([]byte(p+"\n\n"), in...), nil
}
