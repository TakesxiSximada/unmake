UNMAKE := $${GOPATH}/bin/unmake

.PHONY: build
build:
	@# build unmake

	go fmt *.go
	golint -set_exit_status=1
	go build


.PHONY: example
example:
	@# show example usage

	@$(UNMAKE) examples/Makefile
