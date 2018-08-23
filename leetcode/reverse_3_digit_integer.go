package leetcode

func reverseInteger(num int) int {
	sum := 0
	for num > 0 {
		sum = sum*10 + num%10
		num /= 10
	}
	return sum
}
