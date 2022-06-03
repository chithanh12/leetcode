package breath_first_search

import "fmt"

/*
[1,0,1,0,0,0,1,1,1,1],
[1,1,1,0,1,1,1,1,0,1],
[1,0,1,1,1,1,0,1,0,0],
[1,0,1,1,1,0,1,1,1,1],
[1,1,0,1,1,1,1,0,0,0],
[1,1,0,0,1,0,1,1,0,1],
[1,1,1,1,1,1,1,1,1,1],
[1,1,0,0,0,1,1,1,0,0],
[0,1,1,1,0,0,1,0,1,1],
[1,1,0,0,0,1,0,1,1,0]
*/

type point struct {
	x int
	y int
}

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	tracking := make([]bool, m*n, m*n)
	untracking := make(map[int]*point)
	printMat(mat)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				tracking[i*n+j] = true
			} else {
				untracking[i*n+j] = &point{x: j, y: i}
			}
		}
	}

	for len(untracking) > 0 {
		removeKeys := []int{}
		for key, p := range untracking {
			found := false
			val := m + n

			if p.y-1 >= 0 && tracking[(p.y-1)*n+p.x] {
				val = mat[p.y-1][p.x] + 1
				found = true
			}

			if p.y+1 <= m-1 && tracking[(p.y+1)*n+p.x] {
				val = min(val, mat[p.y+1][p.x]+1)
				found = true
			}

			if p.x-1 >= 0 && tracking[p.y*n+p.x-1] {
				val = min(val, mat[p.y][p.x-1]+1)
				found = true
			}

			if p.x+1 <= n-1 && tracking[p.y*n+p.x+1] {
				val = min(val, mat[p.y][p.x+1]+1)
				found = true
			}

			if found {
				mat[p.y][p.x] = val
				tracking[p.y*n+p.x] = true
				removeKeys = append(removeKeys, key)
			}
		}

		for _, k := range removeKeys {
			delete(untracking, k)
		}

		printMat(mat)
	}

	return mat
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func printMat(mat [][]int) {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			fmt.Printf("%v  ", mat[i][j])
		}
		fmt.Println()
	}
	fmt.Println("------------------------------------")
}
