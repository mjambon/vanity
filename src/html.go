// Format a dictionary to stylable HTML

package main

import (
	"fmt"
	"html"
	"encoding/hex"
)

// Return hex-encoded ID which is URL- and HTML-friendly.
func getTermId(term string) string {
	return hex.EncodeToString([]byte(term))
}

func outputHtmlDef(def Definition) {
	term := def.Term
	fmt.Printf(`
<p class="ad-def">
  <a name="ad-%s"><strong class="ad-term">%s</strong></a>:
  <span class="ad-contents">`,
		html.EscapeString(getTermId(term)),
		html.EscapeString(term),
	)
	for _, elt := range def.Contents {
		if elt.Kind == DefinedTerm {
			term := elt.Text
			fmt.Printf(`<a href="#%s" class="ad-term-link">%s</a>`,
				html.EscapeString(getTermId(term)),
				html.EscapeString(term),
			)
		} else {
			fmt.Printf(`%s`, html.EscapeString(elt.Text))
		}
	}
	fmt.Printf(`</span>
</p>
`)
}

// Print HTML to be included in a document to stdout.
func outputHtml(doc []Definition) {
	for _, def := range doc {
		outputHtmlDef(def)
	}
}

// Print a single HTML page with basic styling.
func outputHtmlPage(doc []Definition) {
	fmt.Printf(`<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>definitions</title>
</head>
<body>
`)
	outputHtml(doc)
	fmt.Printf(`
</body>
`)
}
