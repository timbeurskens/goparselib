package ini

import (
	"log"
	"testing"

	"goparselib"
)

func TestBasic(t *testing.T) {
	n, err := goparselib.Parse(basic1, root)
	if err != nil {
		t.Error(err)
	}
	n.Populate(basic1)
	n2, err := n.Reduce(goparselib.Blank, eof)
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

const (
	basic1 = `
test = 22

[h]
age = hoi

[a]
a=1
#`
)
