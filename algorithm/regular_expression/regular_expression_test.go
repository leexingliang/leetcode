package regular_expression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Pattern(t *testing.T) {
	Pattern("abcdabb*ooo")
	ret := isMatch("abcooo")
	assert.False(t, ret)

	ret = isMatch("abcdabboooooo")
	assert.True(t, ret)
}
