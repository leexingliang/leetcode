package max_profit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test01(t *testing.T) {
	src := []int{3, 2, 6, 5, 0, 3}

	ret := maxProfit(src)
	assert.Equal(t, 4, ret)
}

func Test02(t *testing.T) {
	src := []int{7, 1, 5, 3, 6, 4}

	ret := maxProfit(src)
	assert.Equal(t, 5, ret)
}
