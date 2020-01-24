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
	termId := getTermId(term)
	fmt.Printf(`
<p class="ad-def" name="ad-%s">
  <span class="ad-term">%s</span>:
  <span class="ad-contents">%s</span>
</p>
`,
		html.EscapeString(termId),
		html.EscapeString(term),
		"",
	)
}

func outputHtml(doc []Definition) {
	for _, def := range doc {
		outputHtmlDef(def)
	}
}
