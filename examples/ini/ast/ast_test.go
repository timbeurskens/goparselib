package ast

import (
	"encoding/json"
	"os"
	"testing"

	"goparselib"
	"goparselib/examples/ini"
)

func TestFile(t *testing.T) {
	tree, err := goparselib.ParseFile("test.ini", ini.Root)
	if err != nil {
		t.Error(err)
	}

	subtree, err := tree.Reduce(ini.Layout...)
	if err != nil {
		t.Error(err)
	}

	t.Log(subtree)

	file, err := LoadFile(subtree)
	if err != nil {
		t.Error(err)
	}
	t.Log(file)

	fout, err := os.Create("test.json")
	if err != nil {
		t.Error(err)
	}

	defer fout.Close()

	encoder := json.NewEncoder(fout)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(&file)
	if err != nil {
		t.Error(err)
	}
}
