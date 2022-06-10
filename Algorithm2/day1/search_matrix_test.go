package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	items := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	pos := findRow(items, 0, 2, 11)
	assert.Equal(t, 1, pos)

	pos = findRow([][]int{{1}, {3}}, 0, 1, 3)
	assert.Equal(t, 1, pos)
}