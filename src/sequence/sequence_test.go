package sequence

import (
	"core"
	"os"
	"path/filepath"
	"testing"
)

func TestSequenceBuildDot(t *testing.T) {
	prefix := "TestSequenceBuildDot"

	seq := &Sequence{}
	config := NewSequenceConfig()

	seq.AddParticipant(&Participant{Name: "UE", Font: config.ParticipantFont})
	seq.AddParticipant(&Participant{Name: "ZXUN B200", Font: config.ParticipantFont, IsFocus: true})
	seq.AddParticipant(&Participant{Name: "I/S-CSCF", Font: config.ParticipantFont})
	seq.AddParticipant(&Participant{Name: "SCC AS", Font: config.ParticipantFont})

	seq.AddMessage(&Message{From: "UE", To: "ZXUN B200", Name: "INVITE", Font: config.MsgFont, Seq: 1})
	seq.AddMessage(&Message{From: "ZXUN B200", To: "I/S-CSCF", Name: "INVITE", Font: config.MsgFont, Seq: 2})
	seq.AddMessage(&Message{From: "I/S-CSCF", To: "SCC AS", Name: "INVITE", Font: config.MsgFont, Seq: 3})
	seq.AddMessage(&Message{From: "SCC AS", To: "I/S-CSCF", Name: "INVITE 180", Font: config.MsgFont, Seq: 4})
	seq.AddMessage(&Message{From: "I/S-CSCF", To: "ZXUN B200", Name: "INVITE 180", Font: config.MsgFont, Seq: 5})
	seq.AddMessage(&Message{From: "ZXUN B200", To: "UE", Name: "INVITE 180", Font: config.MsgFont, Seq: 6})
	seq.AddMessage(&Message{From: "SCC AS", To: "UE", Name: "test1", Font: config.MsgFont, Seq: 7})
	seq.AddMessage(&Message{From: "UE", To: "SCC AS", Name: "test2", Font: config.MsgFont, Seq: 8})

	seq.AddNote(&Note{OverParticipant: "ZXUN B200", Name: "A", Font: config.NoteFont})

	//fmt.Print("dot =\n", dot)

	dotStandard := os.Args[len(os.Args)-1] + "\\src\\testdata\\test_standard\\test1.gv"
	dotOutput := os.Args[len(os.Args)-1] + "\\src\\testdata\\test_output\\test1.gv"

	pngStandard := os.Args[len(os.Args)-1] + "\\src\\testdata\\test_standard\\test2.png"
	pngOutput := os.Args[len(os.Args)-1] + "\\src\\testdata\\test_output\\test2.png"

	seq.BuildAndGenDotPng(pngOutput, config)

	seq.BuildDotFile(dotOutput, config)
	if !core.FileEqual(dotStandard, dotOutput) {
		t.Errorf("%s: ouput file \"%s\" is not equal standard file \"%s\"", prefix, filepath.Base(dotOutput), filepath.Base(dotStandard))
	}

	if !core.FileEqual(pngStandard, pngOutput) {
		t.Errorf("%s: ouput file \"%s\" is not equal standard file \"%s\"", prefix, filepath.Base(pngOutput), filepath.Base(pngStandard))
	}
}
