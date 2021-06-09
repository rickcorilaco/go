package check

import (
	"testing"
)

func TestIf(t *testing.T) {
	var trueValue, falseValue int = 1, 2

	value := If(true, trueValue, falseValue).(int)
	if value != trueValue {
		t.Error("the If function should return the value to true")
	}

	value = If(false, trueValue, falseValue).(int)
	if value != falseValue {
		t.Error("the If function should return the value to false")
	}
}
