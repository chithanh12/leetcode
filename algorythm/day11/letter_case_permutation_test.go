package day11

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLetterCasePermutation(t *testing.T) {
	rs := letterCasePermutation("a1b2")
	fmt.Println(rs)

	assert.Equal(t, true, sliceStringEqual([]string{"a1b2", "a1B2", "A1b2", "A1B2"}, rs))
}

func sliceStringEqual(s, t []string) bool {
	if len(s) != len(t) {
		return false
	}

	for _, ss := range s {
		if !contains(ss, t) {
			return false
		}
	}
	return true
}

func contains(s string, ss []string) bool {
	for _, sss := range ss {
		if s == sss {
			return true
		}
	}
	return false
}

func TestSliceByte(t *testing.T) {
	item := []byte{'1'}
	item = append(item, '2')
	fmt.Println(item)
}
