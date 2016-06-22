ROOT_MAKEFILE := Makefile
.PHONY: run
run:
	@## target=NO
	@#
	@# execute an application.
	@#

	@$${GOPATH}/bin/unmake $(ROOT_MAKEFILE)
	@echo "RUN...."
