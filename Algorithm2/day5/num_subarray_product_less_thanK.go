package day5

func numSubarrayProductLessThanK(nums []int, k int) int {
	left := 0
	total, count := 1, 0

	for right := 0; right < len(nums); right++ {
		total *= nums[right]

		for total >= k && left <= right {
			total = total / nums[left]
			left++
		}

		count += (right - left + 1)
	}
	return count
}
