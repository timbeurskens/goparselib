package sml

import (
	"log"
	"testing"

	"goparselib"
)

func TestBasic(t *testing.T) {
	t.Log(goparselib.Parse(line, timelineLine))
	t.Log(goparselib.Parse(eolExample, eol))

	ok, tree := goparselib.Parse(basic, goparselib.R(root))
	if !ok {
		t.Error("parse failure")
	}

	tree.Populate(basic)

	tree.Output(log.Writer())
	t.Log(tree)

	r, _ := tree.Reduce(eol, eof, space, lopen, rclose, enumEnd, forLit, play, programLit, timelineLit, nil)
	r.Output(log.Writer())
}
