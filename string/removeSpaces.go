package string

// RemoveSpaces 去除前后空字符
func RemoveSpaces(s string) string {
	left, right := 0, len(s)-1
	for right >= 0 && s[right] == ' ' {
		right--
	}
	for left < right && s[left] == ' ' {
		left++
	}
	return s[left : right+1]
}
