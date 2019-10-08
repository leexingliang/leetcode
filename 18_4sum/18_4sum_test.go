package sum4

import (
	"fmt"
	"testing"
)

func Test_Sum_1(t *testing.T) {
	nums := []int{1, 0, -1, 0, -2, 2}

	result := fourSum(nums, 0)
	fmt.Println(result)
}
