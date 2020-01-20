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
it accomodates text better than JSON and is more readable than
XML. Structured data such as lists of synonyms can easily be added
without extending the syntax. The only originality is in the markup
language used in the body of the definitions, which uses its own
conventions to link terms to their definition.

The Go language was chosen for this implementation as it's relatively
friendly to external contributors, and it was a good opportunity for
the author to learn it.

Unimplemented ideas
--

...
