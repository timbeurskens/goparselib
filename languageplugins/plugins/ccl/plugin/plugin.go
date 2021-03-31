package main

import (
	"github.com/timbeurskens/goparselib"
	"github.com/timbeurskens/goparselib/languageplugins/plugins/ccl"
)

func Load() goparselib.Symbol {
	return ccl.Root
}
