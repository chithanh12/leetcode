package day2

func findPeakElement(nums []int) int {
	start, end := 0, len(nums)-1
	if len(nums) == 1 {
		return 0
	}

	for start < end {
		mid := (start + end) / 2
		if nums[mid] > nums[mid+1] {
			end = mid
		} else {
			start = mid
		}
	}
	return start
}
