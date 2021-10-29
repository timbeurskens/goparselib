package testing

import (
	"github.com/timbeurskens/goparselib"
	"github.com/timbeurskens/goparselib/parser"
	"strings"
	"testing"
)

func DoTestInput(t *testing.T, input map[string]string, symbol goparselib.Symbol) {
	for name, testStr := range input {
		t.Run(name, func(t *testing.T) {
			result, err := parser.ParseString(testStr, symbol)

			if err != nil {
				t.Error(err)
			}

			builder := new(strings.Builder)
			result.Walk(func(node goparselib.Node) bool {
				builder.WriteString(node.Contents)
				return true
			})

			if testStr != builder.String() {
				t.Errorf("Mismatch between source and result:\n BEFORE:%s\nAFTER:%s\n", testStr, builder.String())
			}
		})
	}
}
