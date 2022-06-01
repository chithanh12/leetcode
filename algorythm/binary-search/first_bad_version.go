package binary_search

func firstBadVersion(n int) int {
	if isBadVersion(1) {
		return 1
	}

	return badVersion(1, n-1)
}

func badVersion(start int, end int) int {
	if isBadVersion(start) {
		return start
	}

	if !isBadVersion(end) {
		return end + 1
	}

	middle := (start + end) / 2
	if isBadVersion(middle) {
		if middle-1 >= 1 && !isBadVersion(middle-1) {
			return middle
		}

		return badVersion(start+1, middle-1)
	}

	return badVersion(middle+1, end-1)
}

func isBadVersion(version int) bool {
	//NOPE
	return true
}
