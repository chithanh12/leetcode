package day12

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return triangle[0][0]
	}
	tracking := make([][]int, len(triangle), len(triangle))
	tracking[0] = triangle[0]
	//fmt.Println(triangle)
	for i := 1; i < len(triangle); i++ {
		tracking[i] = make([]int, len(triangle[i]), len(triangle[i]))
		for j := 1; j < len(triangle[i])-1; j++ {
			tracking[i][j] = min2(tracking[i-1][j-1]+triangle[i][j], tracking[i-1][j]+triangle[i][j])
		}

		tracking[i][0] = tracking[i-1][0] + triangle[i][0]
		tracking[i][len(triangle[i])-1] = tracking[i-1][len(tracking[i-1])-1] + triangle[i][len(triangle[i])-1]
	}

	//fmt.Println(tracking)

	return min(tracking[len(triangle)-1])
}

func min(items []int) int {
	tt := items[0]
	for i := 1; i < len(items); i++ {
		tt = min2(tt, items[i])
	}
	return tt
}

func min2(x, y int) int {
	if x > y {
		return y
	}

	return x
}
