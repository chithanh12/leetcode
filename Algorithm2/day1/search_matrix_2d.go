package day1

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	row := findRow(matrix, 0, len(matrix)-1, target)
	if row == -1 {
		return false
	}

	return biSearch(matrix[row], 0, len(matrix[row])-1, target)
}

func biSearch(row []int, start, end, target int) bool {
	if start > end {
		return false
	}
	if start == end {
		return row[start] == target
	}

	mid := (start + end) / 2
	if row[mid] == target {
		return true
	}

	if row[mid] > target {
		return biSearch(row, start, mid-1, target)
	}

	return biSearch(row, mid+1, end, target)
}

func findRow(matrix [][]int, startRow, endRow, target int) int {
	if startRow >= endRow {
		return startRow
	}
	if startRow == endRow-1 {
		if matrix[endRow][0] <= target {
			return endRow
		}
		return startRow
	}

	mid := (startRow + endRow) / 2

	if matrix[mid][0] == target {
		return mid
	}

	if matrix[mid][0] > target {
		return findRow(matrix, startRow, mid-1, target)
	}

	return findRow(matrix, mid, endRow, target)
}
