package arr

import (
	"fmt"
	"testing"
)

func TestContainInt(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	ret := ContainInt(nums, 3)
	fmt.Println(ret)
}

func TestContain(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	fmt.Println(Contain(nums, 3))
	strs := []string{"fdas", "fa", "sagger", "thr"}
	fmt.Println(Contain(strs, "fa"))
}

func TestContainString(t *testing.T) {
	strs := []string{"fdas", "fa", "sagger", "thr"}
	ret := ContainString(strs, "fa")
	fmt.Println(ret)
}
