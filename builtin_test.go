package goparselib

import "testing"

func TestDecorate(t *testing.T) {
	decorator := CTerminal("!")
	container := Concat{CTerminal("H"), CTerminal("e"), CTerminal("l"), CTerminal("l"), CTerminal("o")}

	result := Decorate(container, decorator)

	t.Log([]Symbol(result.(Concat)))

	if (len(container)*2 + 1) != len(result.(Concat)) {
		t.Error("lengths do not match", len(container), len(result.(Concat)))
	}

	for i := 0; i < len(result.(Concat)); i += 2 {
		if result.(Concat)[i] != decorator {
			t.Error("Expected decorator, got", result.(Concat)[i])
		}
	}
}
