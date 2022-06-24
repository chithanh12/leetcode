package day5

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) || len(p) == 0 {
		return []int{}
	}

	second := characterCount(p)
	first := characterCount(string(s[:len(p)]))
	result := make([]int, 0)
	if uint8SliceEqual(first, second) {
		result = append(result, 0)
	}

	start := 0
	for i := len(p); i < len(s); i++ {
		first[int(s[start]-97)]--
		first[int(s[i]-97)]++
		if uint8SliceEqual(first, second) {
			result = append(result, i-len(p)+1)
		}
		start++
	}

	return result
}

func uint8SliceEqual(first, second []uint8) bool {
	for i := 0; i < 26; i++ {
		if first[i] != second[i] {
			return false
		}
	}

	return true
}

func characterCount(s string) []uint8 {
	result := make([]uint8, 26, 26)
	for _, ss := range s {
		result[int(ss)-97] += 1
	}
	return result
}
