package movements

func Move(direction Direction, values [][]int) {
	switch direction {
	case UP:
		moveUp(values)
	case DOWN:
		moveDown(values)
	case LEFT:
		moveLeft(values)
	case RIGHT:
		moveRight(values)
	}
}
