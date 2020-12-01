package goparselib

import "testing"

func TestLang(t *testing.T) {
	_, err1 := ParseString(example1, R(root))
	if err1 != nil {
		t.Error(err1)
	}
	_, err2 := ParseString(example2, R(root))
	if err2 != nil {
		t.Error(err2)
	}
	node, err3 := ParseString(example3, R(root))
	if err3 != nil {
		t.Error(err3)
	}

	_, nok1 := ParseString(counterexample1, R(root))
	if nok1 == nil {
		t.Fail()
	}
	t.Log(nok1)
	_, nok2 := ParseString(counterexample2, R(root))
	if nok2 == nil {
		t.Fail()
	}
	t.Log(nok2)
	_, nok3 := ParseString(counterexample3, R(root))
	if nok3 == nil {
		t.Fail()
	}
	t.Log(nok3)
	node.Populate(example3)

	t.Log(node)
}
