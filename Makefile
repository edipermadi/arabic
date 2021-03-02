GO	:= $(shell which go)

.PHONY: test
test:
	$(GO) test -v ./...
