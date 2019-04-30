package maximum_subarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MaxSubArray1(t *testing.T) {
	src := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	ret := maxSubArray(src)
	assert.Equal(t, 6, ret)
}

func Test_MaxSubArray2(t *testing.T) {
	src := []int{-1}

	ret := maxSubArray(src)
	assert.Equal(t, -1, ret)
}
