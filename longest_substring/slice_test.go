package longest_substring

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlice(t *testing.T) {
	num := []int{1, 2, 3, 4}
	num1 := num[0:2]
	fmt.Println(num1)
}

func TestImpmutation(t *testing.T) {
	assert.Equal(t, true, checkInclusion("ab", "eidbaooo"))
}
