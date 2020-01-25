// A program to check and format an acyclic dictionary.
// See example.yml for sample input.

package main

import (
	"io/ioutil"
	"log"
	"os"
)

type EltKind int

const (
	DefinedTerm EltKind = iota
	Text
)

type DefContentsElt struct {
	Kind EltKind
	Text string
}

type Definition struct {
	Term string
	Contents []DefContentsElt
	Synonyms []string
}

type Dictionary struct {
	// A lookup table, needed to resolve synonyms into the canonical term,
	// for linking to a definition.
	Map map[string]Definition

	// The sequence of definitions in order specified in the original input
	// file.
	Sequence []Definition
}

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	doc := loadData(string(data))
	outputHtmlPage(doc)
}
