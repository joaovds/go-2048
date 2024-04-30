package logic

import (
	"math/rand"
)

func NewValue() int {
	random := rand.Intn(10)
	if random <= 8 {
		return 2
	}
	return 4
}

func RandomEmptyPosition(values [][]int) (int, int) {
	for {
		row := rand.Intn(4)
		col := rand.Intn(4)
		if values[row][col] == 0 {
			return row, col
		}
	}
}

func SetValueEmptyPosition(values [][]int) {
	row, col := RandomEmptyPosition(values)
	values[row][col] = NewValue()
}
