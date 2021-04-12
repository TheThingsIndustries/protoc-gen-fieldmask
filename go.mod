module github.com/TheThingsIndustries/protoc-gen-fieldmask

go 1.15

replace github.com/lyft/protoc-gen-star => github.com/TheThingsIndustries/protoc-gen-star v0.5.2-gogo.1

replace github.com/golang/protobuf => github.com/golang/protobuf v1.3.5

replace github.com/envoyproxy/protoc-gen-validate => github.com/TheThingsIndustries/protoc-gen-validate v0.4.2-0.20210412112233-c6a2014668e2

require (
	github.com/envoyproxy/protoc-gen-validate v0.4.2-0.20210412095728-272bdc88182a
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.5
	github.com/kr/pretty v0.1.0
	github.com/lyft/protoc-gen-star v0.5.2
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/smartystreets/assertions v1.0.1
)
