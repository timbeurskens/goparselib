package ccl

import (
	"github.com/timbeurskens/goparselib/parser"
	"strings"
	"testing"

	"github.com/timbeurskens/goparselib"
)

func TestSimple(t *testing.T) {
	tree, err := parser.ParseString(simple, Root)
	if err != nil {
		t.Error(err)
	}
	t.Log(tree)
}

func TestCompleteFile(t *testing.T) {
	tree, err := parser.ParseFile("test.ccl", Root)
	if err != nil {
		t.Error(err)
	}
	t.Log(tree)
}

func TestComplete(t *testing.T) {
	tree, err := parser.ParseString(complete, Root)
	if err != nil {
		t.Error(err)
	}
	t.Log(tree)

	reduced, err := tree.Reduce(goparselib.Blank, goparselib.BlankOpt, nil)
	if err != nil {
		t.Error(err)
	}

	writer := &strings.Builder{}
	err = reduced.Output(writer)
	if err != nil {
		t.Error(err)
	}

	t.Log(writer.String())
}

func TestStorage(t *testing.T) {
	tree, err := parser.ParseString(storageStr, storage)
	if err != nil {
		t.Error(err)
	}
	t.Log(tree)
}

func TestEmptyBody(t *testing.T) {
	tree, err := parser.ParseString(emptyBody, storageBody)
	if err != nil {
		t.Error(err)
	}
	t.Log(tree)
}

func TestRegionBody(t *testing.T) {
	tree, err := parser.ParseString(regionStr, storageBody)
	if err != nil {
		t.Error(err)
	}
	t.Log(tree)
}

func TestList(t *testing.T) {
	tree, err := parser.ParseString(listStr, goparselib.List(goparselib.Ident, goparselib.Comma, goparselib.Blank))
	if err != nil {
		t.Error(err)
	}
	t.Log(tree)
}

func TestEmptyList(t *testing.T) {
	tree, err := parser.ParseString(emptyBody, goparselib.List(goparselib.Ident, goparselib.Comma, goparselib.Blank))
	if err != nil {
		t.Error(err)
	}
	t.Log(tree)
}

const (
	emptyBody = `
`

	listStr = "a,b,c,d"

	regionStr = "region: Bogota"

	storageStr = `storage my_db {
    
    region: Bogota,
    region: Bogota

}`

	simple = `resource my_cluster {
    storage my_db {
        region: Bogota
    },
    computing my_server {
        region: Bogota
    },
    test
}`

	complete = `resource my_cluster {
    storage my_db {
        region: Bogota,
        engine: MySQL,
        CPU: 2 cores,
        memory: 2 GB,
        IPV6: no,
        storage: BLS of 16 GB
    },
    computing my_server {
        region: Bogota,
        OS: Linux,
        IPV6: yes,
        storage: SSD of 256 GB,
        CPU: 4 cores,
        memory: 8 GB
    },
    my_server,
    computing second_server {
        OS: Windows Server 2019,
        storage: BLS of 1024GB,
        CPU: 8 cores
    }
}`
)
