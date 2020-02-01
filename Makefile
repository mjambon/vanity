# Build the 'vanity' executable using 'make'.
#
.PHONY: all
all: build

.PHONY: build
build:
	$(MAKE) -C src build

.PHONY: release
release:
	@echo "Building release binaries."
	$(MAKE) -C src release
	@echo "Don't forget to tag the git commit:"
	@echo "  git tag `cat VERSION`"
	@echo "  git push origin master --tags"
	@echo "Then create a new release on Github and upload the binaries."

# Build the examples. Requires 'dot' command from Graphviz.
.PHONY: examples
examples: build
	$(MAKE) -C examples

# Run test suite
.PHONY: test
test:
	$(MAKE) -C test

# Accept new test results
.PHONY: accept
accept:
	$(MAKE) -C test accept

.PHONY: install
install:
	sudo mkdir -p /usr/local/bin
	sudo cp bin/vanity /usr/local/bin

.PHONY: uninstall
uninstall:
	sudo rm /usr/local/bin/vanity

.PHONY: clean
clean:
	git clean -dfX
