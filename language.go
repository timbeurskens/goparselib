package goparselib

var (
    // start = CTerminal("\\`")
    end = CTerminal("x")
    number = CTerminal("[0-9]+")
    operator = CTerminal("\\+")
    rclose = CTerminal("\\)")
    lopen = CTerminal("\\(")

    expression = new (Symbol)
    root = new (Symbol)
)

const (
    example1 = "25x"
    example2 = "(20+10)x"
    example3 = "((1844+2546)+(3+(4+5)))x"

    counterexample1 = "(20+10x"
    counterexample2 = "20+10x"
    counterexample3 = "(20)x"
)

func init() {
    *expression = Union{number, Concat{lopen, R(expression), operator, R(expression), rclose}}
    *root = Concat{R(expression), end}
}