package ini

import . "goparselib"

var (
	eof        = CTerminal("#")
	open       = CTerminal("\\[")
	close      = CTerminal("\\]")
	eq         = CTerminal("=")
	str        = CTerminal("[a-z]*")
	value      = Union{Float, str}
	property   = Concat{Ident, BlankOpt, eq, BlankOpt, value, EOL}
	properties = new(Symbol)
	sectionId  = Concat{open, Ident, close, EOL}
	section    = Concat{sectionId, R(properties)}
	sections   = new(Symbol)
	root       = Concat{R(properties), R(sections), eof}
)

func init() {
	Define(sections, Union{nil, Concat{EOL, R(sections)}, Concat{section, R(sections)}})
	Define(properties, Union{nil, Concat{EOL, R(properties)}, Concat{property, R(properties)}})
}
