package cnf

import (
	"encoding/json"
	"os"
	"testing"

	"goparselib"
)

func TestParse(t *testing.T) {
	t.Log(goparselib.ParseString("(a&-a&-a)#", root))
	t.Log("------")
	str := "((a&-b&c)|-(-a|(a|b|c)))#"
	tree, _ := goparselib.ParseString(str, root)
	tree.Populate(str)

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(tree); err != nil {
		t.Error(err)
	}
}
