package implwin

import (
	"core"
	"win"
)

type CanvasWin struct {
	hdc          win.HDC
	hwnd         win.HWND
	dpix         int
	dpiy         int
	doNotDispose bool
}

func (this *CanvasWin) SetHdc(hdc win.HDC)    { this.hdc = hdc }
func (this *CanvasWin) SetHwnd(hwnd win.HWND) { this.hwnd = hwnd }

func NewCanvasWinFromHDC(hdc win.HDC) (core.Canvas, error) {
	if hdc == 0 {
		return nil, core.NewError("invalid hdc")
	}

	return (&CanvasWin{hdc: hdc, doNotDispose: true}).init()
}

func (this *CanvasWin) init() (core.Canvas, error) {
	this.dpix = int(win.GetDeviceCaps(this.hdc, win.LOGPIXELSX))
	this.dpiy = int(win.GetDeviceCaps(this.hdc, win.LOGPIXELSY))

	if win.SetBkMode(this.hdc, win.TRANSPARENT) == 0 {
		return nil, core.NewError("SetBkMode failed")
	}

	switch win.SetStretchBltMode(this.hdc, win.HALFTONE) {
	case 0, win.ERROR_INVALID_PARAMETER:
		return nil, core.NewError("SetStretchBltMode failed")
	}

	if !win.SetBrushOrgEx(this.hdc, 0, 0, nil) {
		return nil, core.NewError("SetBrushOrgEx failed")
	}

	return this, nil
}

func (this *CanvasWin) Dispose() {
	if !this.doNotDispose && this.hdc != 0 {
		if this.hwnd == 0 {
			win.DeleteDC(this.hdc)
		} else {
			win.ReleaseDC(this.hwnd, this.hdc)
		}

		this.hdc = 0
	}
}

func (this *CanvasWin) DrawLine(from, to core.Point, pen core.Pen) error {
	return nil
}

func (this *CanvasWin) DrawRectangle(rect core.Rectangle, pen core.Pen) error {
	return nil
}

func (this *CanvasWin) FillRectangle() {

}

func (this *CanvasWin) withGdiObj(handle win.HGDIOBJ, f func() error) error {
	oldHandle := win.SelectObject(this.hdc, handle)
	if oldHandle == 0 {
		return core.NewError("SelectObject failed")
	}
	defer win.SelectObject(this.hdc, oldHandle)

	return f()
}

func (this *CanvasWin) withPen(pen PenWin, f func() error) error {
	return this.withGdiObj(win.HGDIOBJ(pen.handle()), f)
}

func (this *CanvasWin) withBrush(brush BrushWin, f func() error) error {
	return this.withGdiObj(win.HGDIOBJ(brush.handle()), f)
}

func (this *CanvasWin) withBrushAndPen(brush BrushWin, pen PenWin, f func() error) error {
	return this.withBrush(brush, func() error {
		return this.withPen(pen, f)
	})
}
