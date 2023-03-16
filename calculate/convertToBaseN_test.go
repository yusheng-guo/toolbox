package calculate

import (
	"fmt"
	"testing"
)

func TestConvertToBaseN(t *testing.T) {
	ret := ConvertToBaseN(100, 7)
	fmt.Println(ret)
}
