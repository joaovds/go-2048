package ui


type Layout struct {
}

func NewLayout() *Layout {
	return &Layout{
	}
}

func (l *Layout) Render() {
	println("Rendering layout")
}
