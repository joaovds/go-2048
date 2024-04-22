package ui

import "fmt"

type Tile struct {
	Value        int
	Row          int
	Col          int
	spacesBefore int
	spacesAfter  int
}

func NewTile(value, row, col int) Tile {
	tile := Tile{
		Value:        value,
		Row:          row,
		Col:          col,
		spacesBefore: 0,
		spacesAfter:  0,
	}

	return tile
}

func (t *Tile) Render() {
	t.CalculateSpaces()

	MoveCursor(t.Row*TILE_HEIGHT+TILE_OFFSET+1, t.Col*TILE_WIDTH+1)
	fmt.Print(DoubleTopLeftCorner)
	for range 8 {
		fmt.Print(DoubleHorizontalLine)
	}
	fmt.Println(DoubleTopRightCorner)

	MoveCursor(t.Row*TILE_HEIGHT+TILE_OFFSET+2, t.Col*TILE_WIDTH+1)
	fmt.Printf("%s        %s\n", DoubleVerticalLine, DoubleVerticalLine)

	MoveCursor(t.Row*TILE_HEIGHT+TILE_OFFSET+3, t.Col*TILE_WIDTH+1)
	if t.Value == 0 {
		fmt.Printf("%s        %s\n", DoubleVerticalLine, DoubleVerticalLine)
	} else {
		fmt.Printf("%s%*d%*s\n", DoubleVerticalLine, t.spacesBefore, t.Value, t.spacesAfter, DoubleVerticalLine)
	}

	MoveCursor(t.Row*TILE_HEIGHT+TILE_OFFSET+4, t.Col*TILE_WIDTH+1)
	fmt.Printf("%s        %s\n", DoubleVerticalLine, DoubleVerticalLine)

	MoveCursor(t.Row*TILE_HEIGHT+TILE_OFFSET+5, t.Col*TILE_WIDTH+1)
	fmt.Print(DoubleBottomLeftCorner)
	for range 8 {
		fmt.Print(DoubleHorizontalLine)
	}
	fmt.Println(DoubleBottomRightCorner)
}

func (t *Tile) CalculateSpaces() {
	if t.Value <= 100 {
		t.spacesBefore = 5
		t.spacesAfter = 4
	} else if t.Value <= 10000 {
		t.spacesBefore = 6
		t.spacesAfter = 3
	}
}
