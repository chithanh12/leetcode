package breath_first_search

import "fmt"

func permute(nums []int) [][]int {
	fmt.Printf("nums =%v\n", nums)
	if len(nums) == 1 {
		return [][]int{nums}
	}
	val := nums[len(nums)-1]
	previous := permute(nums[:len(nums)-1])

	return insertPermute(val, previous)
}

func insertPermute(val int, nums [][]int) [][]int {
	rs := make([][]int, 0)
	for _, set := range nums {
		for i := 0; i <= len(set); i++ {
			tmp := make([]int, len(set), len(set))
			copy(tmp, set)
			rs = append(rs, insert(tmp, val, i))
		}
	}

	return rs
}

func insert(a []int, c int, i int) []int {
	return append(a[:i], append([]int{c}, a[i:]...)...)
}
