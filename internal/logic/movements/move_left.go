package movements

func moveLeft(values [][]int) {
	for row := range len(values) {
		newRow := make([]int, len(values[row]))
		insertAt := 0

		for col := range len(values[row]) {
			current := values[row][col]

			if current == 0 {
				continue
			}

			if insertAt == 0 {
				newRow[insertAt] = current
				insertAt++
				continue
			}

			if newRow[insertAt-1] == current {
				newRow[insertAt-1] = current * 2
				insertAt++
			} else {
				if newRow[insertAt-1] == 0 {
					newRow[insertAt-1] = current
					continue
				}

				newRow[insertAt] = current
				insertAt++
			}
		}

		values[row] = newRow
	}
}
