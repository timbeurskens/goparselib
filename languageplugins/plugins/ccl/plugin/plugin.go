package main

import (
	"github.com/timbeurskens/goparselib"
	"github.com/timbeurskens/goparselib/languageplugins/plugins/ccl"
)

func Load() goparselib.Plugin {
	return goparselib.MakePlugin(ccl.Root, nil)
}
