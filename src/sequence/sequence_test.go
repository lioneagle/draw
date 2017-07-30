package sequence

import (
	"core"
	"os"
	"path/filepath"
	"testing"
)

func TestSequenceBuildDot(t *testing.T) {
	prefix := "TestSequenceBuildDot
	
	seq := &Sequence{}

	config := NewSequenceConfig()

	seq.AddObject(&Object{Name: "UE", Font: config.ObjectFont})
	seq.AddObject(&Object{Name: "ZXUN B200", Font: config.ObjectFont, IsMain: true})
	seq.AddObject(&Object{Name: "I/S-CSCF", Font: config.ObjectFont})
	seq.AddObject(&Object{Name: "SCC AS", Font: config.ObjectFont})

	seq.AddMessage(&Message{FromObj: "UE", ToObj: "ZXUN B200", Name: "INVITE", Font: config.MsgFont, Seq: 1})
	seq.AddMessage(&Message{FromObj: "ZXUN B200", ToObj: "I/S-CSCF", Name: "INVITE", Font: config.MsgFont, Seq: 2})
	seq.AddMessage(&Message{FromObj: "I/S-CSCF", ToObj: "SCC AS", Name: "INVITE", Font: config.MsgFont, Seq: 3})
	seq.AddMessage(&Message{FromObj: "SCC AS", ToObj: "I/S-CSCF", Name: "INVITE 180", Font: config.MsgFont, Seq: 4})
	seq.AddMessage(&Message{FromObj: "I/S-CSCF", ToObj: "ZXUN B200", Name: "INVITE 180", Font: config.MsgFont, Seq: 5})
	seq.AddMessage(&Message{FromObj: "ZXUN B200", ToObj: "UE", Name: "INVITE 180", Font: config.MsgFont, Seq: 6})
	seq.AddMessage(&Message{FromObj: "SCC AS", ToObj: "UE", Name: "test1", Font: config.MsgFont, Seq: 7})
	seq.AddMessage(&Message{FromObj: "UE", ToObj: "SCC AS", Name: "test2", Font: config.MsgFont, Seq: 8})

	seq.AddNote(&Note{Obj: "ZXUN B200", Name: "A", Font: config.NoteFont})

	//fmt.Print("dot =\n", dot)

	dotStandard := os.Args[len(os.Args)-1] + "\\src\\testdata\\test_standard\\test1.gv"
	dotOutput := os.Args[len(os.Args)-1] + "\\src\\testdata\\test_output\\test1.gv"

	pngStandard := os.Args[len(os.Args)-1] + "\\src\\testdata\\test_standard\\test2.png"
	pngOutput := os.Args[len(os.Args)-1] + "\\src\\testdata\\test_output\\test2.png"

	seq.BuildAndGenDotPng(pngOutput, config)

	seq.BuildDotFile(dotOutput, config)
	if !core.FileEqual(dotStandard, dotOutput) {
		t.Errorf("%s: ouput file \"%s\" is not equal standard file \"%s\"",prefix, filepath.Base(dotOutput), filepath.Base(dotStandard))
	}

	if !core.FileEqual(pngStandard, pngOutput) {
		t.Errorf("%s: ouput file \"%s\" is not equal standard file \"%s\"",prefix, filepath.Base(pngOutput), filepath.Base(pngStandard))
	}
}
