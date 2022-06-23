package day3

func backspaceCompare(s string, t string) bool {
	ss := trim([]byte(s))
	tt := trim([]byte(t))
	if len(ss) != len(tt) {
		return false
	}

	for i := 0; i < len(ss); i++ {
		if ss[i] != tt[i] {
			return false
		}
	}

	return true
}

func trim(ss []byte) []byte {
	i, j := 0, 0
	rs := make([]byte, len(ss), len(ss))

	for j < len(ss) {
		if ss[j] == '#' {
			i = max(i-1, 0)
		} else {
			rs[i] = ss[j]
			i++
		}
		j++
	}

	return rs[:i]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
