package pivot_index

func pivotIndex(nums []int) int {
	sum, halfSum := 0, 0
	for _, n := range nums {
		sum += n
	}
	for i, n := range nums {
		if sum-n == 2*halfSum {
			return i
		}
		halfSum += n
	}
	return -1
}
