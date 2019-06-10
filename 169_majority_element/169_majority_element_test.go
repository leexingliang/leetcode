package majority_element

import (
	"fmt"
	"testing"
)

func Test_MajorityElem(t *testing.T) {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))

	nums = []int{3}
	fmt.Println(majorityElement(nums))
}
