package leetcode

import "testing"

func TestReverseInteger(t *testing.T) {
	AssertEqual(t, reverseInteger(0), 0)
	AssertEqual(t, reverseInteger(321), 123)
	AssertEqual(t, reverseInteger(100), 1)
	AssertEqual(t, reverseInteger(900), 9)
}
