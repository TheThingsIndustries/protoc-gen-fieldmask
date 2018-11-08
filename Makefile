.default: build

.PHONY: test

test:
	go test -coverprofile=coverage.out ./...

.PHONY: build

build: internal/extensions/gogoproto/gogo.pb.go
	CGO_ENABLED=0 go build -ldflags "-s -w" -o dist/protoc-gen-fieldmask-$(shell go env GOOS)-$(shell go env GOARCH)$(shell go env GOEXE) .

WORKDIR:=$(shell mkdir -p $(PWD)/.work && mktemp -d -p $(PWD)/.work)

DOCKER ?= docker
PROTOC_DOCKER_IMAGE ?= thethingsindustries/protoc:3.0.15
PROTOC_DOCKER_ARGS := run --user `id -u` --rm \
										 --mount type=bind,src=$(WORKDIR),dst=$(WORKDIR) \
										 --mount type=bind,src=$(PWD)/vendor,dst=$(PWD)/vendor \
										 --mount type=bind,src=$(PWD)/internal/extensions,dst=$(PWD)/internal/extensions \
										 -w $(PWD)
PROTOC ?= $(DOCKER) $(PROTOC_DOCKER_ARGS) $(PROTOC_DOCKER_IMAGE)

internal/extensions/gogoproto/gogo.pb.go: vendor/github.com/gogo/protobuf/gogoproto/gogo.proto
	perl \
		-pe 's!(.*option[[:space:]]+.*go_package.*=.*"github.com/)gogo/protobuf(/gogoproto".*)!\1TheThingsIndustries/protoc-gen-fieldmask/internal/extensions\2!' \
		$< > $(WORKDIR)/gogo.proto
	$(PROTOC) -I$(WORKDIR) -I$(PWD)/vendor --go_out=$(WORKDIR) $(WORKDIR)/gogo.proto
	mv $(WORKDIR)/github.com/TheThingsIndustries/protoc-gen-fieldmask/internal/extensions/gogoproto/gogo.pb.go $@

.PHONY: extensions

extensions: internal/extensions/gogoproto/gogo.pb.go

.PHONY: clean

clean:
	rm -rf .work
