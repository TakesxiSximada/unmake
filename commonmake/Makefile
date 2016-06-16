.PHONY: help
help:
	@# Display usage.
	@./commonmake/usage Makefile

commnmake-sync:
	@# Sync commonmake
	@curl -L -o ./commonmake/usage https://raw.githubusercontent.com/TakesxiSximada/commonmake/develop/commonmake/usage
	@curl -L -o ./commonmake/Makefile https://raw.githubusercontent.com/TakesxiSximada/commonmake/develop/commonmake/Makefile
	@chmod 755 ./commonmake/usage
	@git add ./commonmake/usage ./commonmake/Makefile
	@git commit -m "sync commonmake" ./commonmake/usage ./commonmake/Makefile

commnmake-remove:
	@# Remove commonmake
	@./commonmake/usage Makefile
