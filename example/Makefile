# Build the example HTML document and a PNG graph of terms using 'make'.
#
.PHONY: all
all: example.html

AD = ../bin/ad
DEPS = ../bin/ad

example.html: example.yml $(DEPS)
	../bin/ad < example.yml > example.html

example.dot: example.yml $(DEPS)
	touch example.dot

example.png: example.dot $(DEPS)
	#dot -Tpng example.dot -o example.png
	touch example.png
