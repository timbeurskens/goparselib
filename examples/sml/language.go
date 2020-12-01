package sml

import . "goparselib"

var (
	programLit   = CTerminal("program")
	timelineLit  = CTerminal("timeline")
	lopen        = CTerminal("\\{")
	rclose       = CTerminal("\\}")
	eol          = CTerminal("[[:blank:]]*\\\n")
	eof          = CTerminal("#")
	space        = CTerminal("[[:blank:]]*")
	integer      = CTerminal("[0-9]+")
	enumEnd      = CTerminal("\\)")
	durationUnit = CTerminal("milliseconds|seconds|minutes|hours")
	play         = CTerminal("play")
	forLit       = CTerminal("for")
	end          = CTerminal("end")
	progName     = CTerminal("[a-zA-Z_]+")
	duration     = Concat{integer, space, durationUnit}
	playCmd      = Concat{play, space, progName, space, forLit, space, duration}
	command      = Union{playCmd}
	timelineLine = Concat{space, integer, enumEnd, space, command, eol}
	programBody  = new(Symbol)
	timelineBody = new(Symbol)
	timeline     = new(Symbol)
	program      = new(Symbol)
	section      = new(Symbol)
	root         = new(Symbol)
)

func init() {
	Define(programBody, Union{space})
	Define(timelineBody, Union{nil, Concat{timelineLine, R(timelineBody)}})
	Define(program, Concat{programLit, space, lopen, eol, R(programBody), rclose})
	Define(timeline, Concat{timelineLit, space, lopen, eol, R(timelineBody), rclose})
	Define(section, Union{R(program), R(timeline)})
	Define(root, Union{eof, Concat{R(section), R(root)}})
}
