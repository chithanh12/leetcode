package move_zeros

func moveZeroes(nums []int) {
	countZero := 0
	for idx, num := range nums {
		if num == 0 {
			countZero++
			continue
		}
		nums[idx-countZero] = num
	}
	for i := 1; i <= countZero; i++ {
		nums[len(nums)-i] = 0
	}
}
