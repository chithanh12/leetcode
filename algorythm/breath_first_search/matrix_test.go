package breath_first_search

import (
	"fmt"
	"testing"
)

func TestMatrix(t *testing.T) {
	mat := [][]int{
		{1, 1, 0, 0, 1, 0, 0, 1, 1, 0},
		{1, 0, 0, 1, 0, 1, 1, 1, 1, 1},
		{1, 1, 1, 0, 0, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 0, 1, 1, 1, 1, 1},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1, 1, 0, 1, 1, 1},
		{0, 1, 1, 1, 1, 1, 1, 0, 0, 1},
		{1, 1, 1, 1, 1, 0, 0, 1, 1, 1},
		{0, 1, 0, 1, 1, 0, 1, 1, 1, 1},
		{1, 1, 1, 0, 1, 0, 1, 1, 1, 1},
	}

	m := updateMatrix(mat)
	println(m)
}

func TestSliceIndesx(t *testing.T) {
	tmp := []int{0, 1, 2, 3}
	fmt.Println(tmp[:3])
	fmt.Println(tmp[:0])
}

func TestPermutation(t *testing.T) {
	tmp := []int{1, 2, 3}
	rs := permute(tmp)
	fmt.Println("--- result--")
	fmt.Println(rs)
}
