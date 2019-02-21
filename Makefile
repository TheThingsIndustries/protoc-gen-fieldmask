# Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

WORKDIR := $(shell mkdir -p $(PWD)/.work && mktemp -d "$(PWD)/.work/tmp.XXX")

DOCKER ?= docker
PROTOC_DOCKER_IMAGE ?= thethingsindustries/protoc:3.0.24
PROTOC_DOCKER_ARGS := run --user `id -u` --rm \
										 --mount type=bind,src=$(PWD),dst=$(PWD) \
										 -e IN_TEST \
										 -w $(PWD)
PROTOC ?= $(DOCKER) $(PROTOC_DOCKER_ARGS) $(PROTOC_DOCKER_IMAGE)

.DEFAULT_GOAL=build

.PHONY: all

all: build

.PHONY: build

build:
	CGO_ENABLED=0 go build -ldflags "-w -s" -o dist/protoc-gen-fieldmask .

.PHONY: clean

clean:
	rm -rf .work dist
	find ./testdata -name '*.pb.go' -delete -or -name '*.pb.fm.go' -delete

.PHONY: test

test:
	$(info Regenerating golden files...)
	@TMPDIR="$(WORKDIR)" WORKDIR="$(WORKDIR)" PROTOC="$(PROTOC)" go test -regenerate
	$(info Running tests...)
	@TMPDIR="$(WORKDIR)" WORKDIR="$(WORKDIR)" PROTOC="$(PROTOC)" go test -coverprofile=coverage.out ./...
