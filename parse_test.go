package goparselib

import "testing"

func TestLang(t *testing.T) {
    ok1, _ := Parse(example1, R(root))
    ok2, _ := Parse(example2, R(root))
    ok3, node := Parse(example3, R(root))

    if !(ok1 && ok2 && ok3) {
        t.Error("not passing")
    }

    nok1, _ := Parse(counterexample1, R(root))
    nok2, _ := Parse(counterexample2, R(root))
    nok3, _ := Parse(counterexample3, R(root))

    if nok1 || nok2 || nok3 {
        t.Error("passing")
    }

    node.Populate(example3)

    t.Log(node)
}
