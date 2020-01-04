terraform-provider-littlewarden: littlewarden/*.go main.go go.mod .tool-versions
	go build

.PHONY: build
build: terraform-provider-littlewarden

.PHONY: test
test:
	go test
