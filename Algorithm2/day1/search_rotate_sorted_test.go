package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotate(t *testing.T) {

	p := findPivot([]int{3, 4, 5, 1, 2}, 0, 4)
	assert.Equal(t, 3, p)

	tt := search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	assert.Equal(t, 4, tt)
	ttt := search([]int{5, 1, 3}, 3)
	assert.Equal(t, 2, ttt)
	xx := search([]int{1, 3}, 3)
	assert.Equal(t, 1, xx)

	yy := search([]int{1, 3, 5}, 1)
	assert.Equal(t, 0, yy)

}
