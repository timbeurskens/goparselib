package durum

import . "github.com/timbeurskens/goparselib"

var (
	FDelim       = CTerminal("---")
	Def          = CTerminal("def")
	StateLit     = CTerminal("state")
	Start        = CTerminal("start")
	OnLit        = CTerminal("on")
	GotoLit      = CTerminal("goto")
	EventLit     = CTerminal("event")
	ActionLit    = CTerminal("action")
	Do           = CTerminal("do")
	DefParameter = Union{Integer, Float}

	Definition = Decorate(Concat{Def, Ident, Eq, Ident, Optional(DefParameter)}, BlankOpt)
	Transition = Decorate(Concat{OnLit, Ident, GotoLit, Ident}, BlankOpt)
	RunAction  = Decorate(Concat{Do, Ident, Optional(Decorate(Concat{LParen, List(Union{Ident, Integer, Float, nil}, Comma, Plus(Blank)), RParen}, BlankOpt))}, BlankOpt)
	StateBody  = List(Union{Transition, RunAction, nil}, LF, Plus(Union{LF, Blank}))
	ActionBody = Concat{nil}
	State      = Decorate(Concat{Optional(Start), StateLit, Ident, LBracket, StateBody, RBracket}, BlankOpt)
	Event      = Decorate(Concat{EventLit, Ident, Eq, Ident, Ident}, BlankOpt)
	Action     = Decorate(Concat{ActionLit, Ident, Optional(Decorate(Concat{LParen, List(Union{Ident, nil}, Comma, Plus(Blank)), RParen}, BlankOpt)), LBracket, ActionBody, RBracket}, BlankOpt)
	Root       = List(Union{Definition, Event, State, Action, nil}, LF, Plus(Union{LF, Blank}))

	RootCheck = Decorate(Concat{Root, FDelim}, BlankOpt)

	Layout = []Symbol{
		LF, Blank, FDelim, nil,
	}
)
