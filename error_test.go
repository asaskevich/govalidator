package govalidator

import (
	"fmt"
	"testing"
)

func TestErrorsToString(t *testing.T) {
	t.Parallel()
	customErr := &Error{Name: "custom error name", Err: fmt.Errorf("stdlib error")}
	customErrWithCustomErrorMessage := &Error{Name: "custom error name 2", Err: fmt.Errorf("bad stuff happened"), CustomErrorMessageExists: true}

	var tests = []struct {
		param1   Errors
		expected string
	}{
		{Errors{}, ""},
		{Errors{fmt.Errorf("error 1")}, "error 1"},
		{Errors{fmt.Errorf("error 1"), fmt.Errorf("error 2")}, "error 1;error 2"},
		{Errors{customErr, fmt.Errorf("error 2")}, "custom error name: stdlib error;error 2"},
		{Errors{fmt.Errorf("error 123"), customErrWithCustomErrorMessage}, "bad stuff happened;error 123"},
	}
	for _, test := range tests {
		actual := test.param1.Error()
		if actual != test.expected {
			t.Errorf("Expected Error() to return '%v', got '%v'", test.expected, actual)
		}
	}
}
