package day12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHouseRober(t *testing.T) {
	tt := rob([]int{1, 2, 3, 1})
	assert.Equal(t, 4, tt)
}
