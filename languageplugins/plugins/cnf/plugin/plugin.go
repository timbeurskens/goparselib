package main

import (
	"github.com/timbeurskens/goparselib"
	"github.com/timbeurskens/goparselib/languageplugins/plugins/cnf"
)

// Load is the initial function
func Load() goparselib.Plugin {
	return goparselib.MakePlugin(cnf.Root, nil)
}
