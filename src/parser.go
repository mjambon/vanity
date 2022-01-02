// The parser for the body of definitions

package main

import (
	"log"
	"gopkg.in/yaml.v3"
	"regexp"
	"strings"
)

type rawDefinition struct {
	Term string // the term being defined
	Contents string `yaml:"def,flow"` // the definition using special markup
	Synonyms []string `yaml:"syn,omitempty"` // synonyms for the term
	Image string `yaml:"img,omitempty"` // image file name matching [a-z0-9_.-]+
}

// A document is an ordered list of definitions. Since Yaml doesn't guarantee
// the order of fields within a map, the document is a list of definitions
// rather than a map from terms to definitions.
type rawDocument []rawDefinition

func normalizeText(s string) string {
	return strings.ToLower(s)
}

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
			text := token[1:len(token)-1]
			elt := DefContentsElt{
				Kind: DefinedTerm,
				Text: text,
				NormalizedText: normalizeText(text),
			}
			res = append(res, elt)
		} else {
			elt := DefContentsElt{
				Kind: Text,
				Text: token,
				NormalizedText: normalizeText(token),
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
			_, ok := defs[elt.NormalizedText]
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

func checkDuplicates(
	term string,
	defs map[string]Definition,
	synonyms []string) {

	_, exists := defs[term]
	if exists {
		log.Fatalf("error: duplicate definition for term '%s'", term)
	}
	for _, syn := range synonyms {
		_, exists := defs[syn]
		if exists {
			log.Fatalf("error: synonym '%s' of term '%s' is duplicate", syn, term)
		}
	}
}

func validateImageName(name string) {
	matched, err := regexp.MatchString("[a-z0-9._-]+", name)
	if err != nil {
		log.Fatalf("error: cannot validate image name %v: %v", name, err)
	} else {
		if !matched {
			log.Fatalf("error: invalid image name %v", name)
		}
	}
}

func loadData(data string) (doc Dictionary) {
	seq := []Definition{}
	rawDoc := loadYamlData(data)
	defs := make(map[string]Definition)
	for _, rawDef := range rawDoc {
		contents := parseDefContents(rawDef.Contents)
		term := rawDef.Term
		image := rawDef.Image
		if image != "" {
			validateImageName(image)
		}
		normalizedTerm := normalizeText(term)
		normalizedSynonyms := make([]string, len(rawDef.Synonyms))
		for i, orig := range rawDef.Synonyms {
			normalizedSynonyms[i] = normalizeText(orig)
		}
		def := Definition{
			Term: term,
			NormalizedTerm: normalizedTerm,
			Contents: contents,
			NormalizedSynonyms: normalizedSynonyms,
			Image: image,
		}
		checkDuplicates(term, defs, def.NormalizedSynonyms)

		// allow term (and its synonyms) in its own definition
		defs[normalizedTerm] = def
		for _, syn := range def.NormalizedSynonyms {
			defs[syn] = def
		}

		checkDef(defs, def)
		seq = append(seq, def)
	}
	doc = Dictionary{
		Sequence: seq,
		Map: defs,
	}
	return doc
}
