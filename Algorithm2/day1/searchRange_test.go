package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchRange(t *testing.T) {
	rs := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	assert.Equal(t, []int{3, 4}, rs)
}
