package longest_substring

func lengthOfLongestSubstring(s string) int {
	sub := make(map[byte]int, 0)
	currentMax := 0
	startCountAtIndex := 0

	for idx, ss := range s {
		if lastPos, ok := sub[byte(ss)]; ok && lastPos >= startCountAtIndex {
			if currentMax < (idx - startCountAtIndex) {
				currentMax = idx - startCountAtIndex
			}

			startCountAtIndex = sub[byte(ss)] + 1
		}

		sub[byte(ss)] = idx
	}
	if currentMax < (len(s) - startCountAtIndex) {
		currentMax = len(s) - startCountAtIndex
	}

	return currentMax
}
