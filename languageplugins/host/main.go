package main

import (
	"flag"
	"github.com/timbeurskens/goparselib"
	"github.com/timbeurskens/goparselib/parser"
	"log"
	"plugin"
)

var (
	pluginFlag = flag.String("plugin", "", "The language plugin to use")
	infileFlag = flag.String("input", "", "The input file to parse")
)

func main() {
	flag.Parse()

	var languagePlugin *plugin.Plugin
	var loadSymbol plugin.Symbol
	var parseResult goparselib.Node
	var err error

	if languagePlugin, err = plugin.Open(*pluginFlag); err != nil {
		log.Fatal(err)
	}

	if loadSymbol, err = languagePlugin.Lookup("Load"); err != nil {
		log.Fatal(err)
	}

	loadFunc, ok := loadSymbol.(goparselib.LoadFunction)
	if !ok {
		log.Fatal("The load symbol is not a load function")
	}

	languageRoot := loadFunc()

	if parseResult, err = parser.ParseFile(*infileFlag, languageRoot.Root()); err != nil {
		log.Fatal(err)
	}

	if languageRoot.Layout() != nil {
		parseResult, err = parseResult.Reduce(languageRoot.Layout()...)
	}

	log.Println(parseResult)
}
