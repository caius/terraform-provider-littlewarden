terraform-provider-littlewarden: littlewarden/*.go main.go go.mod .tool-versions
	go build

.PHONY: build
build: terraform-provider-littlewarden

terraform-provider-littlewarden.test: littlewarden/*.go main.go go.mod .tool-versions
	go test -c

.PHONY: test
test: terraform-provider-littlewarden.test

.PHONY: example
example: terraform-provider-littlewarden
	cd example && terraform init
	cd example && terraform plan
