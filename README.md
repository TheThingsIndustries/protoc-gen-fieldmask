# protoc-gen-fieldmask

A protoc plug-in, which generates fieldmask utilities. Compatible with gogoproto extensions.

## Installation:

```sh
go install .
```

## Usage:

For example, in `TheThingsIndustries/lorawan-stack`:

```sh
protoc -I $(dirname $PWD) -I $GOPATH/src -I /usr/local/include --fieldmask_out=$GOPATH/src $PWD/api/*.proto
```
