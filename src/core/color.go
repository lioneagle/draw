package core

import (
	"fmt"
)

type Color struct {
	Red   byte
	Green byte
	Blue  byte
}

func NewColor(Red, Green, Blue byte) *Color {
	return &Color{Red: Red, Green: Green, Blue: Blue}
}

func (this *Color) String() string {
	return fmt.Sprintf("#%02x%02x%02x", this.Red, this.Green, this.Blue)
}

type ColorPool struct {
	data map[string]*Color
}

func NewColorPool() *ColorPool {
	return &ColorPool{}
}

func (this *ColorPool) AddColor(name string, color *Color) {
	this.data[name] = color
}

func (this *ColorPool) GetColor(name string) (color *Color, ok bool) {
	color, ok = this.data[name]
	return color, ok
}

var (
	ColorBlack *Color = &Color{}
	ColorWhite *Color = NewColor(255, 255, 255)
	ColorRed   *Color = &Color{Red: 255}
	ColorGreen *Color = &Color{Green: 255}
	ColorBlue  *Color = &Color{Blue: 255}
	ColorAqua  *Color = &Color{Red: 0, Green: 255, Blue: 255}
	ColorGold  *Color = &Color{Red: 255, Green: 201, Blue: 14}
)
