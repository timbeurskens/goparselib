package goparselib

import "fmt"

type SyntaxError struct {
	At       int64
	Expected Symbol
	Got      string
	Err      error
}

func (s SyntaxError) Error() string {
	return fmt.Sprintf("Syntax error at %d: expected %v, but got %s {\n%s\n}", s.At, s.Expected, s.Got, s.Err)
}

type InputError struct {
	At  int64
	Err error
}

func (i InputError) Error() string {
	return fmt.Sprintf("Input error at: %d (%s)", i.At, i.Err)
}

type ParserError struct {
	Symbol
}

func (p ParserError) Error() string {
	return fmt.Sprintf("Parser error when parsing symbol: %v", p.Symbol)
}
