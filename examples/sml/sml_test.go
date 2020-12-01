package sml

import (
	"log"
	"testing"

	"goparselib"
)

func TestBasic(t *testing.T) {
	t.Log(goparselib.Parse(line, timelineLine))
	t.Log(goparselib.Parse(eolExample, eol))

	tree, err := goparselib.Parse(basic, goparselib.R(root))
	if err != nil {
		t.Error(err)
	}

	tree.Populate(basic)

	tree.Output(log.Writer())
	t.Log(tree)

	r, _ := tree.Reduce(eol, eof, space, lopen, rclose, enumEnd, forLit, play, programLit, timelineLit, nil)
	r.Output(log.Writer())
}

const (
	eolExample = ` 
`
	basic = `timeline {
 1) play rainbow for 10 seconds
 2) play test for 20 milliseconds
}#`
	line = `1) play rainbow for 10 seconds
`

	example1 = `

program rainbow {
}

timeline {
  1) play rainbow for 10 seconds
  2) wait for 5 seconds
  3) jump to step 2
}
#
`
)
