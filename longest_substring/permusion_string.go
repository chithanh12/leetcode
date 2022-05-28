package longest_substring

func checkInclusion(s2 string, s1 string) bool {
	tmp := map[byte]struct{}{}
	for i := 0; i < len(s2); i++ {
		tmp[s2[i]] = struct{}{}
	}
	for i := 0; i < len(s1)-len(s2); i++ {
		if _, ok := tmp[s1[i]]; !ok {
			continue
		}

		if isImutable(s2, s1, i) {
			return true
		}
	}
	return false
}

func isImutable(s2, s1 string, start int) bool {
	tmp := map[byte]int{}
	tmp1 := map[byte]int{}

	for idx := 0; idx < len(s2); idx++ {
		tmp[s1[idx+start]] += 1
		tmp1[s2[idx]] += 1
	}

	for k, v1 := range tmp {
		if v2, ok := tmp1[k]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}
