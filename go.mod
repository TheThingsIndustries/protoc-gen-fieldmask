module github.com/TheThingsIndustries/protoc-gen-fieldmask

go 1.15

replace github.com/lyft/protoc-gen-star => github.com/TheThingsIndustries/protoc-gen-star v0.5.2-gogo.1

replace github.com/golang/protobuf => github.com/golang/protobuf v1.3.5

replace github.com/envoyproxy/protoc-gen-validate => github.com/TheThingsIndustries/protoc-gen-validate v0.5.1-fieldmask.2

require (
	github.com/envoyproxy/protoc-gen-validate v0.5.1
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.0
	github.com/kr/pretty v0.1.0
	github.com/lyft/protoc-gen-star v0.5.2
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/smartystreets/assertions v1.0.1
)
