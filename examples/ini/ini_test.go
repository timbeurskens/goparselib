package ini

import (
	"log"
	"testing"

	"goparselib"
)

func TestBasic(t *testing.T) {
	n, err := goparselib.ParseString(basic1, Root)
	if err != nil {
		t.Error(err)
	}
	n.Populate(basic1)
	n2, err := n.Reduce(goparselib.Blank)
	if err != nil {
		t.Error(err)
	}
	n2.Output(log.Writer())
}

func TestEmpty(t *testing.T) {
	n, err := goparselib.ParseString(empty, Root)
	if err != nil {
		t.Error(err)
	}
	n.Populate(empty)
	n2, err := n.Reduce(goparselib.Blank)
	if err != nil {
		t.Error(err)
	}
	n2.Output(log.Writer())
}

func TestTools(t *testing.T) {
	reader, err := New(`
hello = world
test = 10

[menu]
size = 4
#
`)
	if err != nil {
		t.Error(err)
	}

	t.Log(reader)

	t.Log(reader.Property("hello"))

	s, err := reader.Section("menu")
	if err != nil {
		t.Error(err)
	}

	t.Log(s.Property("size"))
}

func TestFile(t *testing.T) {
	tree, err := goparselib.ParseFile("test.ini", Root)
	if err != nil {
		t.Error(err)
	}
	t.Log(tree)
}

const (
	basic1 = `
test = 22

[h]
age = hoi

[a]
a=1
#\x0`

	empty = `#`
)
