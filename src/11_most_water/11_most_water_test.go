package most_water

import "testing"
import "github.com/stretchr/testify/assert"

func Test_MaxArea(t *testing.T) {
	var src []int = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	assert.Equal(t, 49, maxArea(src))
}
