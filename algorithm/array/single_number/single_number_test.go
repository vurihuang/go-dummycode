package single_num

import (
	"testing"

	"github.com/upeoe/go-dummycode/helper"
)

func TestSingleNnmber(t *testing.T) {
	helper.AssertEqual(t, singleNumber([]int{2, 2, 1}), 1)
	helper.AssertEqual(t, singleNumber([]int{4, 1, 2, 1, 2}), 4)
}
