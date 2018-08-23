package leetcode

import (
	"testing"
)

func TestLowercaseToUppercase(t *testing.T) {
	AssertEqualValue(t, LowercaseToUppercase('a'), 'A')
	AssertEqualValue(t, LowercaseToUppercase('s'), 'S')
	AssertEqualValue(t, LowercaseToUppercase('z'), 'Z')
}
