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

export GO111MODULE=on

.DEFAULT_GOAL=build

.PHONY: all

all: build

.PHONY: build

build:
	CGO_ENABLED=0 go build -ldflags "-w -s" -o dist/protoc-gen-fieldmask .

.PHONY: clean

clean:
	rm -rf dist .tools

.tools/protoc-gen-gogo: go.mod go.sum
	go build -o $@ github.com/gogo/protobuf/protoc-gen-gogo

vendor/github.com/gogo/protobuf/gogoproto/gogo.proto: go.mod go.sum
	go mod vendor

.PHONY: test

PROTOC ?= protoc
PROTOC += --plugin=protoc-gen-gogo=.tools/protoc-gen-gogo

test: .tools/protoc-gen-gogo vendor/github.com/gogo/protobuf/gogoproto/gogo.proto
	$(info Regenerating golden files...)
	@PROTOC="$(PROTOC)" go test -regenerate
	$(info Running tests...)
	@PROTOC="$(PROTOC)" go test -coverprofile=coverage.out ./...
