package movements

func moveRight(values [][]int) {
	for row := range len(values) {
		newRow := make([]int, len(values[row]))
		insertAt := len(newRow) - 1

		for col := len(values[row]) - 1; col >= 0; col-- {
			current := values[row][col]

			if current == 0 {
				continue
			}

			if insertAt == len(newRow)-1 {
				newRow[insertAt] = current
				insertAt--
				continue
			}

			if newRow[insertAt+1] == current {
				newRow[insertAt+1] = current * 2
			} else {
				newRow[insertAt] = current
				insertAt--
			}
		}

		values[row] = newRow
	}
}
