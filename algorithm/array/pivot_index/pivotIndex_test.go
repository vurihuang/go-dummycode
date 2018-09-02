package pivot_index

import (
	"testing"

	"github.com/upeoe/go-dummycode/helper"
)

func TestPivotIndex(t *testing.T) {
	helper.AssertEqual(t, pivotIndex([]int{1, 7, 3, 6, 5, 6}), 3)
	helper.AssertEqual(t, pivotIndex([]int{1, 2, 3}), -1)
}
