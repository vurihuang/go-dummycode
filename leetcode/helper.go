package leetcode

import (
	"bytes"
	"fmt"
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

	t.Fatal(messageFormatEqual(actual, expected, msg))
}

func AssertEqualValue(t *testing.T, actual, expected interface{}, msg ...interface{}) {
	if EqualValue(actual, expected) {
		return
	}

	t.Fatal(messageFormatEqual(actual, expected, msg))
}

func messageFormat(msg ...interface{}) string {
	if len(msg) == 1 {
		return msg[0].(string)
	}
	if len(msg) > 1 {
		return fmt.Sprintf(msg[0].(string), msg[1:]...)
	}
	return ""
}

func messageFormatEqual(actual, expected interface{}, msg ...interface{}) string {
	message := ""
	if msg == nil || len(msg) == 0 {
		message = fmt.Sprintf("%v != %v", actual, expected)
	} else {
		message = messageFormat(msg)
	}
	return message
}
