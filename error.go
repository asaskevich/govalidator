package govalidator

import (
	"strings"
)

// Errors is an array of multiple errors and conforms to the error interface.
type Errors []error

// Errors returns itself.
func (es Errors) Errors() []error {
	return es
}

func (es Errors) Error() string {
	var errs []string
	for _, e := range es {
		errs = append(errs, e.Error())
	}
	return strings.Join(errs, ";")
}

// Error encapsulates a name, an error and whether there's a custom error message or not.
type Error struct {
	Err                      error
	CustomErrorMessageExists bool

	// Validator indicates the name of the validator that failed
	Validator string
	Path      []string
}

// Name returns the name of the last element in the path
func (e Error) Name() string {
	return e.Path[len(e.Path)-1]
}

// FullPath returns a dot separated path of where the error occured
func (e Error) FullPath() string {
	return strings.Join(e.Path, ".")
}

func (e Error) Error() string {
	if e.CustomErrorMessageExists {
		return e.Err.Error()
	}

	return e.FullPath() + ": " + e.Err.Error()
}

// PrependPathToErrors prepends the paths to the desired error.
// This recursively loops through all children prepending the path at every error it encounters
func PrependPathToErrors(err error, path ...string) error {
	switch err2 := err.(type) {
	case Error:
		err2.Path = append(path, err2.Path...)
		return err2
	case *Error:
		err2.Path = append(path, err2.Path...)
		return err2
	case Errors:
		errors := err2.Errors()
		for i, err3 := range errors {
			errors[i] = PrependPathToErrors(err3, path...)
		}
		return err2
	}
	return err
}

// AppendPathToErrors appends the paths to the desired error.
// This recursively loops through all children appending the path at every error it encounters
func AppendPathToErrors(err error, path ...string) error {
	switch err2 := err.(type) {
	case Error:
		err2.Path = append(err2.Path, path...)
		return err2
	case *Error:
		err2.Path = append(err2.Path, path...)
		return err2
	case Errors:
		errors := err2.Errors()
		for i, err3 := range errors {
			errors[i] = AppendPathToErrors(err3, path...)
		}
		return err2
	}
	return err
}
