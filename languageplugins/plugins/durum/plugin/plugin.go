package main

import (
	"github.com/timbeurskens/goparselib"
	"github.com/timbeurskens/goparselib/languageplugins/plugins/durum"
)

// Load is the initial function
func Load() goparselib.Plugin {
	return goparselib.MakePlugin(durum.Root, durum.Layout)
}
