//+build ignore

package main

import (
	"bytes"
	"io/ioutil"

	"github.com/TheThingsIndustries/protoc-gen-fieldmask/utils"
)

func main() {
	buf := new(bytes.Buffer)
	utils.Generate(buf, "utils")
	ioutil.WriteFile("utils.generated.go", buf.Bytes(), 0644)
}
