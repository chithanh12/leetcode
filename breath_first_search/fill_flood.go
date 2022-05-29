package breath_first_search

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldValue := image[sr][sc]
	num := len(image) * len(image[0])
	colCnt := len(image[0])
	track := make([]bool, num, num)

	fill(image, sr, sc, oldValue, newColor, track, colCnt)
	return image
}

func fill(imgs [][]int, r, c, oldColor, color int, track []bool, colCount int) {
	if r < 0 || c < 0 || r >= len(imgs) || c >= len(imgs[0]) {
		return
	}

	if imgs[r][c] != oldColor {
		return
	}

	if track[r*colCount+c] {
		return
	}

	imgs[r][c] = color
	track[r*colCount+c] = true

	fill(imgs, r+1, c, oldColor, color, track, colCount)
	fill(imgs, r-1, c, oldColor, color, track, colCount)
	fill(imgs, r, c+1, oldColor, color, track, colCount)
	fill(imgs, r, c-1, oldColor, color, track, colCount)
}
