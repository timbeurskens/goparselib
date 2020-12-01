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
	Comma    = CTerminal("\\,")
)

// Plus creates a symbol matching one or more of the symbol 'of'
func Plus(of Symbol) Symbol {
	plus := new(Symbol)
	Define(plus, Union{
		of,
		Concat{of, R(plus)},
	})
	return R(plus)
}

// Optional creates a symbol matching zero or one of the symbol 'of'
func Optional(of Symbol) Symbol {
	return Union{nil, of}
}

// List creates a symbol matching a list of elements 'of' separated by 'separator', with optional 'blank'
func List(of, separator, blank Symbol) Symbol {
	list := new(Symbol)

	Define(list, Union{
		of,
		Concat{of, separator, R(list)},
		Concat{blank, R(list)},
	})

	// a list is either an element following a list, or it is empty
	return Union{R(list), Optional(blank)}
}
