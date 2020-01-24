.PHONY: build
build:
	$(MAKE) -C src build

.PHONY: test
test: build
	./bin/ad < example.yml

.PHONY: clean
clean:
	git clean -dfX
