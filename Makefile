UNMAKE := $${GOPATH}/bin/unmake

.PHONY: build
build:
	@# build unmake

	go build


.PHONY: example
example:
	@# show example usage

	@$(UNMAKE) examples/Makefile
