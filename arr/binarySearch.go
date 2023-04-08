package arr

// BinarySearch Bifurcation Find Target Element
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// BinarySearchFirst Bifurcate to find the first occurrence of the target value
func BinarySearchFirst(arr []int, target int) int {
	left, right := 0, len(arr)-1
	firstPos := -1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			firstPos = mid
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return firstPos
}

// BinarySearchLast Bifurcate to find the last occurrence of the target value
func BinarySearchLast(arr []int, target int) int {
	left, right := 0, len(arr)-1
	lastPos := -1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			lastPos = mid
			left = mid + 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return lastPos
}
