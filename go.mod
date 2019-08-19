module github.com/TheThingsIndustries/protoc-gen-fieldmask

go 1.12

replace github.com/lyft/protoc-gen-star => github.com/TheThingsIndustries/protoc-gen-star v0.4.11-gogo

replace github.com/envoyproxy/protoc-gen-validate => github.com/TheThingsIndustries/protoc-gen-validate v0.2.0-java-fieldmask

require (
	github.com/envoyproxy/protoc-gen-validate v0.2.0-java
	github.com/gogo/protobuf v1.2.1
	github.com/golang/protobuf v1.3.2
	github.com/kr/pretty v0.1.0
	github.com/lyft/protoc-gen-star v0.4.11
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/smartystreets/assertions v1.0.1
	golang.org/x/tools v0.0.0-20190819174341-15fda70baffd // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)
