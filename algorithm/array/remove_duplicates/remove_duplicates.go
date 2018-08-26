package remove_dups

func removeDuplicatesV1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[index] != nums[i] {
			index++
			nums[index] = nums[i]
		}
	}
	return index + 1
}

func removeDuplicatesV2(nums []int) int {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			count++
		} else {
			nums[i-count] = nums[i]
		}
	}
	return len(nums) - count
}
