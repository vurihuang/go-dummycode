package pkg0027

func removeElementV1(nums []int, val int) int {
	ret := 0
	for _, n := range nums {
		if n != val {
			nums[ret] = n
			ret++
		}
	}
	return ret
}

func removeElementV2(nums []int, val int) int {
	ret, l := 0, len(nums)
	for ret < l {
		if nums[ret] == val {
			nums[ret] = nums[l-1]
			l--
		} else {
			ret++
		}
	}
	return ret
}
