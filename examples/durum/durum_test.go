package durum

import (
	"testing"

	"github.com/timbeurskens/goparselib"
)

var definitions = map[string]string{
	"basic integer": `def counter = int32_t`,
	"parameter int": `def led = gpio_input 10`,
}

var events = map[string]string{
	"timer tick": `event timer_tick = timer tick`,
}

var actions = map[string]string{
	"no_param":     `action test {}`,
	"single_param": `action test (a) {}`,
}

var states = map[string]string{
	"basic":         `state initial {}`,
	"basic_initial": `start state initial {}`,
	"basic_eol": `start state initial {
}`,
	"1 transition": `start state initial {
	on some_event goto next
}`,
	"multi transition": `start state initial {
	on some_event goto next
	on other_event goto previous
	on tick goto this_state
}`,
	"single_action_call": `start state initial {
	on some_event goto next
	do some_action
}`,
	"single_param_action_call": `start state initial {
	on some_event goto next
	do some_action (10)
}`,
	"multi_param_action_call": `start state initial {
	on some_event goto next
	do some_action (a, 10, zozo)
}`,
}

var full = map[string]string{
	"basic": `def counter = int32_t 10

event tick = counter overflow

start state initial {
	on tick goto next
}

state next {
	on tick goto initial
	do action
}
`,
}

func TestStateDefinition(t *testing.T) {
	t.Parallel()
	DoTestInput(t, states, State)
}

func TestDefinitions(t *testing.T) {
	t.Parallel()
	DoTestInput(t, definitions, Definition)
}

func TestFullPartial(t *testing.T) {
	t.Parallel()
	DoTestInput(t, definitions, Root)
	DoTestInput(t, events, Root)
	DoTestInput(t, states, Root)
	DoTestInput(t, actions, Root)
}

func TestFull(t *testing.T) {
	t.Parallel()
	DoTestInput(t, full, Root)
}

func TestEvents(t *testing.T) {
	t.Parallel()
	DoTestInput(t, events, Event)
}

func TestActions(t *testing.T) {
	t.Parallel()
	DoTestInput(t, actions, Action)
}

func DoTestInput(t *testing.T, input map[string]string, symbol goparselib.Symbol) {
	for name, testStr := range input {
		t.Run(name, func(t *testing.T) {
			result, err := goparselib.ParseString(testStr, symbol)
			if err != nil {
				t.Error(err)
			}
			t.Log(result)
		})
	}
}
