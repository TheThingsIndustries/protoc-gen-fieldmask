module github.com/TheThingsIndustries/protoc-gen-fieldmask

go 1.19

replace github.com/envoyproxy/protoc-gen-validate => github.com/TheThingsIndustries/protoc-gen-validate v0.5.1-fieldmask.6

require (
	github.com/envoyproxy/protoc-gen-validate v0.10.1
	github.com/kr/pretty v0.3.1
	github.com/lyft/protoc-gen-star/v2 v2.0.3
	github.com/smartystreets/assertions v1.13.0
	golang.org/x/exp v0.0.0-20230321023759-10a507213a29
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.10.0 // indirect
	github.com/spf13/afero v1.9.5 // indirect
	golang.org/x/mod v0.9.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	golang.org/x/tools v0.7.0 // indirect
)
