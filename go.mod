module github.com/TheThingsIndustries/protoc-gen-fieldmask

go 1.13

replace github.com/lyft/protoc-gen-star => github.com/TheThingsIndustries/protoc-gen-star v0.4.14-gogo.2

replace github.com/envoyproxy/protoc-gen-validate => github.com/TheThingsIndustries/protoc-gen-validate v0.2.0-java-fieldmask.6

require (
	github.com/envoyproxy/protoc-gen-validate v0.2.0-java
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.3
	github.com/kr/pretty v0.1.0
	github.com/lyft/protoc-gen-star v0.4.14
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/smartystreets/assertions v1.0.1
)
