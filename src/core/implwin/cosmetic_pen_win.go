package implwin

import (
	"core"
	"win"
)

type CosmeticPen struct {
	hPen  win.HPEN
	style core.PenStyle
	color core.Color
}

func NewCosmeticPen(style core.PenStyle, color core.Color) (*CosmeticPen, error) {
	lb := &win.LOGBRUSH{LbStyle: win.BS_SOLID, LbColor: win.COLORREF(color)}

	style.Type = core.PEN_TYPE_COSMETIC
	winPenStyle := getPenStyleWin(style)

	hPen := win.ExtCreatePen(winPenStyle, 1, lb, 0, nil)
	if hPen == 0 {
		return nil, core.NewError("ExtCreatePen failed")
	}

	return &CosmeticPen{hPen: hPen, style: style, color: color}, nil
}

func (this *CosmeticPen) Dispose() {
	if this.hPen != 0 {
		win.DeleteObject(win.HGDIOBJ(this.hPen))

		this.hPen = 0
	}
}

func (this *CosmeticPen) handle() win.HPEN {
	return this.hPen
}

func (this *CosmeticPen) Style() core.PenStyle {
	return this.style
}

func (this *CosmeticPen) Color() core.Color {
	return this.color
}

func (this *CosmeticPen) Width() int {
	return 1
}
