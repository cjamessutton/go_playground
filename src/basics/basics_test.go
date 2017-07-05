package basics_test

import (
	"basics"
	"fmt"
	. "test_util"
	"testing"
)

func TestExport(t *testing.T) {
	Equals(t, "This is exported", basics.Exported)
	fmt.Println(basics.Exported)
}

func TestAdd(t *testing.T) {
	for name, test := range testAddTable {
		t.Logf("Running test case %s", name)
		result, _ := basics.Add(test.input[0], test.input[1])
		Equals(t, test.output, result)
	}
}

var testAddTable = map[string]struct {
	input  [2]int
	output int
	err    error
}{
	"test1": {
		input:  [2]int{1, 2},
		output: 3,
		err:    nil,
	},
}

func TestFor(t *testing.T) {
	t.Logf("Testing ForFunc")
	for name, test := range testForTable {
		t.Logf("Running test case %s", name)
		x, y, _ := basics.ForFunc(test.input)
		Equals(t, test.output[0], x)
		Equals(t, test.output[1], y)
	}
}

var testForTable = map[string]struct {
	input  int
	output [2]int
	err    error
}{
	"'proper input'": {
		input:  100,
		output: [2]int{110, 120},
		err:    nil,
	},
}

func TestSwitch(t *testing.T) {
	result, _ := basics.SwitchFunc("1")
	if result != "First String" {
		t.Errorf("whoops")
	}
}

func TestDefer(t *testing.T) {
	basics.DeferFunc()
}

func TestArrays(t *testing.T) {
	basics.Arrays()
}

func TestSlices(t *testing.T) {
	basics.Slices()
}

func TestRange(t *testing.T) {
	basics.Range()
}

func TestMaps(t *testing.T) {
	basics.Maps()
}

func TestTakesValue(t *testing.T) {
	basics.TakesValue(basics.AsValue)
}

func TestClosures(t *testing.T) {
	fn := basics.Closures()
	fmt.Println(fn())
}
