LOCAL_BIN:=$(CURDIR)/bin

.PHONY: install-bins
install-bins:
	GOBIN=$(LOCAL_BIN) go install github.com/bufbuild/buf/cmd/buf@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0

.PHONY: generate
generate:
	$(LOCAL_BIN)/buf generate proto
