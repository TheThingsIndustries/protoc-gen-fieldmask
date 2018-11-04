//+build ignore

package main

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/TheThingsIndustries/protoc-gen-fieldmask/util"
)

const fileName = "util.generated.go"

func init() {
	log.SetFlags(0)
}

func main() {
	buf := &bytes.Buffer{}
	_, err := util.Write(buf, "util")
	if err != nil {
		log.Fatalf("Failed to write to buffer: %s", err)
	}

	if err = ioutil.WriteFile(fileName, buf.Bytes(), 0644); err != nil {
		log.Fatalf("Failed to write buffer to %s: %s", fileName, err)
	}
}
