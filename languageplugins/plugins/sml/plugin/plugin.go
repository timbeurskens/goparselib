package main

import (
	"github.com/timbeurskens/goparselib"
	"github.com/timbeurskens/goparselib/languageplugins/plugins/sml"
)

func Load() goparselib.Plugin {
	return goparselib.MakePlugin(goparselib.R(sml.Root), nil)
}
