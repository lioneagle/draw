package implwin

import (
	"core"
	"win"
)

type GeometricPen struct {
	hPen  win.HPEN
	style core.PenStyle
	brush core.Brush
	width int
}

func NewGeometricPen(style core.PenStyle, width int, brush BrushWin) (*GeometricPen, error) {
	if brush == nil {
		return nil, core.NewError("brush cannot be nil")
	}
	style.Type = core.PEN_TYPE_GEOMETRIC
	winPenStyle := getPenStyleWin(style)

	hPen := win.ExtCreatePen(winPenStyle, uint32(width), brush.logbrush(), 0, nil)
	if hPen == 0 {
		return nil, core.NewError("ExtCreatePen failed")
	}

	return &GeometricPen{
		hPen:  hPen,
		style: style,
		width: width,
		brush: brush,
	}, nil
}

func (this *GeometricPen) Dispose() {
	if this.hPen != 0 {
		win.DeleteObject(win.HGDIOBJ(this.hPen))

		this.hPen = 0
	}
}

func (this *GeometricPen) handle() win.HPEN {
	return this.hPen
}

func (this *GeometricPen) Style() core.PenStyle {
	return this.style
}

func (this *GeometricPen) Width() int {
	return this.width
}

func (this *GeometricPen) Brush() core.Brush {
	return this.brush
}
