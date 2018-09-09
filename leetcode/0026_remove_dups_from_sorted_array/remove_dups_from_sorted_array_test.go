package pkg0026

import (
	"testing"

	"github.com/upeoe/go-dummycode/helper"
)

func TestRemoveDuplicatesV1(t *testing.T) {
	arr1 := []int{1, 1, 2}
	len := removeDuplicatesV1(arr1)
	helper.AssertEqual(t, len, 2)
	helper.AssertEqual(t, arr1[:len], []int{1, 2})

	arr1 = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	len = removeDuplicatesV1(arr1)
	helper.AssertEqual(t, len, 5)
	helper.AssertEqual(t, arr1[:len], []int{0, 1, 2, 3, 4})
}

func TestRemoveDuplicatesV2(t *testing.T) {
	arr1 := []int{1, 1, 2}
	len := removeDuplicatesV2(arr1)
	helper.AssertEqual(t, len, 2)
	helper.AssertEqual(t, arr1[:len], []int{1, 2})

	arr1 = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	len = removeDuplicatesV2(arr1)
	helper.AssertEqual(t, len, 5)
	helper.AssertEqual(t, arr1[:len], []int{0, 1, 2, 3, 4})
}
