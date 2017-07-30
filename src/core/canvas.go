package core

type Canvas interface {
	DrawLine(from, to Point, pen Pen) error
	DrawRectangle(rect Rectangle, pen Pen) error
	FillRectangle()
}
