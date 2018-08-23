package leetcode

import (
	"testing"
)

func TestEqual(t *testing.T) {
	if !Equal(nil, nil) {
		t.Error("Equal should return true")
	}
	if !Equal(123, 123) {
		t.Error("Equal should return true")
	}
	if !Equal("123", "123") {
		t.Error("Equal should return true")
	}
	if Equal(nil, 123) {
		t.Error("Equal should return false")
	}
	if Equal(123, "123") {
		t.Error("Equal should return false")
	}
}

func TestEqualValue(t *testing.T) {
	if !EqualValue(nil, nil) {
		t.Error("EqualValue should return true")
	}
	if !EqualValue(123, 123) {
		t.Error("EqualValue should return true")
	}
	if !EqualValue('a', 'a') {
		t.Error("EqualValue should return true")
	}
	if !EqualValue(uint8('a'), 'a') {
		t.Error("EqualValue should return true")
	}
	if !EqualValue(97, uint8(97)) {
		t.Error("EqualValue should return true")
	}
	if !EqualValue("a", 97) {
		t.Error("EqualValue should return true")
	}
	if EqualValue([]byte{'a'}, 'a') {
		t.Error("EqualValue should return false")
	}
}

func TestMessageFormat(t *testing.T) {
	if !Equal(messageFormat(), "") {
		t.Error("MessageFormat should return true")
	}
	if !Equal(messageFormat("a"), "a") {
		t.Error("MessageFormat should return true")
	}
	if !Equal(messageFormat("%s-%s-%s", "a", "b", "c"), "a-b-c") {
		t.Error("MessageFormat should return true")
	}
	if Equal(messageFormat(), nil) {
		t.Error("MessageFormat should return false")
	}
	if Equal(messageFormat("."), "z") {
		t.Error("MessageFormat should return false")
	}
	if Equal(messageFormat("a", "b"), "ab") {
		t.Error("MessageFormat should return false")
	}
}
