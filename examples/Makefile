# Build the example HTML document and a PNG graph of terms using 'make'.
#
.PHONY: all
all: simple.html simple.png complex.html complex.png vanity.html vanity.png

DEPS = ../bin/vanity

### A simple example. Exercises few options, relying on defaults.

simple.html: simple.yml $(DEPS)
	../bin/vanity -s < simple.yml > simple.html

simple.dot: simple.yml $(DEPS)
	../bin/vanity -t dot < simple.yml > simple.dot

# Render graph. Needs 'dot' program, normally from a 'graphviz' package.
simple.png: simple.dot $(DEPS)
	dot -Tpng simple.dot -o simple.png

### A complex example. Tries to demonstrate most available features.

complex.html: complex.yml $(DEPS)
	../bin/vanity -s \
	  --title='Glossary' \
	  --include-in-header=include/in-header.html \
	  --include-before-body=include/before-body.html \
	  --include-after-body=include/after-body.html \
	  < complex.yml > complex.html

complex.dot: complex.yml $(DEPS)
	../bin/vanity -t dot \
	  --include-before-body=include/before-body.dot \
	  < complex.yml > complex.dot

complex.png: complex.dot $(DEPS)
	dot -Tpng complex.dot -o complex.png

### Dogfooding

vanity.html: vanity.yml $(DEPS)
	../bin/vanity < vanity.yml > vanity.html

vanity.dot: vanity.yml $(DEPS)
	../bin/vanity -t dot < vanity.yml > vanity.dot

# Render graph. Needs 'dot' program, normally from a 'graphviz' package.
vanity.png: vanity.dot $(DEPS)
	dot -Tpng vanity.dot -o vanity.png
