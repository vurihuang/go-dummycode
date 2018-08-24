package pkg0001

func twoSumV1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumV2(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, n := range nums {
		m[n] = i
	}

	for i, n := range nums {
		if j, ok := m[target-n]; ok {
			return []int{i, j}
		}
	}
	return nil
}

func twoSumV3(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, n := range nums {
		if j, ok := m[target-n]; ok {
			return []int{j, i}
		}
		m[n] = i
	}
	return nil
}
