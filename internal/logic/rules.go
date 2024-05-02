package logic

func IsGameOver(values [][]int) bool {
	if HasEmptyCell(values) {
		return false
	}

	for row := range SIZE {
		if row == 0 {
			for col := range SIZE {
				if col != SIZE-1 {
					if isEqualRight(values, row, col) {
						return false
					}
				}
			}
		} else {
			for col := range SIZE {
				if col == SIZE-1 {
					if isEqualTop(values, row, col) {
						return false
					}
				} else {
					if isEqualRight(values, row, col) || isEqualTop(values, row, col) {
						return false
					}
				}
			}
		}
	}

	return true
}

func isEqualTop(values [][]int, row, col int) bool {
	return values[row][col] == values[row-1][col]
}

func isEqualRight(values [][]int, row, col int) bool {
	return values[row][col] == values[row][col+1]
}

func HasEmptyCell(values [][]int) bool {
	for row := range SIZE {
		for col := range SIZE {
			if values[row][col] == 0 {
				return true
			}
		}
	}

	return false
}

func changed(values, newValues [][]int) bool {
	if len(values) != len(newValues) {
		return true
	}

	for row := range len(values) {
		if len(values[row]) != len(newValues[row]) {
			return true
		}

		for col := range len(values[row]) {
			if values[row][col] != newValues[row][col] {
				return true
			}
		}
	}

	return false
}
