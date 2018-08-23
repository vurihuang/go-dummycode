package leetcode

func LowercaseToUppercase(c byte) byte {
	if c >= 'a' && c <= 'z' {
		c -= 32
	}
	return c
}
