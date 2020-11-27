package cnf

import . "goparselib"

var (
	variable       = CTerminal("[a-z][a-z0-9]*")
	end            = CTerminal("#")
	and            = CTerminal("&")
	or             = CTerminal("\\|")
	implies        = CTerminal("=>")
	biimplies      = CTerminal("<=>")
	not            = CTerminal("-")
	binaryoperator = Union{and, or, implies, biimplies}
	unaryoperator  = Union{not}
	term           = Union{variable, Concat{unaryoperator, variable}}
	lopen          = CTerminal("\\(")
	rclose         = CTerminal("\\)")
	conjunction    = new(Symbol)
	disjunction    = new(Symbol)
	expression     = new(Symbol)
	root           = Concat{R(expression), end}
)

func init() {
	Define(conjunction, Union{term, Concat{term, and, R(conjunction)}})
	Define(disjunction, Union{term, Concat{term, or, R(disjunction)}})
	Define(expression, Union{
		term,
		Concat{lopen, R(conjunction), rclose},
		Concat{lopen, R(disjunction), rclose},
		Concat{lopen, R(expression), binaryoperator, R(expression), rclose},
		Concat{unaryoperator, R(expression)},
	})
}
