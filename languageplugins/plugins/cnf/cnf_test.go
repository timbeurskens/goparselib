package cnf

import (
	"testing"

	parsetest "github.com/timbeurskens/goparselib/testing"
)

func TestCNF(t *testing.T) {
	parsetest.DoTestInput(t, map[string]string{
		"simple_expr_1":   `(a|b)`,
		"simple_expr_2":   `(a&b)`,
		"composed_expr_1": `((a|b)&b)`,
		"triple_expr_1":   `(a&b&c)`,
		"triple_neg_1":    `(a&-a&-a)`,
		"complex_1":       `((a&-b&c)|-(-a|(a|b|c)))`,
	}, Root)
}
