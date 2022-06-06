package day12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTriangle(t *testing.T) {
	nums := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	m := minimumTotal(nums)
	assert.Equal(t, 11, m)

}
