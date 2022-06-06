package day12

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	if len(nums) == 3 {
		return max(nums[0]+nums[2], nums[1])
	}
	tmp := make([]int, len(nums), len(nums))
	tmp[0] = nums[0]
	tmp[1] = max(nums[0], nums[1])
	tmp[2] = rob(nums[:3])

	for i := 3; i < len(nums); i++ {
		tmp[i] = max(tmp[i-3]+nums[i-1], tmp[i-2]+nums[i])
	}

	return tmp[len(nums)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
