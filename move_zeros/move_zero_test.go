package move_zeros

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveZero(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	assert.Equal(t, []int{1, 3, 12, 0, 0}, nums)

	nums = []int{0, 0, 0, 0, 0, 1, 2, 3, 4}
	moveZeroes(nums)
	assert.Equal(t, []int{1, 2, 3, 4, 0, 0, 0, 0, 0}, nums)

	nums = []int{0, 0, 0, 0, 0, 1, 2, 3, 4, 0, 0}
	moveZeroes(nums)
	assert.Equal(t, []int{1, 2, 3, 4, 0, 0, 0, 0, 0, 0, 0}, nums)
}
