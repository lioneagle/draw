package implwin

import (
	"core"
	"win"
)

type CosmeticPenWin struct {
	hPen  win.HPEN
	style core.PenStyle
	color core.Color
}

func NewCosmeticPenWin(style core.PenStyle, color core.Color) (*CosmeticPenWin, error) {
	lb := &win.LOGBRUSH{LbStyle: win.BS_SOLID, LbColor: win.COLORREF(color)}

	style.Type = core.PEN_TYPE_COSMETIC
	winPenStyle := getPenStyleWin(style)

	hPen := win.ExtCreatePen(winPenStyle, 1, lb, 0, nil)
	if hPen == 0 {
		return nil, core.NewError("ExtCreatePen failed")
	}

	return &CosmeticPenWin{hPen: hPen, style: style, color: color}, nil
}

func (this *CosmeticPenWin) Dispose() {
	if this.hPen != 0 {
		win.DeleteObject(win.HGDIOBJ(this.hPen))

		this.hPen = 0
	}
}

func (this *CosmeticPenWin) handle() win.HPEN {
	return this.hPen
}

func (this *CosmeticPenWin) Style() core.PenStyle {
	return this.style
}

func (this *CosmeticPenWin) Color() core.Color {
	return this.color
}

func (this *CosmeticPenWin) Width() int {
	return 1
}
