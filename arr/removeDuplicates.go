package arr

// RemoveDuplicatesByLoop Filtering duplicate elements by double loops
//   Pros: No change in order after de-duplication
//   Disadvantage: slow de-duplication speed
func RemoveDuplicatesByLoop[T comparable](slc []T) []T {
	result := []T{}
	for i := range slc {
		repeat := false
		for j := range result {
			if slc[i] == result[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			result = append(result, slc[i])
		}
	}
	return result
}

// RemoveDuplicatesByMap Filtering duplicate elements using map
//   Pros: Fast speed
//   Disadvantages: disorder occurs
func RemoveDuplicatesByMap[T comparable](slc []T) []T {
	result := []T{}
	tempMap := map[T]bool{}
	for _, item := range slc {
		if _, ok := tempMap[item]; !ok {
			tempMap[item] = true
			result = append(result, item)
		}
	}
	return result
}
