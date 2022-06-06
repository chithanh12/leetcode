package day12

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	ll := make([]int, n, n)
	ll[0] = 1
	ll[1] = 2
	for i := 2; i < n; i++ {
		ll[i] = ll[i-1] + ll[i-2]
	}
	return ll[n-1]
}
