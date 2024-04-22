package ui

type Board struct {
	tiles [][]Tile
}

func NewBoard() *Board {
	tiles := make([][]Tile, SIZE)

	for r := range tiles {
		tiles[r] = make([]Tile, SIZE)

		for c := range tiles[r] {
			tiles[r][c] = NewTile(0, r, c)
		}
	}

	return &Board{
		tiles: tiles,
	}
}

func (b *Board) Render(values [][]int) {
	for r, row := range b.tiles {
		for c, tile := range row {
			tile.Value = values[r][c]
			tile.Value = 32
			tile.Render()
		}
	}
}
