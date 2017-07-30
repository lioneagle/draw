package implwin

import (
	"core"
	"win"
)

type nullPenWin struct {
	hPen win.HPEN
}

func newNullPen() *nullPenWin {
	lb := &win.LOGBRUSH{LbStyle: win.BS_NULL}

	hPen := win.ExtCreatePen(win.PS_COSMETIC|win.PS_NULL, 1, lb, 0, nil)
	if hPen == 0 {
		panic("failed to create null brush")
	}

	return &nullPenWin{hPen: hPen}
}

func (p *nullPenWin) Dispose() {
	if p.hPen != 0 {
		win.DeleteObject(win.HGDIOBJ(p.hPen))

		p.hPen = 0
	}
}

func (p *nullPenWin) handle() win.HPEN {
	return p.hPen
}

func (p *nullPenWin) Style() core.PenStyle {
	return core.PenStyle{}
}

func (p *nullPenWin) Width() int {
	return 0
}

var nullPenSingleton core.Pen = newNullPen()

func NullPen() core.Pen {
	return nullPenSingleton
}
