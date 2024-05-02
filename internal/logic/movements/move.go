package movements

func Move(direction Direction, values [][]int, updateSignal chan<- struct{}) {
	switch direction {
	case UP:
		println("UP")
		updateSignal <- struct{}{}
	case DOWN:
		println("DOWN")
		updateSignal <- struct{}{}
	case LEFT:
		println("LEFT")
		updateSignal <- struct{}{}
	case RIGHT:
		moveRight(values)
		updateSignal <- struct{}{}
	}
}
