package two_pointers

func rotate(nums []int, k int) {
	k = k % len(nums)
	tmp := make([]int, k, k)
	l := len(nums)
	for i := 0; i < k; i++ {
		tmp[i] = nums[l-k+i]
	}

	for i := l - 1; i >= k; i-- {
		nums[i] = nums[i-k]
	}

	for i := 0; i < k; i++ {
		nums[i] = tmp[i]
	}
}
