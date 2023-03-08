package arr

// GetColumns 获取几列
func GetColumns(arr [][]any, start, end int) [][]any {
	if start > end || end > len(arr[0]) {
		panic("GetRows Err")
	}
	tmp := make([][]any, len(arr))
	for i := 0; i < len(arr); i++ {
		for j := start; j < end; j++ {
			tmp[i] = append(tmp[i], arr[i][j])
		}
	}
	return tmp
}
