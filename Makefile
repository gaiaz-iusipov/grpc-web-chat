LOCAL_BIN:=$(CURDIR)/bin

.PHONY: install-bins
install-bins:
	GOBIN=$(LOCAL_BIN) go install github.com/bufbuild/buf/cmd/buf@latest

.PHONY: generate
generate:
	$(LOCAL_BIN)/buf generate proto
