package reverse

func ReverseUint(x int) int {
	ret := 0
	for x != 0 {
		ret = ret*10 + x%10
		x /= 10
	}
	return ret
}

func ReverseString(s string) string {
	str := []byte(s)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	return string(str)
}
