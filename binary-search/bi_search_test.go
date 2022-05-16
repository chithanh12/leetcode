package binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 2, 3, 7, 9}
	f := search(nums, 3)
	assert.Equal(t, 2, f, "Not found at pos 2")

	f = search(nums, 0)
	assert.Equal(t, -1, f, "Not found. Should return -1")

	f = search(nums, 1)
	assert.Equal(t, 0, f, "Should return at begining")

	f = search(nums, 9)
	assert.Equal(t, 4, f, "Should return at the end")
}
