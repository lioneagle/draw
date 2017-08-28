package sequence

import (
	"core"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestSequenceBuildDot(t *testing.T) {
	prefix := "TestSequenceBuildDot"

	seq := &Sequence{}
	config := NewSequenceConfig()

	seq.AddParticipant(&Participant{Name: "ue", Label: "UE", Font: config.ParticipantFont})
	seq.AddParticipant(&Participant{Name: "sbc", Label: "ZXUN B200", Font: config.ParticipantFont, IsFocus: true})
	seq.AddParticipant(&Participant{Name: "cscf", Label: "I/S-CSCF", Font: config.ParticipantFont})
	seq.AddParticipant(&Participant{Name: "scc_as", Label: "SCC AS", Font: config.ParticipantFont})

	seq.AddMessage(&Message{From: "ue", To: "sbc", Name: "INVITE", Font: config.MsgFont, Seq: 1})
	seq.AddMessage(&Message{From: "sbc", To: "cscf", Name: "INVITE", Font: config.MsgFont, Seq: 2})
	seq.AddMessage(&Message{From: "cscf", To: "scc_as", Name: "INVITE", Font: config.MsgFont, Seq: 3})
	seq.AddMessage(&Message{From: "scc_as", To: "cscf", Name: "INVITE 180", Font: config.MsgFont, Seq: 4})
	seq.AddMessage(&Message{From: "cscf", To: "sbc", Name: "INVITE 180", Font: config.MsgFont, Seq: 5})
	seq.AddMessage(&Message{From: "sbc", To: "ue", Name: "INVITE 180", Font: config.MsgFont, Seq: 6})
	seq.AddMessage(&Message{From: "scc_as", To: "ue", Name: "test1", Font: config.MsgFont, Seq: 7})
	seq.AddMessage(&Message{From: "ue", To: "scc_as", Name: "test2", Font: config.MsgFont, Seq: 8})

	seq.AddNote(&Note{OverParticipant: "sbc", Name: "A", Font: config.NoteFont})

	//fmt.Print("dot =\n", dot)

	dotStandard := os.Args[len(os.Args)-1] + "\\src\\testdata\\test_standard\\test1.gv"
	dotOutput := os.Args[len(os.Args)-1] + "\\src\\testdata\\test_output\\test1.gv"

	pngStandard := os.Args[len(os.Args)-1] + "\\src\\testdata\\test_standard\\test1_dot.png"
	pngOutput := os.Args[len(os.Args)-1] + "\\src\\testdata\\test_output\\test1_dot.png"

	seq.BuildAndGenDotPng(pngOutput, config)

	seq.BuildDotFile(dotOutput, config)
	if !core.FileEqual(dotStandard, dotOutput) {
		t.Errorf("%s: ouput file \"%s\" is not equal standard file \"%s\"", prefix, filepath.Base(dotOutput), filepath.Base(dotStandard))
	}

	if !core.FileEqual(pngStandard, pngOutput) {
		t.Errorf("%s: ouput file \"%s\" is not equal standard file \"%s\"", prefix, filepath.Base(pngOutput), filepath.Base(pngStandard))
	}
}

func TestSequenceBuildplantuml(t *testing.T) {
	prefix := "TestSequenceBuildDot"

	fmt.Println("os.Args[0] =", os.Args[0])

	seq := &Sequence{}
	config := NewSequenceConfig()
	config.PlantumlJarPath = "F:\\DevCode\\go_code\\src\\my_code\\draw\\"

	seq.AddParticipant(&Participant{Name: "ue", Label: "UE", Font: config.ParticipantFont})
	seq.AddParticipant(&Participant{Name: "sbc", Label: "ZXUN B200", Font: config.ParticipantFont, IsFocus: true})
	seq.AddParticipant(&Participant{Name: "cscf", Label: "I/S-CSCF", Font: config.ParticipantFont})
	seq.AddParticipant(&Participant{Name: "scc_as", Label: "SCC AS", Font: config.ParticipantFont})

	seq.AddMessage(&Message{From: "ue", To: "sbc", Name: "INVITE", Font: config.MsgFont, Seq: 1})
	seq.AddMessage(&Message{From: "sbc", To: "cscf", Name: "INVITE", Font: config.MsgFont, Seq: 2})
	seq.AddMessage(&Message{From: "cscf", To: "scc_as", Name: "INVITE", Font: config.MsgFont, Seq: 3})
	seq.AddMessage(&Message{From: "scc_as", To: "cscf", Name: "INVITE 180", Font: config.MsgFont, Seq: 4})
	seq.AddMessage(&Message{From: "cscf", To: "sbc", Name: "INVITE 180", Font: config.MsgFont, Seq: 5})
	seq.AddMessage(&Message{From: "sbc", To: "ue", Name: "INVITE 180", Font: config.MsgFont, Seq: 6})
	seq.AddMessage(&Message{From: "scc_as", To: "ue", Name: "test1", Font: config.MsgFont, Seq: 7})
	seq.AddMessage(&Message{From: "ue", To: "scc_as", Name: "test2", Font: config.MsgFont, Seq: 8})

	seq.AddNote(&Note{OverParticipant: "sbc", Name: "A", Font: config.NoteFont})

	plantumlStandard := os.Args[len(os.Args)-1] + "\\src\\testdata\\test_standard\\test1_puml.puml"
	plantumlOutput := os.Args[len(os.Args)-1] + "\\src\\testdata\\test_output\\test1_puml.puml"

	pngStandard := os.Args[len(os.Args)-1] + "\\src\\testdata\\test_standard\\test1_puml.png"
	pngOutput := os.Args[len(os.Args)-1] + "\\src\\testdata\\test_output\\test1_puml.png"

	seq.BuildAndGenPlantumlPng(pngOutput, config)

	seq.BuildPlantumlFile(plantumlOutput, config)
	if !core.FileEqual(plantumlStandard, plantumlOutput) {
		t.Errorf("%s: ouput file \"%s\" is not equal standard file \"%s\"", prefix, filepath.Base(plantumlOutput), filepath.Base(plantumlStandard))
	}

	if !core.FileEqual(pngStandard, pngOutput) {
		t.Errorf("%s: ouput file \"%s\" is not equal standard file \"%s\"", prefix, filepath.Base(pngOutput), filepath.Base(pngStandard))
	}
}
