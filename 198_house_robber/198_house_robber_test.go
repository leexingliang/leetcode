package house_robber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_robber_1(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	result := rob(nums)
	assert.Equal(t, 4, result)
}

func Test_robber_2(t *testing.T) {
	nums := []int{9, 8, 9, 8, 6, 9}
	result := rob(nums)
	assert.Equal(t, 27, result)
}
