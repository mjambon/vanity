# Build the example HTML document and a PNG graph of terms using 'make'.
#
.PHONY: all
all: example.html example.png

AD = ../bin/ad
DEPS = ../bin/ad

example.html: example.yml $(DEPS)
	../bin/ad < example.yml > example.html

example.dot: example.yml $(DEPS)
	../bin/ad -t dot < example.yml > example.dot

example.png: example.dot $(DEPS)
	dot -Tpng example.dot -o example.png
