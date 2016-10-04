UNMAKE := $${GOPATH}/bin/unmake

.PHONY: build
build:
	@# build unmake

	go fmt *.go
	go build


.PHONY: example
example:
	@# show example usage

	@$(UNMAKE) examples/Makefile
