package main

import (
	"core"
	//"os"
	"sequence"
)

func main() {
	seq := &sequence.Sequence{}

	color := core.NewColor()
	objectFont := &core.Font{Name: "ObjectFont", Style: "FangSong", Color: *color, Size: 14}
	messageFont := &core.Font{Name: "MessageFont", Style: "FangSong", Color: *color, Size: 9}
	noteFont := &core.Font{Name: "NodeFont", Style: "FangSong", Color: *color, Size: 9}

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

	seq.BuildAndGenDotPng(pngOutput)
}
