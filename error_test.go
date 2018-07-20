package govalidator

import (
	"fmt"
	"testing"
)

func TestErrorsToString(t *testing.T) {
	t.Parallel()
	customErr := &Error{Path: []string{"Custom Error Name"}, Err: fmt.Errorf("stdlib error")}
	customErrWithCustomErrorMessage := &Error{Path: []string{"Custom Error Name 2"}, Err: fmt.Errorf("Bad stuff happened"), CustomErrorMessageExists: true}

	var tests = []struct {
		param1   Errors
		expected string
	}{
		{Errors{}, ""},
		{Errors{fmt.Errorf("Error 1")}, "Error 1"},
		{Errors{fmt.Errorf("Error 1"), fmt.Errorf("Error 2")}, "Error 1;Error 2"},
		{Errors{customErr, fmt.Errorf("Error 2")}, "Custom Error Name: stdlib error;Error 2"},
		{Errors{fmt.Errorf("Error 123"), customErrWithCustomErrorMessage}, "Error 123;Bad stuff happened"},
	}
	for _, test := range tests {
		actual := test.param1.Error()
		if actual != test.expected {
			t.Errorf("Expected Error() to return '%v', got '%v'", test.expected, actual)
		}
	}
}

func TestPrependPathToErrors(t *testing.T) {
	var tests = []struct {
		err      Errors
		expected Errors
	}{
		{Errors{&Error{Err: fmt.Errorf("Some Error Occured"), Path: []string{"Field"}}}, Errors{&Error{Err: fmt.Errorf("Some Error Occured"), Path: []string{"foo", "bar", "Field"}}}},
	}
	for _, test := range tests {
		prependedErrors := PrependPathToErrors(test.err, "foo", "bar")

		if prependedErrors.Error() != test.expected.Error() {
			t.Errorf("Expected Error() to return '%v', got '%v'", test.expected.Error(), prependedErrors.Error())
		}
	}
}

func TestAppendPathToErrors(t *testing.T) {
	var tests = []struct {
		err      Errors
		expected Errors
	}{
		{Errors{&Error{Err: fmt.Errorf("Some Error Occured"), Path: []string{"Field"}}}, Errors{&Error{Err: fmt.Errorf("Some Error Occured"), Path: []string{"Field", "foo", "bar"}}}},
	}
	for _, test := range tests {
		appendedErrors := AppendPathToErrors(test.err, "foo", "bar")

		if appendedErrors.Error() != test.expected.Error() {
			t.Errorf("Expected Error() to return '%v', got '%v'", test.expected.Error(), appendedErrors.Error())
		}
	}
}
