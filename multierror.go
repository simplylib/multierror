// Package multierror contains a Multierror for combining multiple errors and compiling with errors.Is
package multierror

import (
	"errors"
	"strings"
)

// Errors is a collection of errors combined and treated as one.
type Errors []error

// Error returns combined error.Error output by \n
func (me Errors) Error() string {
	builder := strings.Builder{}
	length := len(me) - 1

	for i, e := range me {
		if i == length {
			builder.WriteString(e.Error())
		} else {
			builder.WriteString(e.Error() + "\n")
		}
	}

	return builder.String()
}

// Unwrap MultiError for conforming to errors.Unwrap
func (me Errors) Unwrap() error {
	switch len(me) {
	case 0:
		return nil
	case 1:
		return me[0]
	default:
		return me[1:]
	}
}

// Is a target error, complies with errors.Is
func (me Errors) Is(target error) bool {
	for _, e := range me {
		if errors.Is(e, target) {
			return true
		}
	}

	return false
}

// Append new and old error together as a Multierror, reusing a multierror if it already exists.
// If either error is nil, the other error is returned.
func Append(oldErr error, newErr error) error {
	if oldErr == nil {
		return newErr
	}

	if newErr == nil {
		return oldErr
	}

	switch t := oldErr.(type) {
	case Errors:
		return append(t, newErr)
	default:
		return Errors{oldErr, newErr}
	}
}
