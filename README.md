Constructive Dictionaries
==

This project is a proof-of-concept command-line tool for building and
presenting dictionaries as a strict hierarchy of terms.

The goal is to produce precise technical and philosophical dictionaries and
glossaries without circular definitions. The solution is to highlight
and link technical terms, and prevent references to terms that have not
yet been defined.

Here's an example of a glossary rendered with `vanity` into an HTML page:
![Example](screenshot.png)

The source for this glossary is:

```yaml
# A sample glossary.
---
- term: furniture
  def: >
    collective noun designating all sorts of large but movable items used in
    indoor settings with a primary non-decorative function.
  syn:
    - piece of furniture
- term: table
  def: >
    a [piece of furniture] with a horizontal surface lying above the floor
    that was designed to support smaller objects.
  syn:
    - tables
- term: seat
  def: >
    a [piece of furniture] designed for a single person to sit on it.
  syn:
    - seats
- term: chair
  def: >
    a mobile [seat] with a backrest.
  syn:
    - chairs
- term: stool
  def: >
    a [seat] without a backrest.
  syn:
    - stools
- term: bench
  def: >
    an item designed for several people to sit on, with a hard surface.
  syn:
    - benches
- term: couch
  def: >
    a [piece of furniture] designed for multiple people to sit on, with
    a cushy surface.
  syn:
    - couches
```

Implementation
--

This proof-of-concept is implemented as a command-line program that
reads a dictionary in source form, checks its validity, and produces a
readable document.

The input is a list of term definitions. The YAML syntax was chosen as
it accommodates text better than JSON and is more readable than
XML. Structured data such as lists of synonyms can easily be added
without extending the syntax. The only originality is in the markup
language used in the body of the definitions, which uses its own
conventions to link terms to their definition.

The Go language was chosen for this implementation as it's relatively
friendly to external contributors, and it was a good opportunity for
the author to learn it.

Installation
--

[Download](https://github.com/mjambon/vanity/releases) a
statically-linked executable for your platform.

If your platform is not in there or if you want to try a development
version, you'll have to build it from source. We have some
instructions in [DEV.md](DEV.md). After installing the prerequisites,
you can build and install `vanity` for your Unixy platform using
`make && make install`.

Documentation
--

Once installed, check out the output of `vanity --help`.
The following output formats are supported:

* HTML snippet or standalone page
* graph in the dot format understood by Graphviz

Input format reference
--

A valid input document is a [yaml](https://yaml.org/) document,
consisting of an ordered list (yaml array) of definitions. Each
definition has 2 mandatory fields `term` and `def`, and an optional
field `syn`.

* `term`: string that represents the standard form of the term being
  defined.
* `def`: string that holds the definition for the term. Links to other
  terms are placed within square brackets, such as in
  `The sky is [cloudy] today.`. Only links to previous definitions
  or to the current definition are permitted. The text of the link must
  be a term from a `term` field or one of its synonyms from a `syn`
  field.
* `syn`: array of strings that are considered synonyms with the term
  being defined, in the sense that any reference to a synonym will
  link to the standard term. This can be used to hold different
  variations of a word, such as plural forms, gendered forms, plain
  synonyms, conjugated forms of verbs, etc. (this will become annoying
  for some languages other than English)

The definitions must be sorted such that all links refer to terms that
were defined earlier. For example, the following is valid input:

```yaml
- term: potato
  def: the edible tuber from the potato plant
- term: French fries
  def: deep-fried [potato] chunks
```

We can add `potatoes` as a synonym of `potato` and link to `potatoes`
instead of `potato`:

```yaml
- term: potato
  def: the edible tuber from the potato plant
  syn:
    - potatoes
- term: French fries
  def: deep-fried [potatoes]
```

Order matters. The following is illegal:

```yaml
- term: French fries
  def: deep-fried [potato] chunks
- term: potato
  def: the edible tuber from the potato plant
```

In that case we get an error message:
```sh
$ vanity < glossary.yml > glossary.html
error: definition for term 'French fries' uses undefined term: 'potato'.
```

Suggested features
--

Public awareness:
* Write an introductory article explaining why and when this thing can
  be useful.
* Use the tool and publish reports on its usefulness and lessons
  learned.

Maintenance and distribution:
* Add automatic testing using one of Travis, CircleCI, Github Actions,
  etc.
* Add contribution guidelines (highly recommended to do before
  accepting contributions).

User-facing features:
* Document the input format.
* Produce a graph even if it has cycles, as an aid to see what's going
  on.
* Use topological sorting to implement some of the following features:
  - Rearrange the input document in an order compatible with the
    dependencies. This is a conversion from yaml to yaml.
  - Automatically sort the input document topologically so that the
    author doesn't have to. This is only for checking purposes.
    This doesn't produce new input or different output.
  - Sort the definitions in the output document in dependency order
    or reverse dependency order, depending on user preference.
* Add an option to sort the terms alphabetically.
* Add support for multiple senses via some dedicated syntax. It could
  be something like `something_2` where the term is identified by the
  full string `something_2` but rendered as just `something` or
  `something (2)`, and links to the correct definition.
* Offer out-of-the-box option for showing definition preview on hover
  or single-tap on mobile. This would work like the Wikipedia mobile
  app or [Wikiwand](https://www.wikiwand.com/en/Hippopotamus).
* Export to PDF or whichever format is in
  demand. [Pandoc](https://pandoc.org/) is an excellent tool for
  this. Part of the work would consist in making the original output
  of vanity fully understood by pandoc. Perhaps the best format for
  this isn't HTML but some other language best suited for pandoc input.
