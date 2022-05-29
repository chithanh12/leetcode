package breath_first_search

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldValue := image[sr][sc]
	track := make([][]bool, len(image), len(image))
	for i := 0; i < len(image); i++ {
		track[i] = make([]bool, len(image[0]), len(image[0]))
	}
	fill(image, sr, sc, oldValue, newColor, track)
	return image
}

func fill(imgs [][]int, r, c, oldColor, color int, track [][]bool) {
	if r < 0 || c < 0 || r >= len(imgs) || c >= len(imgs[0]) {
		return
	}

	if imgs[r][c] != oldColor {
		return
	}

	if track[r][c] {
		return
	}

	imgs[r][c] = color
	track[r][c] = true

	fill(imgs, r+1, c, oldColor, color, track)
	fill(imgs, r-1, c, oldColor, color, track)
	fill(imgs, r, c+1, oldColor, color, track)
	fill(imgs, r, c-1, oldColor, color, track)
}
