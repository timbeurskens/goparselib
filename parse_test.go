package goparselib

import "testing"

func TestLang(t *testing.T) {
	_, err1 := Parse(example1, R(root))
	if err1 != nil {
		t.Error(err1)
	}
	_, err2 := Parse(example2, R(root))
	if err2 != nil {
		t.Error(err2)
	}
	node, err3 := Parse(example3, R(root))
	if err3 != nil {
		t.Error(err3)
	}

	_, nok1 := Parse(counterexample1, R(root))
	if nok1 == nil {
		t.Fail()
	}
	t.Log(nok1)
	_, nok2 := Parse(counterexample2, R(root))
	if nok2 == nil {
		t.Fail()
	}
	t.Log(nok2)
	_, nok3 := Parse(counterexample3, R(root))
	if nok3 == nil {
		t.Fail()
	}
	t.Log(nok3)
	node.Populate(example3)

	t.Log(node)
}
