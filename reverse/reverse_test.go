package reverse

import (
	"fmt"
	"testing"
)

func TestReverseUint(t *testing.T) {
	ret := ReverseUint(12345)
	fmt.Println(ret)
}

func TestReverseString(t *testing.T) {
	ret := ReverseString("abcdef")
	fmt.Println(ret)
}
