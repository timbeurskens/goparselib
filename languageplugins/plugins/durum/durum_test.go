package durum

import (
	"testing"

	"github.com/timbeurskens/goparselib/parser"

	"github.com/timbeurskens/goparselib"
)

func TestStateDefinition(t *testing.T) {
	t.Parallel()
	DoTestInput(t, map[string]string{
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
	}, State)
}

func TestDefinitions(t *testing.T) {
	t.Parallel()
	DoTestInput(t, map[string]string{
		"basic integer": `def counter = int32_t`,
		"parameter int": `def led = gpio_input 10`,
	}, Definition)
}

func TestFromFile(t *testing.T) {
	result, err := parser.ParseFile("examples/basic.du", Root)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestFullPartial(t *testing.T) {
	t.Parallel()
	DoTestInput(t, map[string]string{
		"basic integer": `def counter = int32_t`,
		"parameter int": `def led = gpio_input 10`,
	}, Root)
	DoTestInput(t, map[string]string{
		"timer tick": `event timer_tick = timer tick`,
	}, Root)
	DoTestInput(t, map[string]string{
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
	}, Root)
	DoTestInput(t, map[string]string{
		"no_param":     `action test {}`,
		"single_param": `action test (a) {}`,
	}, Root)
}

func TestFull(t *testing.T) {
	t.Parallel()
	DoTestInput(t, map[string]string{
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
	}, Root)
}

func TestEvents(t *testing.T) {
	t.Parallel()
	DoTestInput(t, map[string]string{
		"timer tick": `event timer_tick = timer tick`,
	}, Event)
}

func TestActions(t *testing.T) {
	t.Parallel()
	DoTestInput(t, map[string]string{
		"no_param":     `action test {}`,
		"single_param": `action test (a) {}`,
	}, Action)
}

func DoTestInput(t *testing.T, input map[string]string, symbol goparselib.Symbol) {
	for name, testStr := range input {
		t.Run(name, func(t *testing.T) {
			result, err := parser.ParseString(testStr, symbol)
			if err != nil {
				t.Error(err)
			}
			t.Log(result)
		})
	}
}
