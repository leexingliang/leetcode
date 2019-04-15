package remove

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Remove(t *testing.T) {
	src := []int{1, 1}

	ret := removeDuplicates(src)
	assert.Equal(t, 1, ret)
}
