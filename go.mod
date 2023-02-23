module github.com/TheThingsIndustries/protoc-gen-fieldmask

go 1.18

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

require (
	github.com/iancoleman/strcase v0.1.2 // indirect
	github.com/kr/text v0.1.0 // indirect
	github.com/spf13/afero v1.4.1 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/text v0.3.8 // indirect
	golang.org/x/tools v0.1.12 // indirect
)
