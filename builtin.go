package goparselib

var (
    Ident    = CTerminal("[a-zA-Z_][a-zA-Z0-9_]*")
    EOL      = CTerminal("[[:blank:]]*\\\n")
    Blank    = CTerminal("[[:blank:]]+")
    BlankOpt = CTerminal("[[:blank:]]*")
    Float    = CTerminal("[+\\-]?([0]|[1-9][0-9]*)(\\.[0-9]+)?")
    Integer  = CTerminal("[+\\-]?([0]|[1-9][0-9]*)")
    Natural  = CTerminal("[0]|[1-9][0-9]*")
    LBracket = CTerminal("\\{")
    RBracket = CTerminal("\\}")
    Comma = CTerminal("\\,")
)
