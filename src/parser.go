// The parser for the body of definitions

package main

import (
	"log"
	"gopkg.in/yaml.v2"
	"regexp"
)

type rawDefinition struct {
	Term string // the term being defined
	Contents string `yaml:"def,flow"` // the definition using special markup
	Synonyms []string `yaml:"syn,omitempty"` // synonyms for the term
}

// A document is an ordered list of definitions. Since Yaml doesn't guarantee
// the order of fields within a map, the document is a list of definitions
// rather than a map from terms to definitions.
type rawDocument []rawDefinition

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
	re, _ := regexp.Compile("\\[[^\\]]*\\]|[^\\[]+")
	tokens := re.FindAllString(s, -1) // wth is the 2nd argument?
	res := []DefContentsElt{}
	for _, token := range tokens {
		if token[0] == '[' {
			elt := DefContentsElt{
				Kind: DefinedTerm,
				Text: token[1:len(token)-1],
			}
			res = append(res, elt)
		} else {
			elt := DefContentsElt{
				Kind: Text,
				Text: token,
			}
			res = append(res, elt)
		}
	}
	return res
}

func loadYamlData(data string) (doc rawDocument) {
	err := yaml.Unmarshal([]byte(data), &doc)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return doc
}

func checkDef(defs map[string]Definition, def Definition) {
	for _, elt := range def.Contents {
		if elt.Kind == DefinedTerm {
			_, ok := defs[elt.Text]
			if !ok {
				log.Fatalf(
					"error: definition for term '%s' uses undefined term: '%s'.",
					def.Term,
					elt.Text,
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
