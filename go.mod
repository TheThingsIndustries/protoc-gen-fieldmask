module github.com/TheThingsIndustries/protoc-gen-fieldmask

go 1.13

replace github.com/lyft/protoc-gen-star => github.com/TheThingsIndustries/protoc-gen-star v0.4.12-gogo

replace github.com/envoyproxy/protoc-gen-validate => github.com/TheThingsIndustries/protoc-gen-validate v0.2.0-java-fieldmask.5

require (
	github.com/envoyproxy/protoc-gen-validate v0.2.0-java
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.2
	github.com/kr/pretty v0.1.0
	github.com/lyft/protoc-gen-star v0.4.12
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/smartystreets/assertions v1.0.1
	golang.org/x/lint v0.0.0-20191125180803-fdd1cda4f05f // indirect
	golang.org/x/tools v0.0.0-20191206204035-259af5ff87bd // indirect
)
