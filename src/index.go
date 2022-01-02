// Show a list of all the dictionary terms in a compact format and
// sorted alphabetically.

package main

import (
	"fmt"
	"sort"
)

// Creating a custom type is necessary for sorting.
// (based on https://gobyexample.com/sorting-by-functions)
type AlphaDefs []Definition

func (a AlphaDefs) Len() int {
	return len(a)
}

func (a AlphaDefs) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a AlphaDefs) Less(i, j int) bool {
	return a[i].NormalizedTerm < a[j].NormalizedTerm
}

// Print sorted list of terms linking to their definition
func outputIndex(doc Dictionary) {
	src := doc.Sequence
	sorted := make([]Definition, len(src))
	copy(sorted, src)
	sort.Sort(AlphaDefs(sorted))
	fmt.Printf(`<p class="vanity-index">`)
	for i, def := range sorted {
		outputHtmlTermLink(def.Term, def.NormalizedTerm)
		if i < len(sorted) - 1 {
			fmt.Printf(`, `)
		}
	}
	fmt.Printf(`</p>`)
}
