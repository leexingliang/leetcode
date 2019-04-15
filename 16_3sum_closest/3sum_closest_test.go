package sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_First(t *testing.T) {
	src := []int{-1, 2, 1, -4}

	ret := threeSumClosest(src, 2)
	assert.Equal(t, 2, ret)
}
