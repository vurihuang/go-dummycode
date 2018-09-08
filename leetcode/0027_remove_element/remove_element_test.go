package pkg0027

import (
	"testing"

	"github.com/upeoe/go-dummycode/helper"
)

func TestRemoveElementV1(t *testing.T) {
	var (
		arr1 = []int{1, 1, 2}
		arr2 = []int{3, 2, 2, 3}
		arr3 = []int{0, 1, 2, 2, 3, 0, 4, 2}
	)
	var ret int

	ret = removeElementV1(arr1, 2)
	helper.AssertEqual(t, ret, 2)
	helper.AssertEqual(t, arr1[:ret], []int{1, 1})

	ret = removeElementV1(arr2, 3)
	helper.AssertEqual(t, ret, 2)
	helper.AssertEqual(t, arr2[:ret], []int{2, 2})

	ret = removeElementV1(arr3, 2)
	helper.AssertEqual(t, ret, 5)
	helper.AssertEqual(t, arr3[:ret], []int{0, 1, 3, 0, 4})
}

func TestRemoveElementV2(t *testing.T) {
	var (
		arr1 = []int{1, 1, 2}
		arr2 = []int{3, 2, 2, 3}
		arr3 = []int{0, 1, 2, 2, 3, 0, 4, 2}
	)
	var ret int

	ret = removeElementV2(arr1, 2)
	helper.AssertEqual(t, ret, 2)
	helper.AssertEqual(t, arr1[:ret], []int{1, 1})

	ret = removeElementV2(arr2, 3)
	helper.AssertEqual(t, ret, 2)
	helper.AssertEqual(t, arr2[:ret], []int{2, 2})

	ret = removeElementV2(arr3, 2)
	helper.AssertEqual(t, ret, 5)
	helper.AssertEqual(t, arr3[:ret], []int{0, 1, 4, 0, 3})
}
