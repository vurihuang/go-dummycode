package pkg0003

func lengthOfLongestSubstringV1(s string) int {
	m, max := make(map[string]string), 0

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			tmp := s[i : j+1]
			if _, ok := m[tmp]; !ok && isUnique(tmp) {
				m[tmp] = tmp
				if len(tmp) > max {
					max = len(tmp)
				}
			}
		}
	}
	return max
}

func isUnique(s string) bool {
	m := make(map[byte]int)
	for i := range s {
		if _, ok := m[s[i]]; ok {
			return false
		}
		m[s[i]] = i
	}
	return true
}

func lengthOfLongestSubstringV2(s string) int {
	m, max := make(map[byte]int), 0

	for i := range s {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = i
			if len(m) > max {
				max = len(m)
			}
		} else {
			tmp := m[s[i]]
			m[s[i]] = i
			for k, v := range m {
				if v < tmp {
					delete(m, k)
				}
			}
		}
	}
	return max
}

func lengthOfLongestSubstringV3(s string) int {
	m := make(map[byte]int)
	max, i := 0, 0

	for j := range s {
		if _, ok := m[s[j]]; ok {
			i = larger(m[s[j]], i)
		}
		max = larger(max, j-i+1)
		m[s[j]] = j + 1
	}

	return max
}

func lengthOfLongestSubstringV4(s string) int {
	m := [128]int{}
	max, i := 0, 0

	for j := range s {
		i = larger(m[s[j]], i)
		max = larger(max, j-i+1)
		m[s[j]] = j + 1
	}

	return max
}

func larger(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
