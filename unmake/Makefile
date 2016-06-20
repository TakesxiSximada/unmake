.PHONY: help
help:
	@# Display usage.
	@./unmake/usage Makefile

.PHONY: unmake-sync
unmake-sync:
	@# Sync unmake
	@curl -L -o ./unmake/usage https://raw.githubusercontent.com/TakesxiSximada/unmake/develop/unmake/usage
	@curl -L -o ./unmake/usage https://raw.githubusercontent.com/TakesxiSximada/unmake/develop/unmake/search
	@curl -L -o ./unmake/Makefile https://raw.githubusercontent.com/TakesxiSximada/unmake/develop/unmake/Makefile
	@chmod 755 ./unmake/usage
	@chmod 755 ./unmake/search

.PHONY: unmake-save
unmake-save:
	@# Save unmake
	@git add ./unmake
	@git commit -m "Update unmake" ./unmake

.PHONY: unmake-uninstall
unmake-uninstall:
	@# Remove unmake
	@git rm -rf unmake
	@git commit -m "Uninstall unmake" ./unmake
