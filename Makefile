.default: build

.PHONY: test

test:
	go test -coverprofile=coverage.out ./...

.PHONY: build

build:
	CGO_ENABLED=0 go build -ldflags "-s -w" -o dist/protoc-gen-fieldmask-$(shell go env GOOS)-$(shell go env GOARCH)$(shell go env GOEXE) .
