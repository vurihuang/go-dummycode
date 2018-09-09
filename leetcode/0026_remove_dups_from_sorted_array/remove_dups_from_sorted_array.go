package pkg0026

func removeDuplicatesV1(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}

	prev := 0
	for i := 1; i < len(nums); i++ {
		if nums[prev] != nums[i] {
			prev++
			nums[prev] = nums[i]
		}
	}
	return prev + 1
}

func removeDuplicatesV2(nums []int) int {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[i-count] = nums[i]
		} else {
			count++
		}
	}
	return len(nums) - count
}
