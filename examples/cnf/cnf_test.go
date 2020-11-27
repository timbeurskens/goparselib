package cnf

import (
	"encoding/json"
	"os"
	"testing"

	"goparselib"
)

func TestParse(t *testing.T) {
	t.Log(goparselib.Parse("(a&-a&-a)#", root))
	t.Log("------")
	str := "((a&-b&c)|-(-a|(a|b|c)))#"
	_, tree := goparselib.Parse(str, root)
	tree.Populate(str)

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(tree); err != nil {
		t.Error(err)
	}
}
