# Build the 'ad' executable and the examples using 'make'.
#
.PHONY: all
all: examples

.PHONY: build
build:
	$(MAKE) -C src build

.PHONY: examples
examples: build
	$(MAKE) -C examples

.PHONY: install
install:
	sudo mkdir -p /usr/local/bin
	sudo cp bin/ad /usr/local/bin

.PHONY: uninstall
uninstall:
	sudo rm /usr/local/bin/ad

.PHONY: clean
clean:
	git clean -dfX
