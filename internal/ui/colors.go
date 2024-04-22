package ui

import "fmt"

type colors struct{}

var Colors = &colors{}

func SetTextColor(color string) {
	fmt.Printf("\033[%sm", color)
}

func SetTextColor256(color string) {
	fmt.Printf("\033[38;5;%sm", color)
}

func (c *colors) Reset() {
	fmt.Print("\033[0m")
}

func (c *colors) Black() {
	SetTextColor("30")
}

func (c *colors) BrightBlack() {
	SetTextColor("90")
}

func (c *colors) Red() {
	SetTextColor("31")
}

func (c *colors) Green() {
	SetTextColor("32")
}

func (c *colors) BrightBlue() {
	SetTextColor("94")
}

func (c *colors) Magenta() {
	SetTextColor("35")
}

func (c *colors) Cyan() {
	SetTextColor("36")
}

func (c *colors) Yellow() {
	SetTextColor("33")
}

func (c *colors) Orange() {
	SetTextColor256("166")
}

func (c *colors) Purple() {
	SetTextColor256("129")
}
