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

.PHONY: clean
clean:
	git clean -dfX
