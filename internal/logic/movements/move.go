package movements

func Move(direction Direction, values [][]int) {
	switch direction {
	case UP:
		moveUp(values)
	case DOWN:
		println("DOWN")
	case LEFT:
		moveLeft(values)
	case RIGHT:
		moveRight(values)
	}
}
