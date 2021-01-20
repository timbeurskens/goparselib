package goparselib

var (
	LF       = CTerminal("\\\n")
	CRLF     = CTerminal("\\\r\\\n")
	Ident    = CTerminal("[a-zA-Z_][a-zA-Z0-9_]*")
	Colon    = CTerminal("\\:")
	Eq       = CTerminal("\\=")
	Blank    = CTerminal("[[:blank:]]+")
	BlankOpt = CTerminal("[[:blank:]]*")
	Null     = CTerminal("\\x00")
	Float    = CTerminal("[+\\-]?([0]|[1-9][0-9]*)(\\.[0-9]+)?")
	Integer  = CTerminal("[+\\-]?([0]|[1-9][0-9]*)")
	Natural  = CTerminal("[0]|[1-9][0-9]*")
	LBracket = CTerminal("\\{")
	RBracket = CTerminal("\\}")
	LParen   = CTerminal("\\(")
	RParen   = CTerminal("\\)")
	Comma    = CTerminal("\\,")

	// deprecated: replaced by LF and CRLF
	EOL = CTerminal("\\\n?\\\n")
)

// Plus creates a symbol matching one or more of the symbol 'of'
func Plus(of Symbol) Symbol {
	plus := new(Symbol)
	Define(plus, Union{
		Concat{of, R(plus)},
		of,
	})
	return R(plus)
}

// Optional creates a symbol matching zero or one of the symbol 'of'
func Optional(of Symbol) Symbol {
	return Union{of, nil}
}

// List creates a symbol matching a list of elements 'of' separated by 'separator', with optional 'blank'
// If an empty list is supported, add a 'nil' element to the 'of' field (union)
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

func Decorate(of, dec Symbol) Symbol {
	switch of.(type) {
	case Union:
		items := of.(Union)
		n := len(items)
		result := make(Union, (2*n)+1)

		for i := range items {
			result[(2*i)+1] = items[i]
			result[2*i] = dec
		}

		result[2*n] = dec

		return result
	case Concat:
		items := of.(Concat)
		n := len(items)
		result := make(Concat, (2*n)+1)

		for i := range items {
			result[(2*i)+1] = items[i]
			result[2*i] = dec
		}

		result[2*n] = dec

		return result
	default:
		return of
	}
}
