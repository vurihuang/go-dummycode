package dominant_index

import (
	"testing"

	"github.com/upeoe/go-dummycode/helper"
)

func TestDominantIndex(t *testing.T) {
	helper.AssertEqual(t, dominantIndex([]int{3, 6, 1, 0}), 1)
	helper.AssertEqual(t, dominantIndex([]int{1, 2, 3, 4}), -1)
	helper.AssertEqual(t, dominantIndex([]int{1, 2, 3, 4}), -1)
	helper.AssertEqual(t, dominantIndex([]int{1, 0}), 0)
	helper.AssertEqual(t, dominantIndex([]int{0, 0, 1}), 2)
	helper.AssertEqual(t, dominantIndex([]int{0, 0, 1, 2}), 3)

}
