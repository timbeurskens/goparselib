package ini

import . "goparselib"

var (
	parOpen        = CTerminal("\\[")
	parClose       = CTerminal("\\]")
	eq             = CTerminal("=")
	str            = CTerminal("[a-zA-Z0-9_\\-\\ ,.;]*")
	propertyident  = Concat{Ident, BlankOpt, eq, BlankOpt}
	FloatProperty  = Concat{propertyident, Float}
	StringProperty = Concat{propertyident, str}
	Property       = Union{FloatProperty, StringProperty}
	PropertyList   = new(Symbol)
	sectionId      = Concat{parOpen, Ident, parClose}
	Section        = Concat{sectionId, R(PropertyList)}
	sections       = new(Symbol)
	Root           = Concat{R(PropertyList), R(sections)}

	// specify layout symbols for reduction
	Layout = []Symbol{parOpen, parClose, eq, BlankOpt, Blank, nil, EOL}
)

func init() {
	Define(sections, Union{
		Concat{Section, R(sections)},
		Concat{EOL, R(sections)},
		nil,
	})
	Define(PropertyList, Union{
		Concat{Property, R(PropertyList)},
		Concat{EOL, R(PropertyList)},
		nil,
	})
}
