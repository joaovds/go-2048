package movements

func Move(direction Direction, values [][]int) (newPoints int) {
	switch direction {
	case UP:
		newPoints = moveUp(values)
	case DOWN:
		newPoints = moveDown(values)
	case LEFT:
		newPoints = moveLeft(values)
	case RIGHT:
		newPoints = moveRight(values)
	}

	return
}
