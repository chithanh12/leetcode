package binary_search

func search(nums []int, target int) int {
	return searchByPosition(nums, target, 0, len(nums)-1)
}

func searchByPosition(nums []int, target, start, end int) int {
	if start > end {
		return -1
	}

	middle := (start + end) / 2
	if nums[middle] == target {
		return middle
	}

	if nums[middle] > target {
		return searchByPosition(nums, target, start, middle-1)
	}

	return searchByPosition(nums, target, middle+1, end)
}
