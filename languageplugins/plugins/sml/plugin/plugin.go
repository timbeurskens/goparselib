package main

import (
	"github.com/timbeurskens/goparselib"
	"github.com/timbeurskens/goparselib/languageplugins/plugins/sml"
)

func Load() goparselib.Symbol {
	return goparselib.R(sml.Root)
}
