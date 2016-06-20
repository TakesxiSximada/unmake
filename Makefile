include unmake/Makefile *.mk # for commonmake


.PHONY: all
all:
	@unmake/search Makefile
