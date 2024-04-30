package movements

func Move(direction Direction) {
	switch direction {
	case UP:
		// move up
		println("UP")
	case DOWN:
		// move down
		println("DOWN")
	case LEFT:
		// move left
		println("LEFT")
	case RIGHT:
		// move right
		println("RIGHT")
	}
}
