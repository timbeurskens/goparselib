package esystems

import (
	. "github.com/timbeurskens/goparselib"
)

var (
	Root   = List(Concat{}, LF, LF)
	Layout = Root
)
