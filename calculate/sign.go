package calculate

// Sign Determine the positive and negative of a number.
func Sign(num int) int {
	if num == 0 {
		return 0
	}
	if num > 0 {
		return 1
	}
	return -1
}
