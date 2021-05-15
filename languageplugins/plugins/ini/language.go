package ini

import . "github.com/timbeurskens/goparselib"

var (
	parOpen        = CTerminal("\\[")
	parClose       = CTerminal("\\]")
	eq             = CTerminal("=")
	str            = CTerminal("[a-zA-Z0-9_\\-\\ ,.;']*")
	propertyident  = Concat{Ident, BlankOpt, eq, BlankOpt}
	FloatProperty  = Concat{propertyident, Float}
	StringProperty = Concat{propertyident, str}
	Property       = Union{FloatProperty, StringProperty}
	PropertyList   = List(Property, nil, LF)
	sectionId      = Concat{parOpen, Ident, parClose}
	Section        = Concat{sectionId, PropertyList}
	sections       = List(Section, nil, LF)
	Root           = Concat{PropertyList, sections}

	// Layout specifies the layout symbols for reduction
	Layout = []Symbol{parOpen, parClose, eq, BlankOpt, Blank, nil, LF}
)
