package movements

func Move(direction Direction, values [][]int, updateSignal chan<- struct{}) {
	switch direction {
	case UP:
		// move up
		println("UP")
		updateSignal <- struct{}{}
	case DOWN:
		// move down
		println("DOWN")
		updateSignal <- struct{}{}
	case LEFT:
		// move left
		println("LEFT")
		updateSignal <- struct{}{}
	case RIGHT:
		// move right
		println("RIGHT")
		updateSignal <- struct{}{}
	}
}
