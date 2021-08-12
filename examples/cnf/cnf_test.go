package cnf

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/timbeurskens/goparselib"
)

func TestParse(t *testing.T) {
	t.Log(goparselib.ParseString("(a&-a&-a)#", root))
	t.Log("------")
	str := "((a&-b&c)|-(-a|(a|b|c)))#"
	tree, _ := goparselib.ParseString(str, root)

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(tree); err != nil {
		t.Error(err)
	}
}
