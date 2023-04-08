package arr

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 2, 6, 6, 6, 6, 7, 10}
	fmt.Println(BinarySearch(arr, 6))      // 4
	fmt.Println(BinarySearchFirst(arr, 6)) // 3
	fmt.Println(BinarySearchLast(arr, 6))  // 6
}
