package breath_first_search

func maxAreaOfIsland(grid [][]int) int {
	track := make([][]bool, len(grid), len(grid))
	for i := 0; i < len(grid); i++ {
		track[i] = make([]bool, len(grid[0]), len(grid[0]))
	}
	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != 0 && !track[i][j] {
				c := area(grid, i, j, track)
				if max < c {
					max = c
				}
			}
		}
	}
	return max
}

func area(grid [][]int, r, c int, track [][]bool) int {
	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) {
		return 0
	}

	if grid[r][c] == 0 {
		return 0
	}

	if track[r][c] {
		return 0
	}

	track[r][c] = true
	t := area(grid, r-1, c, track)
	d := area(grid, r+1, c, track)
	l := area(grid, r, c+1, track)
	ri := area(grid, r, c-1, track)

	return 1 + t + d + l + ri
}
