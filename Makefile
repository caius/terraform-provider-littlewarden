terraform-provider-littlewarden: littlewarden/*.go main.go go.mod .tool-versions
	go build

.PHONY: build
build: terraform-provider-littlewarden

littlewarden/littlewarden.test: littlewarden/*.go go.mod .tool-versions
	cd littlewarden && go test -c
	littlewarden/littlewarden.test

.PHONY: test
test: littlewarden/littlewarden.test

.PHONY: example
example: terraform-provider-littlewarden
	cd example && terraform init
	cd example && terraform plan
