package string

import (
	"fmt"
	"testing"
)

func TestRemoveSpaces(t *testing.T) {
	ret := RemoveSpaces("  Hello  ")
	fmt.Println(ret)
}
