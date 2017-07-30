package core

type Font struct {
	Name            string
	FrontColor      Color
	BackgroundColor Color
	Style           int
	Size            int
}

const (
	FONT_STYLE_NORMAL = iota
	FONT_STYLE_BOLD
)

func NewFont(Name string, Style, Size int) *Font {
	f := &Font{Name: Name, Style: Style, FrontColor: ColorBlack, BackgroundColor: ColorWhite, Size: Size}
	return f
}

func (this *Font) GetDotName() string {
	s := this.Name
	if this.Style == FONT_STYLE_BOLD {
		s += " bold"
	}
	return s
}
