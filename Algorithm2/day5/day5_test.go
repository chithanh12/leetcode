package day5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAnagram(t *testing.T) {
	s, p := "cbaebabacd", "abc"
	items := findAnagrams(s, p)
	assert.Equal(t, []int{0, 6}, items)
}
