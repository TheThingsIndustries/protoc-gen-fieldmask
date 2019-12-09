# protoc-gen-fieldmask

A protoc plug-in, which generates fieldmask utilities. Compatible with gogoproto extensions.

## Installation:

```sh
GO111MODULE=on go install .
```

## Usage:

For example, from root of this repository:

```sh
protoc -Itestdata -Ivendor --fieldmask_out=lang=gogo:$GOPATH/src testdata/testdata.proto
```
Note, you will need to run `GO111MODULE=on go mod vendor` before running the command above.
