// A command-line program to check and format an acyclic dictionary.
// This is the entrypoint.
//
// See /example/example.yml for sample input.

package main

import (
	"io/ioutil"
	"log"
	"os"

	// using this rather than the built-in "flag" package because it supports
	// --long options.
	"github.com/jessevdk/go-flags"
)

// Type definitions used in the document AST. They could as well go in a
// file of their own.

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

// Command-line parsing
//
// Whenever possible, we use the same options as those offered by 'pandoc',
// a document translation tool.
//

type Options struct {
	// TODO: find a way to make these strings readable and cut at 80 columns.
	OutputFormat string `short:"t" long:"to" default:"html" description:"Specify output format. It can be one of: html (HTML snippet or standalone page), dot (dot format supported by Graphviz)."`

	Standalone bool `short:"s" long:"standalone" description:"Produce a standalone HTML document."`

	Title string `long:"title" default:"Definitions" description:"The document title. Applies to standalone HTML output only."`

	// It would be nice if the following options could be repeated so as to
	// inject multiple files in one command, like the same options do in pandoc.
	IncludeInHeader string `short:"H" long:"include-in-header" description:"Include contents of the given file, verbatim, at the end of the HTML <head> section. This is meant for adding CSS styling or Javascript. Applies to standalone HTML output only."`
	IncludeBeforeBody string `short:"B" long:"include-before-body" description:"Include contents of the given file, verbatim, at the beginning of the HTML <body> section. This is meant for adding introductory material at the beginning of the page. Applies to standalone HTML output only."`
	IncludeAfterBody string `short:"A" long:"include-after-body" description:"Include contents of the given file, verbatim, at the end of the HTML <body> section. This is meant for adding concluding material at the bottom of the page. Applies to standalone HTML output only."`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)

func main() {
	_, err := parser.Parse()
	if err != nil {
		flagsErr, ok := err.(*flags.Error)
		if ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			log.Fatalf("error: %v", err)
		}
	}

	// Load data from stdin because it's simpler. Could read from file as well
	// if that's useful.
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	doc := loadData(string(data))

	switch options.OutputFormat {
	case "html":
		if options.Standalone {
			outputHtmlPage(doc, options)
		} else {
			outputHtml(doc)
		}
	case "dot":
		outputDot(doc)
	}
}
