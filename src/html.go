// Format a dictionary to stylable HTML

package main

import (
	"encoding/hex"
	"fmt"
	"html"
	"io/ioutil"
	"log"
)

// Return hex-encoded ID which is URL- and HTML-friendly.
func getTermId(term string) string {
	return hex.EncodeToString([]byte(term))
}

func outputHtmlDef(defs map[string]Definition, def Definition) {
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
			fmt.Printf(`<a href="#ad-%s" class="ad-term-link">%s</a>`,
				html.EscapeString(getTermId(defs[term].Term)),
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
func outputHtml(doc Dictionary) {
	for _, def := range doc.Sequence {
		outputHtmlDef(doc.Map, def)
	}
}

func readFile(path string) string {
	if len(path) != 0 {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		return string(data)
	} else {
		return ""
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

	outputHtml(doc)

	fmt.Printf(`
%s</body>
`,
		readFile(options.IncludeAfterBody),
	)
}
