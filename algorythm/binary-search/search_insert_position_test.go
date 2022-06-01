package binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	assert.Equal(t, 2, searchInsert(nums, 5), "wrong")
	assert.Equal(t, 1, searchInsert(nums, 2), "wrong")
	assert.Equal(t, 4, searchInsert(nums, 7), "wrong")

}
