package arr

func ContainString(strs []string, s string) bool {
	for _, v := range strs {
		if v == s {
			return true
		}
	}
	return false
}

func ContainInt(nums []int, i int) bool {
	for _, v := range nums {
		if v == i {
			return true
		}
	}
	return false
}

func Contain[T comparable](arr []T, z T) bool {
	for _, v := range arr {
		if v == z {
			return true
		}
	}
	return false
}
