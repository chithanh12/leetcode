package day1

func searchRange(nums []int, target int) []int {
	first, second := -1, -1
	fs, sf := -1, -1

	for i, j := 0, len(nums)-1; i <= j; i, j = i+1, j-1 {
		if nums[i] == target {
			if first == -1 {
				first = i
			} else {
				sf = i
			}
		}

		if nums[j] == target {
			if second == -1 {
				second = j
			} else {
				fs = j
			}
		}

		if first >= 0 && second >= 0 {
			break
		}
	}

	if first == -1 && fs != -1 {
		first = fs
	}
	if second == -1 && sf != -1 {
		second = sf
	}

	if (second == -1 && first != -1) || (first == -1 && second != -1) {
		first = second + first + 1
		second = first
	}

	return []int{first, second}
}
