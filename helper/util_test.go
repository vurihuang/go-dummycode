package helper

import "testing"

func TestMessageFormat(t *testing.T) {
	if MessageFormat() != "" {
		t.Error("MessageFormat should return true")
	}
	if MessageFormat("a") != "a" {
		t.Error("MessageFormat should return true")
	}
	if MessageFormat("%s-%s", "a", "b") != "a-b" {
		t.Error("MessageFormat should return true")
	}
	if MessageFormat("a", "b") == "ab" {
		t.Error("MessageFormat should return false")
	}
}

func TestMessageFormatEqual(t *testing.T) {
	if MessageFormatEqual("a", "b") != "a != b" {
		t.Error("MessageFormatEqual should return true")
	}
	if MessageFormatEqual("a", "b") == "a!=b" {
		t.Error("MessageFormatEqual should return false")
	}
}
