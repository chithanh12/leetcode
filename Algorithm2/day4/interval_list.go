package day4

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	i, j := 0, 0
	rs := make([][]int, 0, 0)
	for true {
		if i >= len(firstList) || j >= len(secondList) {
			break
		}
		intersec, ok, lst := intersection(firstList[i], secondList[j])
		if ok {
			rs = append(rs, intersec)
		}

		if lst == 1 {
			j++
		} else {
			i++
		}
	}

	return rs
}

func intersection(first, second []int) ([]int, bool, int) {
	if first[0] < second[0] {
		if first[1] < second[0] {
			return nil, false, 2
		}

		if first[1] <= second[1] {
			return []int{second[0], first[1]}, true, 2
		}

		return second, true, 1
	}

	if first[0] <= second[1] {
		if first[1] <= second[1] {
			return first, true, 2
		}
		return []int{first[0], second[1]}, true, 1
	}
	return nil, false, 1
}
