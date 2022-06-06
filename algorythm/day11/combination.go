package day11

func combine(n int, k int) [][]int {
	if n < k || k == 0 {
		return [][]int{}
	}
	if k == 1 {
		tmp := make([][]int, 0, 0)
		for idx := 1; idx <= n; idx++ {
			tmp = append(tmp, []int{idx})
		}
		return tmp
	}

	if n == k {
		return collection(k)
	}

	// C(n, k) = C (n-1, k) + insert(n , C(n-1, k-1))
	c1 := combine(n-1, k)
	c2 := addMore(n, combine(n-1, k-1))

	return append(c1, c2...)
}

func collection(n int) [][]int {
	tmp := make([]int, 0)
	for i := 1; i <= n; i++ {
		tmp = append(tmp, i)
	}
	return [][]int{tmp}
}

func addMore(val int, items [][]int) [][]int {
	rs := make([][]int, 0)
	for _, item := range items {
		tmp := append(item, val)
		rs = append(rs, tmp)
	}

	return rs
}
