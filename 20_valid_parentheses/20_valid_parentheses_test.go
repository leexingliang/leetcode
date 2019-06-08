package valid_parentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsValid(t *testing.T) {
	s := "([{}])[]{}()"

	result := isValid(s)
	assert.True(t, result)
}
