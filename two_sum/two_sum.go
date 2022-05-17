package two_sum

func twoSum1(numbers []int, target int) []int {

	f, l := 0, len(numbers)-1

	for true {
		total := numbers[f] + numbers[l]

		if total == target {
			return []int{f + 1, l + 1}
		} else if total < target {
			f = f + 1
		} else {
			l = l - 1
		}
	}

	return []int{}
}

func twoSum(numbers []int, target int) []int {
	for idx, num := range numbers {
		remain := target - num

		anotherIndex := findNum(numbers, remain, idx+1)
		if anotherIndex > 0 {
			return []int{idx + 1, anotherIndex + 1}
		}
	}
	return []int{}
}

func findNum(nums []int, target, startIndex int) int {
	return searchByPosition(nums, target, startIndex, len(nums)-1)
}

func searchByPosition(nums []int, target, start, end int) int {
	if start > end {
		return -1
	}

	middle := (start + end) / 2
	if nums[middle] == target {
		return middle
	}

	if nums[middle] > target {
		return searchByPosition(nums, target, start, middle-1)
	}

	return searchByPosition(nums, target, middle+1, end)
}
