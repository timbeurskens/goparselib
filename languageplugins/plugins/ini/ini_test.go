package ini

import (
	"github.com/timbeurskens/goparselib/parser"
	"log"
	"testing"

	"github.com/timbeurskens/goparselib"
)

func New(input string) (Reader, error) {
	tree, err := parser.ParseString(input, Root)
	if err != nil {
		return Reader{}, err
	}
	ast, err := tree.Reduce(nil, parOpen, parClose, eq, goparselib.Blank, goparselib.EOL, goparselib.BlankOpt)
	if err != nil {
		return Reader{}, err
	}

	return Reader{
		Model: ast,
	}, nil
}

func TestBasic(t *testing.T) {
	n, err := parser.ParseString(basic1, Root)
	if err != nil {
		t.Error(err)
	}
	n2, err := n.Reduce(goparselib.Blank)
	if err != nil {
		t.Error(err)
	}
	n2.Output(log.Writer())
}

func TestEmpty(t *testing.T) {
	n, err := parser.ParseString(empty, Root)
	if err != nil {
		t.Error(err)
	}
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
	tree, err := parser.ParseFile("test.ini", Root)
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