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
	t.setStyle()
	defer t.unsetStyle()

	MoveCursor(t.Row*TILE_HEIGHT+TILE_OFFSET+1, t.Col*TILE_WIDTH+5)
	fmt.Print(DoubleTopLeftCorner)
	for range 8 {
		fmt.Print(DoubleHorizontalLine)
	}
	fmt.Println(DoubleTopRightCorner)

	MoveCursor(t.Row*TILE_HEIGHT+TILE_OFFSET+2, t.Col*TILE_WIDTH+5)
	fmt.Printf("%s        %s\n", DoubleVerticalLine, DoubleVerticalLine)

	MoveCursor(t.Row*TILE_HEIGHT+TILE_OFFSET+3, t.Col*TILE_WIDTH+5)
	if t.Value == 0 {
		fmt.Printf("%s        %s\n", DoubleVerticalLine, DoubleVerticalLine)
	} else {
		fmt.Printf("%s%*d%*s\n", DoubleVerticalLine, t.spacesBefore, t.Value, t.spacesAfter, DoubleVerticalLine)
	}

	MoveCursor(t.Row*TILE_HEIGHT+TILE_OFFSET+4, t.Col*TILE_WIDTH+5)
	fmt.Printf("%s        %s\n", DoubleVerticalLine, DoubleVerticalLine)

	MoveCursor(t.Row*TILE_HEIGHT+TILE_OFFSET+5, t.Col*TILE_WIDTH+5)
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

func (t *Tile) setStyle() {
	switch t.Value {
	case 0:
		Colors.BrightBlack()
		break
	case 2:
		Colors.Blue()
		break
	case 4:
		Colors.BrightBlue()
		break
	case 8:
		Colors.Green()
		break
	case 16:
		Colors.Cyan()
		break
	case 32:
		Colors.Lime()
		break
	case 64:
		Colors.Orange()
		break
	case 128:
		Colors.Magenta()
		break
	case 256:
		Colors.Purple()
		break
	case 512:
		Colors.Red()
		break
	case 1024:
		Colors.Pink()
		break
	case 2048:
		Colors.WinColor()
		break
	}
}

func (t *Tile) unsetStyle() {
	Colors.Reset()
}
