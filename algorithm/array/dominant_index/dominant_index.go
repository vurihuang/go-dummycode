package dominant_index

func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	max, secMax := 0, 1
	if nums[max] < nums[secMax] {
		max, secMax = secMax, max
	}

	for i := 2; i < len(nums); i++ {
		if nums[i] > nums[max] {
			max, secMax = i, max
		} else if nums[i] > nums[secMax] {
			secMax = i
		}
	}

	if nums[max] >= 2*nums[secMax] {
		return max
	}
	return -1
}
