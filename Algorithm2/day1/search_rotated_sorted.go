package day1

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	pivote := findPivot(nums, 0, len(nums)-1)

	if nums[pivote] == target {
		return pivote
	}

	if nums[len(nums)-1] >= target {
		return findPos(nums, pivote+1, len(nums)-1, target)
	}

	return findPos(nums, 0, pivote-1, target)
}

func findPos(nums []int, start, end, target int) int {
	if start > end {
		return -1
	}

	mid := (start + end) / 2
	if nums[mid] == target {
		return mid
	}

	if nums[mid] > target {
		return findPos(nums, start, mid-1, target)
	}

	return findPos(nums, mid+1, end, target)
}

func findPivot(nums []int, start, end int) int {
	if start == end {
		return start
	}

	if nums[start] < nums[end] {
		return start
	}

	mid := (start + end) / 2

	if nums[start] < nums[mid] {
		return findPivot(nums, mid, end)
	}

	return findPivot(nums, start, mid)
}
