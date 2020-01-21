package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
//	"github.com/biogo/ragel"
)

//go:generate ragel -Z -G2 -o lexer.go lexer.rl
//go:generate goyacc parser.y

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

func loadYamlFile() (doc *YamlDocument) {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = yaml.Unmarshal([]byte(data), &doc)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return doc
}

func main() {
	doc := loadYamlFile()
	fmt.Printf("%v\n", doc)
}
