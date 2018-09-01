package pkg0003

import (
	"testing"

	"github.com/upeoe/go-dummycode/helper"
)

var (
	var1 string = "pwwkew"
	var2 string = "abcabcbb"
	var3 string = "asdfasdasdasdfg"
)

func TestLengthOfLongestSubstringV1(t *testing.T) {
	helper.AssertEqual(t, lengthOfLongestSubstringV1(var1), 3)
	helper.AssertEqual(t, lengthOfLongestSubstringV1(var2), 3)
	helper.AssertEqual(t, lengthOfLongestSubstringV1(var3), 5)
}

func TestLengthOfLongestSubstringV2(t *testing.T) {
	helper.AssertEqual(t, lengthOfLongestSubstringV2(var1), 3)
	helper.AssertEqual(t, lengthOfLongestSubstringV2(var2), 3)
	helper.AssertEqual(t, lengthOfLongestSubstringV2(var3), 5)
}

func TestLengthOfLongestSubstringV3(t *testing.T) {
	helper.AssertEqual(t, lengthOfLongestSubstringV3(var1), 3)
	helper.AssertEqual(t, lengthOfLongestSubstringV3(var2), 3)
	helper.AssertEqual(t, lengthOfLongestSubstringV3(var3), 5)
}

func TestLengthOfLongestSubstringV4(t *testing.T) {
	helper.AssertEqual(t, lengthOfLongestSubstringV4(var1), 3)
	helper.AssertEqual(t, lengthOfLongestSubstringV4(var2), 3)
	helper.AssertEqual(t, lengthOfLongestSubstringV4(var3), 5)
}
