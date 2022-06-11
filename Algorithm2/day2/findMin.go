package day2

func findMin(nums []int) int {
	pos := findPivot(nums, 0, len(nums)-1)
	return nums[pos]
}

func findPivot(nums []int, start, end int) int {
	if start == end {
		return start
	}

	if end-start == 1 {
		if nums[end] > nums[start] {
			return start
		}
		return end
	}

	mid := (start + end) / 2
	if nums[mid] > nums[end] {
		return findPivot(nums, mid, end)
	}
	return findPivot(nums, start, mid)
}
