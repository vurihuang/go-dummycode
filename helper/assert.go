package helper

import (
	"bytes"
	"reflect"
	"testing"
)

func Equal(actual, expected interface{}) bool {
	if actual == nil || expected == nil {
		return actual == expected
	}

	act, ok := actual.([]byte)
	if !ok {
		return reflect.DeepEqual(actual, expected)
	}

	exp, ok := expected.([]byte)
	if !ok {
		return false
	}

	if act == nil || exp == nil {
		return act == nil && exp == nil
	}

	return bytes.Equal(act, exp)
}

func EqualValue(actual, expected interface{}) bool {
	if Equal(actual, expected) {
		return true
	}

	actType := reflect.TypeOf(actual)
	if actType == nil {
		return false
	}

	expVal := reflect.ValueOf(expected)
	if expVal.IsValid() && expVal.Type().ConvertibleTo(actType) {
		return reflect.DeepEqual(expVal.Convert(actType).Interface(), actual)
	}

	return false
}

func AssertEqual(t *testing.T, actual, expected interface{}, msg ...interface{}) {
	if Equal(actual, expected) {
		return
	}
	t.Error(MessageFormatEqual(actual, expected))
}

func AssertEqualValue(t *testing.T, actual, expected interface{}, msg ...interface{}) {
	if EqualValue(actual, expected) {
		return
	}
	t.Error(MessageFormatEqual(actual, expected))
}
