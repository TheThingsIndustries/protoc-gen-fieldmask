# protoc-gen-fieldmask

A protoc plug-in, which generates fieldmask utilities. Compatible with gogoproto extensions.

## Installation:

```sh
dep ensure # If not using Go 1.11 modules
go install .
```

## Usage:

For example, in `TheThingsIndustries/lorawan-stack`:

```sh
protoc -I $(dirname $PWD) -I $GOPATH/src -I /usr/local/include --fieldmask_out=$GOPATH/src $PWD/api/*.proto
```
