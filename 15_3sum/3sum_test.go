package sum

import (
	"fmt"
	"testing"
)

func Test_First(t *testing.T) {
	src := []int{-1, 0, 1, 2, -1, -4}

	ret := threeSum(src)
	fmt.Println(ret)
}
