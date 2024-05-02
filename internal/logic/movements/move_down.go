package movements

func moveDown(values [][]int) {
	for col := range len(values[0]) {
		newCol := make([]int, len(values[0]))
		insertAt := len(newCol) - 1

		for row := len(values) - 1; row >= 0; row-- {
			current := values[row][col]

			if current == 0 {
				continue
			}

			if insertAt == len(newCol)-1 {
				newCol[insertAt] = current
				insertAt--
				continue
			}

			if newCol[insertAt+1] == current {
				newCol[insertAt+1] = current * 2
				insertAt--
			} else {
				if newCol[insertAt+1] == 0 {
					newCol[insertAt+1] = current
					continue
				}

				newCol[insertAt] = current
				insertAt--
			}
		}

		for row := range len(values) {
			values[row][col] = newCol[row]
		}
	}
}
