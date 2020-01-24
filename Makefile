.PHONY: all
all: example

.PHONY: build
build:
	$(MAKE) -C src build

.PHONY: example
example: build
	./bin/ad < example.yml > example.html

.PHONY: clean
clean:
	git clean -dfX
