package cnf

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/timbeurskens/goparselib/parser"
)

func TestParse(t *testing.T) {
	t.Log(parser.ParseString("(a&-a&-a)#", Root))
	t.Log("------")
	str := "((a&-b&c)|-(-a|(a|b|c)))#"
	tree, _ := parser.ParseString(str, Root)

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(tree); err != nil {
		t.Error(err)
	}
}
