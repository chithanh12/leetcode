package sort

func sort(num []int) []int {
	for i := 0; i < len(num)-1; i++ {
		for j := i + 1; j < len(num); j++ {
			if num[i] > num[j] {
				num[i], num[j] = num[j], num[i]
			}
		}
	}

	return num
}
