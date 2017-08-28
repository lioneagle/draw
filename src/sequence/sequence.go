package sequence

import (
	"core"
	"errors"
	"fmt"
	"os"
	"os/exec"
	//"path/filepath"
)

const (
	STEP_TYPE_NORMAL = iota
	STEP_TYPE_NOTE
)

type Step struct {
	Id   int
	Type int
	Name string
	Font *core.Font
}

type Steps struct {
	Data []*Step
}

type Participant struct {
	Name    string
	Label   string
	Id      int
	Font    *core.Font
	IsFocus bool
}

const (
	ACTION_TYPE_MESSAGE = iota
	ACTION_TYPE_NOTE
)

type Action interface {
	Type() int
	GetRow() int
}

type Message struct {
	From string
	To   string
	Name string
	Font *core.Font
	Row  int
	Seq  int
	Id   int
}

func (this *Message) Type() int   { return ACTION_TYPE_MESSAGE }
func (this *Message) GetRow() int { return this.Row }

type Note struct {
	OverParticipant string
	Name            string
	Font            *core.Font
	Id              int
	Row             int
}

func (this *Note) Type() int   { return ACTION_TYPE_NOTE }
func (this *Note) GetRow() int { return this.Row }

type SequenceConfig struct {
	TextBackgroundColor       core.Color
	CrossNeighborMsgColor     core.Color
	FocusParticipantFillColor core.Color
	ParticipantFont           *core.Font
	MsgFont                   *core.Font
	NoteFont                  *core.Font
	PlantumlJarPath           string
}

func NewSequenceConfig() *SequenceConfig {
	s := &SequenceConfig{}

	s.TextBackgroundColor = core.RGB(232, 248, 247)
	s.CrossNeighborMsgColor = core.RGB(68, 170, 205)
	s.FocusParticipantFillColor = core.ColorGold

	s.ParticipantFont = core.NewFont("Arial", core.FONT_STYLE_BOLD, 9)
	s.MsgFont = core.NewFont("Arial", core.FONT_STYLE_BOLD, 9)
	s.NoteFont = core.NewFont("Arial", core.FONT_STYLE_BOLD, 9)

	return s
}

type Sequence struct {
	participants []*Participant
	actions      []Action
}

func (this *Sequence) AddParticipant(obj *Participant) {
	v := this.findParticipant(obj.Name)
	if v == nil {
		this.participants = append(this.participants, obj)
	}
}

func (this *Sequence) findParticipant(name string) *Participant {
	for _, v := range this.participants {
		if v.Name == name {
			return v
		}
	}
	return nil
}

func (this *Sequence) AddMessage(message *Message) {
	this.actions = append(this.actions, message)
}

func (this *Sequence) AddNote(note *Note) {
	this.actions = append(this.actions, note)
}

func (this *Sequence) BuildAndGenDotPng(pngfile string, config *SequenceConfig) {
	dotfile := core.ReplaceFileSuffix(pngfile, "gv")
	this.BuildDotFile(dotfile, config)
	this.GenDotPng(dotfile, pngfile)
}

func (this *Sequence) GenDotPng(dotfile, pngfile string) {
	cmd := exec.Command("dot", "-Kdot", "-Tpng", dotfile, "-o", pngfile)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("GenDotPng failed: ", err.Error())
	}
	os.Remove(dotfile)
}

func (this *Sequence) BuildDotFile(filename string, config *SequenceConfig) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("ERROR: cannot open file %s\r\n", filename)
		return
	}
	defer file.Close()

	file.WriteString(this.BuildDot(config))
}

func (this *Sequence) BuildDot(config *SequenceConfig) string {
	var buf core.Writer

	m, err := this.getSteps()
	if err != nil {
		fmt.Printf("BuildDot failed: %s\n", err.Error())
		return ""
	}

	format := `digraph G {
    rankdir="LR";
    node[shape="point", width=0, height=0, fontname="%s", fontsize=%d];
    edge[arrowhead="none", style="solid", fontname="%s", fontsize=%d];
	
	`

	buf.Write(format,
		config.ParticipantFont.GetDotName(),
		config.ParticipantFont.Size,
		config.MsgFont.GetDotName(),
		config.MsgFont.Size)

	/*buf.WriteString(fmt.Sprintf(format,
	config.ObjectFont.GetDotName(), config.ObjectFont.Size,
	config.MsgFont.GetDotName(), config.MsgFont.Size)
	)*/

	for i, v := range this.participants {
		v.Id = i
	}

	for _, v := range this.participants {
		buf.Writeln("")
		buf.Writeln("    {")
		buf.Writeln("        rank=\"same\";")
		buf.Writeln("        edge[style=\"solid\"];")
		if !v.IsFocus {
			buf.Writeln("        %s[shape=\"box\", label=\"%s\", width=1, height=0.5];", v.Name, v.Label)
		} else {
			buf.Writeln("        %s[shape=\"box\", label=\"%s\", fillcolor=\"%s\", style=filled, width=1, height=0.5];",
				v.Name, v.Label, config.FocusParticipantFillColor.RGBString())
		}

		steps, ok := m[v.Name]
		if !ok {
			fmt.Printf("ERROR: cannot find object \"%s\"\n", v.Name)
			return ""
		}
		for j, s := range steps.Data {
			if j == (len(steps.Data) - 1) {
				buf.Writeln("        %s_step_%d[shape=\"box\", width=0.5, label=\"\"];", v.Name, s.Id)
				break
			}
			if s.Type == STEP_TYPE_NOTE {
				buf.Writeln("        %s_note_%d[shape=\"circle\", label=\"%s\", width=0.51];", v.Name, s.Id, s.Name)
			}
		}

		buf.Write("        %s", v.Name)

		for _, s := range steps.Data {
			if s.Type == STEP_TYPE_NOTE {
				buf.Write(" -> %s_note_%d", v.Name, s.Id)
			} else {
				buf.Write(" -> %s_step_%d", v.Name, s.Id)
			}
		}

		buf.Write(";\r\n    }\r\n")
	}

	buf.Write("\r\n")

	for _, a := range this.actions {
		switch data := a.(type) {
		case *Message:
			from := this.findParticipant(data.From)
			if from == nil {
				fmt.Printf("ERROR: cannot find object \"%s\"\n", data.From)
				return ""
			}
			to := this.findParticipant(data.To)
			if to == nil {
				fmt.Printf("ERROR: cannot find object \"%s\"\n", data.To)
				return ""
			}

			var k int

			if to.Id > from.Id {
				if to.Id == (from.Id + 1) {
					buf.Writeln("    %s_step_%d -> %s_step_%d [label=\"%s\", arrowhead=\"normal\"];",
						from.Name, data.Id, to.Name, data.Id, data.Name)
				} else {
					buf.Writeln("    %s_step_%d -> %s_step_%d [label=\"%s\", color=\"%s\"];",
						from.Name, data.Id, this.participants[from.Id+1].Name, data.Id, data.Name, config.CrossNeighborMsgColor.RGBString())
					for k = from.Id + 1; k < to.Id; k++ {
						if k != (to.Id - 1) {
							buf.Writeln("    %s_step_%d -> %s_step_%d [color=\"%s\"];",
								this.participants[k].Name, data.Id, this.participants[k+1].Name, data.Id, config.CrossNeighborMsgColor.RGBString())
						} else {
							buf.Writeln("    %s_step_%d -> %s_step_%d [arrowhead=\"normal\", color=\"%s\"];",
								this.participants[k].Name, data.Id, this.participants[k+1].Name, data.Id, config.CrossNeighborMsgColor.RGBString())
						}
					}
				}
			} else {
				if to.Id == (from.Id - 1) {
					buf.Writeln("    %s_step_%d -> %s_step_%d [label=\"%s\", arrowhead=\"normal\"];",
						from.Name, data.Id, to.Name, data.Id, data.Name)
				} else {
					buf.Writeln("    %s_step_%d -> %s_step_%d [label=\"%s\", color=\"%s\", ];",
						from.Name, data.Id, this.participants[from.Id-1].Name, data.Id, data.Name, config.CrossNeighborMsgColor.RGBString())
					for k = from.Id - 1; k > to.Id; k-- {
						if k != (to.Id + 1) {
							buf.Writeln("    %s_step_%d -> %s_step_%d [color=\"%s\"];",
								this.participants[k].Name, data.Id, this.participants[k-1].Name, data.Id, config.CrossNeighborMsgColor.RGBString())
						} else {
							buf.Writeln("    %s_step_%d -> %s_step_%d [arrowhead=\"normal\", color=\"%s\"];",
								this.participants[k].Name, data.Id, this.participants[k-1].Name, data.Id, config.CrossNeighborMsgColor.RGBString())
						}
					}
				}
			}
		}
	}

	buf.Writeln("}")

	return buf.String()
}

func (this *Sequence) getSteps() (map[string]*Steps, error) {
	m := make(map[string]*Steps)
	totalSteps := this.getTotalStepNum()

	for _, v := range this.participants {
		m[v.Name] = nil
	}

	for _, v := range this.participants {
		steps := &Steps{}
		id := 0
		for _, a := range this.actions {
			switch data := a.(type) {
			case *Note:
				_, ok := m[data.OverParticipant]
				if !ok {
					return m, errors.New(fmt.Sprintf("ERROR: cannot find object \"%s\"", data.OverParticipant))
				}
				if data.OverParticipant == v.Name {
					data.Id = id
					steps.Data = append(steps.Data, &Step{Name: data.Name, Font: data.Font, Type: STEP_TYPE_NOTE})
					id += 2
				}
			case *Message:
				data.Id = id
				steps.Data = append(steps.Data, &Step{Type: STEP_TYPE_NORMAL, Id: id})
				id++
			}
		}

		for id < totalSteps {
			steps.Data = append(steps.Data, &Step{Type: STEP_TYPE_NORMAL, Id: id})
			id++
		}

		m[v.Name] = steps
	}

	return m, nil
}

func (this *Sequence) getTotalStepNum() int {
	totalSteps := this.getMaxRow()

	if totalSteps == 0 {
		for _, v := range this.actions {
			if v.Type() == ACTION_TYPE_NOTE {
				totalSteps += 2
			} else {
				totalSteps++
			}
		}
	}

	return totalSteps + 1
}

func (this *Sequence) getMaxRow() int {
	maxRow := 0
	for _, v := range this.actions {
		if v.GetRow() > maxRow {
			maxRow = v.GetRow()
		}
	}

	return maxRow
}

func (this *Sequence) BuildAndGenPlantumlPng(pngfile string, config *SequenceConfig) {
	plantUmlfile := core.ReplaceFileSuffix(pngfile, "puml")
	this.BuildDotFile(plantUmlfile, config)
	this.GenPlantumlPng(plantUmlfile, pngfile, config)
}

func (this *Sequence) BuildPlantumlFile(filename string, config *SequenceConfig) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("ERROR: cannot open file %s\r\n", filename)
		return
	}
	defer file.Close()

	file.WriteString(this.BuildPlantUml(config))
}

func (this *Sequence) GenPlantumlPng(plantUmlfile, pngfile string, config *SequenceConfig) {
	cmd := exec.Command("java", "-jar", config.PlantumlJarPath+"plantuml.jar", plantUmlfile, "-v")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = "f:\\"
	err := cmd.Run()
	if err != nil {
		fmt.Println("GenPlantumlPng failed: =", err)
	}
	fmt.Println("cmd.Args =", cmd.Args)
	fmt.Println("cmd.Dir =", cmd.Dir)

}

func (this *Sequence) BuildPlantUml(config *SequenceConfig) string {
	var buf core.Writer

	//m, _ := this.getSteps()

	buf.Write(`@startuml
hide footbox
scale 700*600

skinparam Note {
	BorderColor black
	BackgroundColor white

	FontName %s
	FontColor %s
	FontSize %d
	FontStyle %s
	
}

skinparam ParticipantPadding 20
skinparam BoxPadding 10

skinparam sequence {
	ParticipantBorderColor black
	ParticipantBackgroundColor white

	ParticipantFontName %s
	ParticipantFontColor %s
	ParticipantFontSize %d
	ParticipantFontStyle %s

	LifeLineBorderColor black
	
	ArrowColor black

	ArrowFontName %s
	ArrowFontColor %s
	ArrowFontSize %d
	ArrowFontStyle %s
}

`,
		config.ParticipantFont.GetPlantumlName(), config.ParticipantFont.FrontColor.RGBString(), config.ParticipantFont.Size, config.ParticipantFont.GetStyleName(),
		config.ParticipantFont.GetPlantumlName(), config.ParticipantFont.FrontColor.RGBString(), config.ParticipantFont.Size, config.ParticipantFont.GetStyleName(),
		config.MsgFont.GetPlantumlName(), config.MsgFont.FrontColor.RGBString(), config.MsgFont.Size, config.MsgFont.GetStyleName(),
	)

	for i, v := range this.participants {
		v.Id = i
	}

	for _, v := range this.participants {
		buf.Write("participant \"%s\" as %s", v.Label, v.Name)
		if v.IsFocus {
			buf.Write(" %s", config.FocusParticipantFillColor.RGBString())
		}
		buf.Writeln("")

	}

	for _, a := range this.actions {
		switch data := a.(type) {
		case *Message:
			from := this.findParticipant(data.From)
			if from == nil {
				fmt.Printf("ERROR: cannot find object \"%s\"\n", data.From)
				return ""
			}
			to := this.findParticipant(data.To)
			if to == nil {
				fmt.Printf("ERROR: cannot find object \"%s\"\n", data.To)
				return ""
			}

			if to.Id == (from.Id+1) || from.Id == (to.Id+1) {
				buf.Writeln("%s -> %s: %s", from.Name, to.Name, data.Name)
			} else {
				buf.Writeln("%s -[%s]> %s: <back:%s>%s</back>",
					from.Name,
					config.CrossNeighborMsgColor.RGBString(),
					to.Name,
					config.TextBackgroundColor.RGBString(),
					data.Name)
			}
		case *Note:
			obj := this.findParticipant(data.OverParticipant)
			if obj == nil {
				fmt.Printf("ERROR: cannot find object \"%s\"\n", data.OverParticipant)
				return ""
			}
			buf.Writeln("hnote over %s: %s", obj.Name, data.Name)
		}
	}

	buf.Writeln("")
	buf.Writeln("@enduml")

	return buf.String()
}
