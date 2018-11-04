//+build ignore

package main

import (
	"bytes"
	"io/ioutil"

	"github.com/TheThingsIndustries/protoc-gen-fieldmask/util"
)

func main() {
	buf := new(bytes.Buffer)
	util.Generate(buf, "utils")
	ioutil.WriteFile("util.generated.go", buf.Bytes(), 0644)
}
