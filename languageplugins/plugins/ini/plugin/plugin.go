package main

import (
	"github.com/timbeurskens/goparselib"
	"github.com/timbeurskens/goparselib/languageplugins/plugins/ini"
)

func Load() goparselib.Symbol {
	return ini.Root
}
