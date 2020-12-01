package ccl

import (
	"strings"
	"testing"

	"goparselib"
)

func TestSimple(t *testing.T) {
	tree, err := goparselib.Parse(simple, root)
	if err != nil {
		t.Error(err)
	}
	t.Log(tree)
}

func TestComplete(t *testing.T) {
	tree, err := goparselib.Parse(complete, root)
	if err != nil {
		t.Error(err)
	}
	tree.Populate(complete)
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
	tree, err := goparselib.Parse(storageStr, storage)
	if err != nil {
		t.Error(err)
	}
	t.Log(tree)
}

func TestEmptyBody(t *testing.T) {
	tree, err := goparselib.Parse(emptyBody, storageBody)
	if err != nil {
		t.Error(err)
	}
	t.Log(tree)
}

func TestRegionBody(t *testing.T) {
	tree, err := goparselib.Parse(regionStr, storageBody)
	if err != nil {
		t.Error(err)
	}
	t.Log(tree)
}

func TestList(t *testing.T) {
	tree, err := goparselib.Parse(listStr, goparselib.List(goparselib.Ident, goparselib.Comma, goparselib.Blank))
	if err != nil {
		t.Error(err)
	}
	tree.Populate(listStr)
	t.Log(tree)
}

func TestEmptyList(t *testing.T) {
	tree, err := goparselib.Parse(emptyBody, goparselib.List(goparselib.Ident, goparselib.Comma, goparselib.Blank))
	if err != nil {
		t.Error(err)
	}
	tree.Populate(listStr)
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
