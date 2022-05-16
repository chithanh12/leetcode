package binary_search

func searchInsert(nums []int, target int) int {
	return searchPos(nums, target, 0, len(nums)-1)
}

func searchPos(nums []int, target, start, end int) int {
	if start > end {
		return -1
	}

	middle := (start + end) / 2
	if nums[middle] == target {
		return middle
	}

	if nums[middle] > target {
		f := searchPos(nums, target, start, middle-1)
		if f < 0 {
			return start
		}

		return f
	}

	l := searchPos(nums, target, middle+1, end)
	if l < 0 {
		return end + 1
	}

	return l
}
