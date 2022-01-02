// Format a dictionary to stylable HTML

package main

import (
	"encoding/hex"
	"fmt"
	"html"
)

// Return hex-encoded ID which is URL- and HTML-friendly.
func getTermId(term string) string {
	return hex.EncodeToString([]byte(term))
}

func outputHtmlTermLink(linkText string, normalizedTerm string) {
	fmt.Printf(`<a href="#vanity-%s" class="vanity-term-link">%s</a>`,
		html.EscapeString(getTermId(normalizedTerm)),
		html.EscapeString(linkText),
	)
}

func outputHtmlDef(defs map[string]Definition, def Definition) {
	term := def.Term
	fmt.Printf(`
<p class="vanity-def">
  <a name="vanity-%s"></a><strong class="vanity-term">%s</strong>:
  <span class="vanity-contents">`,
		html.EscapeString(getTermId(term)),
		html.EscapeString(term),
	)
	for _, elt := range def.Contents {
		if elt.Kind == DefinedTerm {
			outputHtmlTermLink(elt.Text, defs[elt.NormalizedText].Term)
		} else {
			fmt.Printf(`%s`, html.EscapeString(elt.Text))
		}
	}
	fmt.Printf(`
  </span>
</p>
`)
	image := def.Image
	if image != "" {
		esc_img := html.EscapeString(image)
		fmt.Printf(`
  <div class="vanity-image-div">
    <a href="img/%s"><img src="img/%s" class="vanity-image-img"/></a>
  </div>
`,
			esc_img, esc_img)
	}
}

// Print HTML to be included in a document to stdout.
func outputHtml(doc Dictionary, options Options) {
	if options.IndexPlacement == "before" {
		outputIndex(doc)
	}
	for _, def := range doc.Sequence {
		outputHtmlDef(doc.Map, def)
	}
	if options.IndexPlacement == "after" {
		outputIndex(doc)
	}
}

// Print a single HTML page with basic styling.
func outputHtmlPage(doc Dictionary, options Options) {
	escapedTitle := html.EscapeString(options.Title)
	fmt.Printf(`<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>%s</title>
%s</head>
<body>
<h1>%s</h1>%s
`,
		escapedTitle,
		readFile(options.IncludeInHeader),
		escapedTitle,
		readFile(options.IncludeBeforeBody),
	)

	outputHtml(doc, options)

	fmt.Printf(`
%s</body>
</html>
`,
		readFile(options.IncludeAfterBody),
	)
}
