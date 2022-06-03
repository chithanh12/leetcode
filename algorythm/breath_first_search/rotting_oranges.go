package breath_first_search

type pos struct {
	row int
	col int
}

func orangesRotting(grid [][]int) int {
	countFresh := 0
	rotting := make([]*pos, 0)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				rotting = append(rotting, &pos{col: j, row: i})
			} else if grid[i][j] == 1 {
				countFresh++
			}
		}
	}
	// there no fresh
	if countFresh == 0 {
		return 0
	}

	minute := 0
	for true {
		newNodes := make([]*pos, 0)
		for _, position := range rotting {
			if n := turnOn(position.row-1, position.col, grid); n != nil {
				newNodes = append(newNodes, n)
			}

			if n := turnOn(position.row+1, position.col, grid); n != nil {
				newNodes = append(newNodes, n)
			}

			if n := turnOn(position.row, position.col+1, grid); n != nil {
				newNodes = append(newNodes, n)
			}

			if n := turnOn(position.row, position.col-1, grid); n != nil {
				newNodes = append(newNodes, n)
			}

		}
		rotting = newNodes
		if len(newNodes) > 0 {
			minute++
			countFresh -= len(newNodes)
		}
		if len(newNodes) == 0 || countFresh <= 0 {
			break
		}
	}

	if countFresh > 0 {
		return -1
	}

	return minute
}

func turnOn(r, c int, grid [][]int) *pos {
	if r < 0 || c < 0 || r > len(grid)-1 || c > len(grid[0])-1 {
		return nil
	}

	if grid[r][c] == 1 {
		grid[r][c] = 2
		return &pos{
			row: r,
			col: c,
		}
	}

	return nil
}
