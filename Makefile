# Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
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
	rm -rf dist .tools vendor

.tools/protoc-gen-go: go.mod go.sum
	go build -o $@ google.golang.org/protobuf/cmd/protoc-gen-go

.PHONY: test

PROTOC ?= protoc
PROTOC += --plugin=protoc-gen-go=.tools/protoc-gen-go

test: .tools/protoc-gen-go
	$(info Regenerating golden files...)
	@PROTOC="$(PROTOC)" go test -regenerate
	$(info Running tests...)
	@PROTOC="$(PROTOC)" go test -coverprofile=coverage.out ./...

benchmark: .tools/protoc-gen-go
	$(info Running benchmarks...)
	@PROTOC="$(PROTOC)" go test -bench .
