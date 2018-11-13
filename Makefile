.DEFAULT_GOAL=build

.PHONY: build

build: internal/extensions/gogoproto/gogo.pb.go
	CGO_ENABLED=0 go build -ldflags "-w -s" -o dist/protoc-gen-fieldmask-$(shell go env GOOS)-$(shell go env GOARCH)$(shell go env GOEXE) .

WORKDIR:=$(shell mkdir -p $(PWD)/.work && mktemp -d -p $(PWD)/.work)

DOCKER ?= docker
PROTOC_DOCKER_IMAGE ?= thethingsindustries/protoc:3.0.15
PROTOC_DOCKER_ARGS := run --user `id -u` --rm \
										 --mount type=bind,src=$(PWD),dst=$(PWD) \
										 -e IN_TEST \
										 -w $(PWD)
PROTOC ?= $(DOCKER) $(PROTOC_DOCKER_ARGS) $(PROTOC_DOCKER_IMAGE)

vendor/github.com/gogo/protobuf/gogoproto/gogo.proto:
	dep ensure

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
	find ./testdata -name '*.pb.go' -delete -or -name '*.pb.fm.go' -delete

.PHONY: test

test:
	$(info Regenerating golden files...)
	@TMPDIR="$(WORKDIR)" WORKDIR="$(WORKDIR)" PROTOC="$(PROTOC)" go test -regenerate
	$(info Running tests...)
	@TMPDIR="$(WORKDIR)" WORKDIR="$(WORKDIR)" PROTOC="$(PROTOC)" go test -coverprofile=coverage.out ./...
