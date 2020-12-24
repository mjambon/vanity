// Output a graph of terms in the dot format for Graphviz.
//
// dot supports various rendering options. For now we just hardcode our own
// style.
//
// https://www.graphviz.org/doc/info/lang.html is the language reference.

package main

import (
	"fmt"
)

func outputDotHead() {
	fmt.Printf(
`digraph G {
  rankdir = TB;
`)
}

func outputDotTail() {
	fmt.Printf("}\n")
}

// dot double-quoted literals only escape double quotes with a backslash.
func quoteDotString(s string) string {
	buf := []byte{'"'}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '"' {
			buf = append(buf, '\\', '"')
		} else {
			buf = append(buf, c)
		}
	}
	buf = append(buf, '"')
	return string(buf)
}

func outputDotDef(defs map[string]Definition, def Definition) {
	srcTerm := def.Term
	srcNormalizedTerm := def.NormalizedTerm
	fmt.Printf(
		"  %s [label=%s];\n",
		quoteDotString(srcNormalizedTerm),
		quoteDotString(srcTerm),
	);
	edges := make(map[string]bool)  // for detecting duplicate outgoing edges
	for _, elt := range def.Contents {
		if elt.Kind == DefinedTerm {
			dstDef := defs[elt.NormalizedText]
			dstNormalizedTerm := dstDef.NormalizedTerm
			if dstNormalizedTerm != srcNormalizedTerm && !edges[dstNormalizedTerm] {
				edges[dstNormalizedTerm] = true
				fmt.Printf("  %s -> %s;\n",
					quoteDotString(srcNormalizedTerm),
					quoteDotString(dstNormalizedTerm),
				)
			}
		}
	}
}

func outputDot(doc Dictionary, options Options) {
	outputDotHead()
	fmt.Printf("%s\n", readFile(options.IncludeBeforeBody));
	for _, def := range doc.Sequence {
		outputDotDef(doc.Map, def)
	}
	fmt.Printf("%s\n", readFile(options.IncludeAfterBody));
	outputDotTail()
}
