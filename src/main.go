// A program to check and format an acyclic dictionary.

package main

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type RawDefinition struct {
	Term string // the term being defined
	Contents string `yaml:"def,flow"` // the definition using special markup
	Synonyms []string `yaml:"syn,omitempty"` // synonyms for the term
}

// A document is an ordered list of definitions. Since Yaml doesn't guarantee
// the order of fields within a map, the document is a list of definitions
// rather than a map from terms to definitions.
type RawDocument []RawDefinition

type DefContents []DefContentsElt

type EltKind int

const (
	DefinedTerm EltKind = iota
	Text
)

type DefContentsElt struct {
	EltKind EltKind
	Elt string
}

type Definition struct {
	Term string
	Contents DefContents
	Synonyms []string
}

type Dictionary []Definition

var example = `
---
- term: body
  def: >
    the physical part necessary and sufficient for an individual to
    function.
- term: head
  def: >
    the part of the [body] of animals containing the brain.
  syn:
    - heads
`

func loadYamlData(data string) (doc RawDocument) {
	err := yaml.Unmarshal([]byte(data), &doc)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return doc
}

func checkDef(defs map[string]Definition, def Definition) {
	for _, elt := range def.Contents {
		if elt.EltKind == DefinedTerm {
			_, ok := defs[elt.Elt]
			if ok {
				log.Fatalf("error: definition for term %s uses undefined term: %s.",
					def.Term,
					elt.Elt,
				)
			}
		}
	}
}

func loadData(data string) (doc []Definition) {
	rawDoc := loadYamlData(data)
	defs := make(map[string]Definition)
	for _, rawDef := range rawDoc {
		contents := parseDefContents(rawDef.Contents)
		term := rawDef.Term
		def := Definition{
			Term: term,
			Contents: contents,
			Synonyms: rawDef.Synonyms,
		}
		_, exists := defs[term]
		if exists {
			log.Fatalf("error: duplicate definition for term %s", term)
		}
		// allow term in its own definition
		defs[term] = def
		checkDef(defs, def)
		doc = append(doc, def)
	}
	return doc
}

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	doc := loadData(string(data))
	outputHtml(doc)
}
