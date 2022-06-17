package day3

func threeSum(nums []int) [][]int {
	//
	mm, values := toMap(nums)
	values = bubbleSort(values)
	rs := make([][]int, 0, 0)

	for idx, val := range values {
		cnt, _ := mm[val]
		if val == 0 && cnt >= 3 {
			rs = append(rs, []int{0, 0, 0})
		}

		for j := idx + 1; j < len(values); j++ {
			remain := -(val + values[j])
			if remain > values[j] {
				if _, ok := mm[remain]; ok {
					rs = append(rs, []int{val, values[j], remain})
				}
			} else {
				break
			}
		}
		if cnt >= 2 && val != 0 {
			if _, ok := mm[-2*val]; ok {
				rs = append(rs, []int{val, val, -2 * val})
			}
		}
	}

	return rs
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func toMap(nums []int) (map[int]int, []int) {
	distint := []int{}
	mm := make(map[int]int, 0)

	for _, val := range nums {
		if _, ok := mm[val]; !ok {
			mm[val] = 1
			distint = append(distint, val)
		} else {
			mm[val] += 1
		}
	}

	return mm, distint
}

func bubbleSort(sums []int) []int {
	for i := 0; i < len(sums)-1; i++ {
		for j := i + 1; j < len(sums); j++ {
			if sums[i] > sums[j] {
				sums[i], sums[j] = sums[j], sums[i]
			}
		}
	}

	return sums
}
