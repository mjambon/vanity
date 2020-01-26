Directed Acyclic Dictionary
==

This project is a one-off proof-of-concept. Its goal is to evaluate
the usefulness of dictionaries in which terms used in definitions
must either have been already defined or appear in a distinctive
font.

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

Suggested features
--

Public awareness:
* Use the tool and publish reports on its usefulness and lessons learned.

Maintenance and distribution:
* Compile and distribute binaries for the most popular platforms. Note
  that Go makes this easy by producing static binaries and
  cross-compiling.
* Add automatic testing using one of Travis, CircleCI, Github Actions,
  etc.
* Add contribution guidelines (highly recommended to do before
  accepting contributions).

User-facing features:
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
  demand. [Pandoc](https://pandoc.org/ is) an excellent tool for
  this. Part of the work would consist in making the original output
  of ad fully understood by pandoc. Perhaps the best format for
  this isn't HTML but some other language best suited for pandoc input.
