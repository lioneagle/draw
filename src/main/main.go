package main

import (
	//"core"
	//"os"
	"fmt"
	"sequence"
	"win"
)

func main() {

	var input win.GdiplusStartupInput
	var output win.GdiplusStartupOutput

	input.GdiplusVersion = 1

	err := win.GdiplusStartup(&input, &output)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer win.GdiplusShutdown()

	hdc := win.CreateCompatibleDC(0)
	defer win.DeleteDC(hdc)

	var graphics *win.GpGraphics
	err = win.GdipCreateFromHDC(hdc, &graphics)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer win.GdipDeleteGraphics(graphics)

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

	err = win.GdipSetSmoothingMode(graphics, win.SmoothingModeAntiAlias)
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
	}

	return

	seq := &sequence.Sequence{}
	config := sequence.NewSequenceConfig()

	objectFont := config.ObjectFont
	messageFont := config.MsgFont
	noteFont := config.NoteFont

	seq.AddObject(&sequence.Object{Name: "UE", Font: objectFont})
	seq.AddObject(&sequence.Object{Name: "ZXUN B200", Font: objectFont, IsMain: true})
	seq.AddObject(&sequence.Object{Name: "I/S-CSCF", Font: objectFont})
	seq.AddObject(&sequence.Object{Name: "SCC AS", Font: objectFont})

	seq.AddMessage(&sequence.Message{FromObj: "UE", ToObj: "ZXUN B200", Name: "INVITE", Font: messageFont, Seq: 1})
	seq.AddMessage(&sequence.Message{FromObj: "ZXUN B200", ToObj: "I/S-CSCF", Name: "INVITE", Font: messageFont, Seq: 2})
	seq.AddMessage(&sequence.Message{FromObj: "I/S-CSCF", ToObj: "SCC AS", Name: "INVITE", Font: messageFont, Seq: 2})
	seq.AddMessage(&sequence.Message{FromObj: "SCC AS", ToObj: "I/S-CSCF", Name: "INVITE 180", Font: messageFont, Seq: 3})
	seq.AddMessage(&sequence.Message{FromObj: "I/S-CSCF", ToObj: "ZXUN B200", Name: "INVITE 180", Font: messageFont, Seq: 3})
	seq.AddMessage(&sequence.Message{FromObj: "ZXUN B200", ToObj: "UE", Name: "INVITE 180", Font: messageFont, Seq: 4})
	seq.AddMessage(&sequence.Message{FromObj: "SCC AS", ToObj: "UE", Name: "test1", Font: messageFont, Seq: 4})
	seq.AddMessage(&sequence.Message{FromObj: "UE", ToObj: "SCC AS", Name: "test2", Font: messageFont, Seq: 4})

	seq.AddNote(&sequence.Note{Obj: "ZXUN B200", Name: "A", Font: noteFont})

	//fmt.Print("dot =\n", dot)

	//dotStandard := os.Args[len(os.Args)-1] + "\\src\\testdata\\test_standard\\test1.gv"
	//dotOutput := os.Args[len(os.Args)-1] + "\\src\\testdata\\test_output\\test1.gv"

	//pngStandard := os.Args[len(os.Args)-1] + "\\src\\testdata\\test_standard\\test2.png"
	pngOutput := `F:\DevCode\go_code\src\my_code\draw\src\testdata\test_output\test2.png`
	//pngOutput := `d:/test2.png`

	seq.BuildAndGenDotPng(pngOutput, config)

	var x win.ULONG

	x = 1

	fmt.Println("x =", x)
}
