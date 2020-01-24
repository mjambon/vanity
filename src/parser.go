// The parser for the body of definitions

package main

import (
	"regexp"
)

// An ugly and simple parser to extract terms enclosed within brackets.
// There is no way to express literal square brackets.
//
// For example, "[polar] bear!" is parsed into two elements:
// - "polar": a defined term
// - " bear!": some text
//
// The result is an array of those elements.
//
func parseDefContents(s string) []DefContentsElt {
	re, _ := regexp.Compile("\\[[^\\]]*\\]|[^\\]]+")
	tokens := re.FindAllString(s, -1) // wth is the 2nd argument?
	res := []DefContentsElt{}
	for _, token := range tokens {
		if token[0] == '[' {
			elt := DefContentsElt{
				EltKind: DefinedTerm,
				Elt: token[1:len(token)-1],
			}
			res = append(res, elt)
		} else {
			elt := DefContentsElt{
				EltKind: Text,
				Elt: token,
			}
			res = append(res, elt)
		}
	}
	return res
}
