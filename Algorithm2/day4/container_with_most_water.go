package day4

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	maxA := 0
	for i < j {
		s := (j - i) * min(height[i], height[j])
		if maxA < s {
			maxA = s
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return maxA
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
