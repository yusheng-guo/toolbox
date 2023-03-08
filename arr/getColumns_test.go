package arr

import (
	"fmt"
	"testing"
)

func TestGetColumns(t *testing.T) {
	arr := [][]any{
		{1, '3', 1},
		{1, 5, 1},
		{'4', 2, '1'},
	}
	ret := GetColumns(arr, 0, 3)
	fmt.Println(ret)
}
