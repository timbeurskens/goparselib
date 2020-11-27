package goparselib

import (
    "fmt"
    "regexp"
)

type Symbol interface {
    symbol()
    fmt.Stringer
}

type Terminal struct {
    Reg *regexp.Regexp
}

func (t Terminal) String() string {
    return fmt.Sprintf("T(%s)", t.Reg.String())
}

func (t Terminal) symbol() {
    panic("implement me")
}

type Union []Symbol

func (u Union) String() string {
    return fmt.Sprintf("U(%d)", len(u))
}

func (u Union) symbol() {
    panic("implement me")
}

type Concat []Symbol

func (c Concat) String() string {
    return fmt.Sprintf("C(%d)", len(c))
}

func (c Concat) symbol() {
    panic("implement me")
}

type Reference struct {
    R *Symbol
}

func (r Reference) String() string {
    return fmt.Sprintf("R(%d)", r.R)
}

func (r Reference) symbol() {
    panic("implement me")
}

func R(s *Symbol) Reference {
    return Reference{s}
}

func CTerminal(expr string) Terminal {
    return Terminal{regexp.MustCompilePOSIX(expr)}
}

func Define(s *Symbol, s2 Symbol) {
    *s = s2
}

type Node struct {
    Start    int64  `json:"start"`
    Size     int64  `json:"size"`
    Contents string `json:"contents"`
    Type     Symbol `json:"-"`
    Children []Node `json:"children"`
}
