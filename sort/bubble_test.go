package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubble(t *testing.T) {
	num := []int{4, 2, 7, 5, 1}
	sort(num)
	assert.Equal(t, []int{1, 2, 4, 5, 7}, num)

	num = []int{9, -1, 4, 5, 2, 3, 8}
	sort(num)

	assert.Equal(t, []int{-1, 2, 3, 4, 5, 8, 9}, num)
}
