# Build the 'ad' executable and the example using 'make'.
#
.PHONY: all
all: example

.PHONY: build
build:
	$(MAKE) -C src build

.PHONY: example
example: build
	$(MAKE) -C example

.PHONY: clean
clean:
	git clean -dfX
