package arr

// GetRows 获取几行
func GetRows(arr [][]any, start, end int) [][]any {
	if start > end || end > len(arr) {
		panic("GetRows Err")
	}
	return arr[start:end]
}
