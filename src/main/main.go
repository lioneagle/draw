package main

import (
	"core"
	"core/implwin"
	"fmt"
	//"os"
	//"parser"
	"sdl"
	"sequence"
	"syscall"
	"win"
)

func pressESCtoQuit() {
	fmt.Println("pressESCtoQuit() function begin ...")
	gameOver := false
	for !gameOver {
		var gameEvent sdl.SDL_Event
		for sdl.SDL_PollEvent(&gameEvent) {
			if gameEvent.SDL_EventType.Type == sdl.SDL_QUIT {
				gameOver = true
			}

			if gameEvent.SDL_EventType.Type == sdl.SDL_KEYUP {
				if gameEvent.SDL_KeyboardEvent().Keysym.Sym == sdl.SDLK_ESCAPE {
					gameOver = true
				}
			}
		}
		fmt.Printf(".")
	}
}

func TestSdl() {
	err := sdl.SDL_Init(sdl.SDL_INIT_EVERYTHING)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sdl.SDL_Quit()

	err = sdl.SDL_WasInit(sdl.SDL_INIT_EVERYTHING)
	if err != nil {
		fmt.Println(err)
	}

	err = sdl.SDL_InitSubSystem(sdl.SDL_INIT_VIDEO)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sdl.SDL_QuitSubSystem(sdl.SDL_INIT_VIDEO)

	//window, err := sdl.SDL_CreateWindow("test1", 50, 50, 640, 480, sdl.SDL_WINDOW_OPENGL)
	//window, err := sdl.SDL_CreateWindow("test1", 50, 50, 640, 480, 0)
	//window, err := sdl.SDL_CreateWindow("test1", 50, 50, 640, 480, sdl.SDL_WINDOW_RESIZABLE)
	//window, err := sdl.SDL_CreateWindow("test1", 50, 50, 640, 480, sdl.SDL_WINDOW_RESIZABLE|sdl.SDL_WINDOWPOS_CENTERED_MASK)
	window, err := sdl.SDL_CreateWindow("test1", sdl.SDL_WINDOWPOS_CENTERED, sdl.SDL_WINDOWPOS_CENTERED,
		640, 480, sdl.SDL_WINDOW_RESIZABLE|sdl.SDL_WINDOW_HIDDEN)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sdl.SDL_DestroyWindow(window)

	fmt.Println("Program is running, press ESC to quit.")
	pressESCtoQuit()
	fmt.Println("GAME OVER")

	return

}

func TestGdi() {
	fmt.Println("w1 =", syscall.StringToUTF16("123abc"))

	w1 := win.StringToWcharPtr("123abc")
	s1 := win.WcharPtrToString(w1)
	fmt.Println("s1 =", s1)

	var input win.GdiplusStartupInput
	var output win.GdiplusStartupOutput

	input.GdiplusVersion = 1

	err := win.GdiplusStartup(&input, &output)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer win.GdiplusShutdown()

	canvas, err := implwin.NewCanvasWin()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer canvas.Dispose()

	/*hdc := win.CreateCompatibleDC(0)
	defer win.DeleteDC(hdc)*/
	hdc := canvas.HDC()

	fmt.Printf("hdc = 0x%08x\n", hdc)

	//hbmp := win.CreateCompatibleBitmap(hdc, 600, 600)
	/*hbmp := win.CreateBitmap(600, 600, 1, 32, nil)
	defer win.DeleteObject(win.HGDIOBJ(hbmp))

	win.SelectObject(hdc, win.HGDIOBJ(hbmp))*/

	bmp, err := implwin.NewBitmap(core.Size{600, 600})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer bmp.Dispose()

	bmp.BeginPaint(canvas)

	//win.FillRect(hdc, &win.RECT{0, 0, 600, 600}, (win.HBRUSH)(win.GetStockObject(win.WHITE_BRUSH)))

	brush_0, err := implwin.NewSolidColorBrush(core.ColorWhite)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer brush_0.Dispose()

	canvas.FillRectangle(brush_0, core.Rectangle{0, 0, 600, 600})

	brush_1, err := implwin.NewSolidColorBrush(core.RGB(0, 255, 255))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer brush_1.Dispose()

	pen_1, err := implwin.NewGeometricPenWin(core.PenStyle{Style: core.PEN_SOLID}, 1, brush_1)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer pen_1.Dispose()

	canvas.DrawLine(pen_1, core.Point{300, 300}, core.Point{350, 350})
	canvas.DrawRectangle(pen_1, core.Rectangle{300, 300, 50, 50})
	canvas.DrawEllipse(pen_1, core.Rectangle{100, 100, 300, 400})
	canvas.DrawCircle(pen_1, core.Point{250, 250}, 150)

	//oldHandle := win.SelectObject(hdc, win.HGDIOBJ(pen_1.HPen))

	/*
		hpen_1 := win.CreatePen(win.PS_SOLID, 10, 0x0000FF00)
		defer win.DeleteObject(win.HGDIOBJ(hpen_1))

		oldHandle := win.SelectObject(hdc, win.HGDIOBJ(hpen_1))

		//oldHandle := win.SelectObject(hdc, win.HGDIOBJ(pen_1.HPen))
		win.MoveToEx(hdc, 300, 300, nil)
		win.LineTo(hdc, 350, 350)

		win.SelectObject(hdc, oldHandle)
		//*/

	/*

	 */

	/*
		var graphics *win.GpGraphics
		err = win.GdipCreateFromHDC(canvas.HDC(), &graphics)
		//err = win.GdipGetImageGraphicsContext(&bitmap.GpImage, &graphics)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer win.GdipDeleteGraphics(graphics)

		fmt.Printf("graphics = %p\n", graphics)

		var hdc2 win.HDC
		win.GdipGetDC(graphics, &hdc2)
		fmt.Printf("graphics.GetDC() = 0x%x\n", hdc2)

		//canvas.DrawRectangle(pen_1, core.Rectangle{100, 100, 50, 50})

		var brush1 *win.GpSolidFill
		err = win.GdipCreateSolidFill(0xFF0000FF, &brush1)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer win.GdipDeleteBrush(&brush1.GpBrush)

		var pen1 *win.GpPen
		err = win.GdipCreatePen1(0xFFFF0000, 1, 2, &pen1)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer win.GdipDeletePen(pen1)

		//err = win.GdipSetSmoothingMode(graphics, win.SmoothingModeAntiAlias)
		err = win.GdipSetSmoothingMode(graphics, win.SmoothingModeDefault)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = win.GdipDrawRectangle(graphics, pen1, 50, 50, 200, 200)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = win.GdipFillRectangle(graphics, &brush1.GpBrush, 50, 50, 100, 100)
		if err != nil {
			fmt.Println(err)
			return
		}

		var brush *win.GpSolidFill
		err = win.GdipCreateSolidFill(0x900000FF, &brush)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer win.GdipDeleteBrush(&brush.GpBrush)
		var lineBrush *win.GpLineGradient
		p1 := &win.GpPoint{0, 0}
		p2 := &win.GpPoint{280, 20}
		err = win.GdipCreateLineBrush(p1, p2, 0xFFFF0000, 0xB00000FF, win.WrapModeTile, &lineBrush)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer win.GdipDeleteBrush(&lineBrush.GpBrush)
		err = win.GdipSetTextRenderingHint(graphics, win.TextRenderingHintAntiAliasGridFit)
		if err != nil {
			fmt.Println(err)
			return
		}

		var fontFamily *win.GpFontFamily
		err = win.GdipCreateFontFamilyFromName("宋体", nil, &fontFamily)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer win.GdipDeleteFontFamily(fontFamily)

		var format *win.GpStringFormat
		err = win.GdipCreateStringFormat(0, win.LANG_NEUTRAL, &format)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer win.GdipDeleteStringFormat(format)

		err = win.GdipSetStringFormatAlign(format, win.StringAlignmentNear)
		if err != nil {
			fmt.Println(err)
			return
		}

		var font *win.GpFont
		err = win.GdipCreateFont(fontFamily, 14, 0, win.UnitPoint, &font)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer win.GdipDeleteFont(font)

		rect := win.RectF{20, 20, 280, 20}
		err = win.GdipDrawString(graphics, "测试：I love Win32 and GdiplusFlat", -1, font, &rect, nil, &brush.GpBrush)
		if err != nil {
			fmt.Println(err)
			return
		}

		rect = win.RectF{20, 60, 280, 20}
		err = win.GdipDrawString(graphics, "测试：I love Win32 and GdiplusFlat", -1, font, &rect, nil, &lineBrush.GpBrush)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = win.GdipDrawString(graphics, "测试：I love Win32 and GdiplusFlat", -1, font, &rect, nil, &lineBrush.GpBrush)
		if err != nil {
			fmt.Println(err)
			return
		} //*/

	bmp.EndPaint()

	bmp.SaveToFile("test1.bmp", "bmp")

	return
}

func CreateConfig() *sequence.SequenceConfig {
	config := sequence.NewSequenceConfig()
	config.PlantumlJarPath = "F:\\DevCode\\go_code\\src\\my_code\\draw\\"
	return config
}

func CreateSequence(config *sequence.SequenceConfig) *sequence.Sequence {
	seq := &sequence.Sequence{}

	participantFont := config.ParticipantFont
	messageFont := config.MsgFont
	noteFont := config.NoteFont

	seq.AddParticipant(&sequence.Participant{Name: "ue", Label: "UE", Font: participantFont})
	seq.AddParticipant(&sequence.Participant{Name: "sbc", Label: "ZXUN B200", Font: participantFont, IsFocus: true})
	seq.AddParticipant(&sequence.Participant{Name: "cscf", Label: "I/S-CSCF", Font: participantFont})
	seq.AddParticipant(&sequence.Participant{Name: "scc_as", Label: "SCC AS", Font: participantFont})

	seq.AddMessage(&sequence.Message{From: "ue", To: "sbc", Name: "INVITE", Font: messageFont, Seq: 1})
	seq.AddMessage(&sequence.Message{From: "sbc", To: "cscf", Name: "INVITE", Font: messageFont, Seq: 2})
	seq.AddMessage(&sequence.Message{From: "cscf", To: "scc_as", Name: "INVITE", Font: messageFont, Seq: 2})
	seq.AddMessage(&sequence.Message{From: "scc_as", To: "cscf", Name: "INVITE 180", Font: messageFont, Seq: 3})
	seq.AddMessage(&sequence.Message{From: "cscf", To: "sbc", Name: "INVITE 180", Font: messageFont, Seq: 3})
	seq.AddMessage(&sequence.Message{From: "sbc", To: "ue", Name: "INVITE 180", Font: messageFont, Seq: 4})
	seq.AddMessage(&sequence.Message{From: "scc_as", To: "ue", Name: "test1", Font: messageFont, Seq: 4})
	seq.AddMessage(&sequence.Message{From: "ue", To: "scc_as", Name: "test2", Font: messageFont, Seq: 4})

	seq.AddNote(&sequence.Note{OverParticipant: "sbc", Name: "A", Font: noteFont})

	return seq
}

func GenDot() {
	config := CreateConfig()
	seq := CreateSequence(config)

	//fmt.Print("dot =\n", dot)

	//dotStandard := os.Args[len(os.Args)-1] + "\\src\\testdata\\test_standard\\test1.gv"
	//dotOutput := os.Args[len(os.Args)-1] + "\\src\\testdata\\test_output\\test1.gv"

	//pngStandard := os.Args[len(os.Args)-1] + "\\src\\testdata\\test_standard\\test2.png"
	pngOutput := `F:\DevCode\go_code\src\my_code\draw\src\testdata\test_output\test2.png`
	//pngOutput := `d:/test2.png`

	seq.BuildAndGenDotPng(pngOutput, config)
}

func GenPlantuml() {
	config := CreateConfig()
	seq := CreateSequence(config)

	//fmt.Print("dot =\n", dot)

	//dotStandard := os.Args[len(os.Args)-1] + "\\src\\testdata\\test_standard\\test1.gv"
	//dotOutput := os.Args[len(os.Args)-1] + "\\src\\testdata\\test_output\\test1.gv"

	//pngStandard := os.Args[len(os.Args)-1] + "\\src\\testdata\\test_standard\\test2.png"
	pngOutput := `F:\DevCode\go_code\src\my_code\draw\src\testdata\test_output\test3.png`
	//pngOutput := `d:/test2.png`

	seq.BuildAndGenPlantumlPng(pngOutput, config)
}

func main() {

	fmt.Printf("parser.TOKEN_CONFIG")
	return
	GenPlantuml()
}
