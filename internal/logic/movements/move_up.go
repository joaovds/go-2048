package movements

func moveUp(values [][]int) (newPoints int) {
	for col := range len(values[0]) {
		newCol := make([]int, len(values[0]))
		insertAt := 0

		for row := range len(values) {
			current := values[row][col]

			if current == 0 {
				continue
			}

			if insertAt == 0 {
				newCol[insertAt] = current
				insertAt++
				continue
			}

			if newCol[insertAt-1] == current {
				newCol[insertAt-1] = current * 2
				newPoints += current * 2
				insertAt++
			} else {
				if newCol[insertAt-1] == 0 {
					newCol[insertAt-1] = current
					continue
				}

				newCol[insertAt] = current
				insertAt++
			}
		}

		for row := range len(values) {
			values[row][col] = newCol[row]
		}
	}

	return
}
