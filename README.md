# protoc-gen-fieldmask

Generate Field Mask Utilities for GoGo protos.

## Installation:

```sh
dep ensure
go install .
```

## Usage:

For example, in `TheThingsIndustries/lorawan-stack`:

```sh
protoc -I $(dirname $PWD) -I $GOPATH/src -I /usr/local/include --fieldmask_out=$GOPATH/src $PWD/api/*.proto
```
