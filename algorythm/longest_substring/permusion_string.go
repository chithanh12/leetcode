package longest_substring

func checkInclusion(s1 string, s2 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	if l2 < l1 {
		return false
	}

	arr1 := make([]byte, 26, 26)
	arr2 := make([]byte, 26, 26)
	for i := 0; i < len(s1); i++ {
		arr1[s1[i]-byte('a')]++
		arr2[s2[i]-byte('a')]++
	}

	if l2 == l1 {
		return sliceEqual(arr1, arr2)
	}

	for i := 0; i < l2-l1; i++ {
		if sliceEqual(arr1, arr2) {
			return true
		}

		arr2[s2[i+l1]-byte('a')]++
		arr2[s2[i]-byte('a')]--
	}

	return sliceEqual(arr1, arr2)
}

func sliceEqual(arr1, arr2 []byte) bool {
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}
