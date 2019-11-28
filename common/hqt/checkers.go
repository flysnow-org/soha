package hqt

import (
	"errors"
	qt "github.com/frankban/quicktest"
	"reflect"
)

var IsSameType qt.Checker = &typeChecker{
	argNames: []string{"got", "want"},
}

type argNames []string

func (a argNames) ArgNames() []string {
	return a
}

type typeChecker struct {
	argNames
}

// Check implements Checker.Check by checking that got and args[0] is of the same type.
func (c *typeChecker) Check(got interface{}, args []interface{}, note func(key string, value interface{})) (err error) {
	if want := args[0]; reflect.TypeOf(got) != reflect.TypeOf(want) {
		if _, ok := got.(error); ok && want == nil {
			return errors.New("got non-nil error")
		}
		return errors.New("values are not of same type")
	}
	return nil
}