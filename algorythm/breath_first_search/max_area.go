package breath_first_search

func maxAreaOfIsland(grid [][]int) int {
	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != 0 && grid[i][j] != -1 {
				c := area(grid, i, j)
				max = maxValue(max, c)
			}
		}
	}

	return max
}

func area(grid [][]int, r, c int) int {
	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) {
		return 0
	}

	if grid[r][c] == 0 || grid[r][c] == -1 {
		return 0
	}

	grid[r][c] = -1

	t := area(grid, r-1, c)
	d := area(grid, r+1, c)
	l := area(grid, r, c+1)
	ri := area(grid, r, c-1)

	return 1 + t + d + l + ri
}

func maxValue(x, y int) int {
	if x > y {
		return x
	}
	return y
}
