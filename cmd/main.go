package main

import (
	"github.com/joaovds/go-2048/internal/logic"
	"github.com/joaovds/go-2048/internal/ui"
)

func main() {
	initialValues := make([][]int, logic.SIZE)
	for i := range initialValues {
		initialValues[i] = make([]int, logic.SIZE)
	}

	logic.SetValueEmptyPosition(initialValues)
	logic.SetValueEmptyPosition(initialValues)

	layout := ui.NewLayout(initialValues)
	layout.Render()
}
