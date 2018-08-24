package pkg0001

import (
	"testing"

	"github.com/upeoe/go-dummycode/helper"
)

func TestTwoSumV1(t *testing.T) {
	helper.AssertEqual(t, twoSumV1([]int{2, 7, 11, 15}, 9), []int{0, 1})
	helper.AssertEqual(t, twoSumV1([]int{1, 2, 11, 23}, 25), []int{1, 3})
}

func TestTwoSumV2(t *testing.T) {
	helper.AssertEqual(t, twoSumV2([]int{2, 7, 11, 15}, 9), []int{0, 1})
	helper.AssertEqual(t, twoSumV2([]int{1, 2, 11, 23}, 25), []int{1, 3})
}

func TestTwoSumV3(t *testing.T) {
	helper.AssertEqual(t, twoSumV3([]int{2, 7, 11, 15}, 9), []int{0, 1})
	helper.AssertEqual(t, twoSumV3([]int{1, 2, 11, 23}, 25), []int{1, 3})
}
