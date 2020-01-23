package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type YamlDefinition struct {
	Term string // the term being defined
	Contents string `yaml:"def,flow"` // the definition using special markup
	Synonyms []string `yaml:"syn,omitempty"` // synonyms for the term
}

// A document is an ordered list of definitions. Since Yaml doesn't guarantee
// the order of fields within a map, the document is a list of definitions
// rather than a map from terms to definitions.
type YamlDocument []YamlDefinition

type DefContents []DefContentsElt

type EltKind int

const (
	Defined EltKind = iota
	Undefined
)

type DefContentsElt struct {
	Elt string
	EltKind EltKind
}

type Definition struct {
	Term string
	Contents DefContents
	Synonyms []string
}

type Dictionary map[string]Definition

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

func loadYamlFile(data string) (doc *YamlDocument) {
	err = yaml.Unmarshal([]byte(data), &doc)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return doc
}

func parseDefContents(data string) (defContents *DefContents) {
	lex := newLexer([]byte(data))
	err := yyParse(lex)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return doc
}

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//doc := loadYamlFile(data)
	doc := parseDefContents(`hello [world]`)
	fmt.Printf("%v\n", doc)
}
