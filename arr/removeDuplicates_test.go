package arr

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	ret1 := RemoveDuplicatesByLoop([]int{1, 1, 1, 2, 3, 5, 1})
	fmt.Println(ret1)
	ret2 := RemoveDuplicatesByMap([]string{"a", "b", "ab", "afds", "fads", "b", "b"})
	fmt.Println(ret2)

	ret3 := RemoveDuplicatesByLoop([]string{"a", "b", "ab", "afds", "fads", "b", "b"})
	fmt.Println(ret3)
	ret4 := RemoveDuplicatesByMap([]rune("adsga"))
	fmt.Println(string(ret4))
}
