.PHONY: build
build:
	$(MAKE) -C src build

.PHONY: clean
clean:
	git clean -dfX
