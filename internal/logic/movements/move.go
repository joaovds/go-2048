package movements

func Move(direction Direction, values [][]int) {
	switch direction {
	case UP:
		println("UP")
	case DOWN:
		println("DOWN")
	case LEFT:
		moveLeft(values)
	case RIGHT:
		moveRight(values)
	}
}
