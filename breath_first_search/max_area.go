package breath_first_search

func maxAreaOfIsland(grid [][]int) int {
	cCnt := len(grid[0])
	track := make([]bool, len(grid)*cCnt, len(grid)*cCnt)
	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != 0 && !track[i*cCnt+j] {
				c := area(grid, i, j, track, cCnt)
				if max < c {
					max = c
				}
			}
		}
	}
	return max
}

func area(grid [][]int, r, c int, track []bool, cCnt int) int {
	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) {
		return 0
	}

	if grid[r][c] == 0 {
		return 0
	}

	if track[r*cCnt+c] {
		return 0
	}

	track[r*cCnt+c] = true
	t := area(grid, r-1, c, track, cCnt)
	d := area(grid, r+1, c, track, cCnt)
	l := area(grid, r, c+1, track, cCnt)
	ri := area(grid, r, c-1, track, cCnt)

	return 1 + t + d + l + ri
}
