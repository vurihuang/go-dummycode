package single_num

func singleNumber(nums []int) int {
	ret := 0
	for _, x := range nums {
		ret ^= x
	}
	return ret
}
