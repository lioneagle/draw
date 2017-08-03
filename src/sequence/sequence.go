package sequence

import (
	"core"
	"fmt"
	"os"
	"os/exec"
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
		fmt.Println("err =", err)
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

	m, _ := this.getSteps()

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
			buf.Writeln("        obj%d[shape=\"box\", label=\"%s\", width=1, height=0.5];", v.Id, v.Name)
		} else {
			buf.Writeln("        obj%d[shape=\"box\", label=\"%s\", fillcolor=\"%s\", style=filled, width=1, height=0.5];",
				v.Id, v.Name, config.FocusParticipantFillColor.RGBString())
		}

		steps, _ := m[v.Name]
		for j, s := range steps.Data {
			if j == (len(steps.Data) - 1) {
				buf.Writeln("        obj%d_step_%d[shape=\"box\", width=0.5, label=\"\"];", v.Id, s.Id)
				break
			}
			if s.Type == STEP_TYPE_NOTE {
				buf.Writeln("        obj%d_note_%d[shape=\"circle\", label=\"%s\", width=0.51];", v.Id, s.Id, s.Name)
			}
		}

		buf.Write("        obj%d", v.Id)

		for _, s := range steps.Data {
			if s.Type == STEP_TYPE_NOTE {
				buf.Write(" -> obj%d_note_%d", v.Id, s.Id)
			} else {
				buf.Write(" -> obj%d_step_%d", v.Id, s.Id)
			}
		}

		buf.Write(";\r\n    }\r\n")
	}

	buf.Write("\r\n")

	for _, a := range this.actions {
		switch data := a.(type) {
		case *Message:
			from := this.findParticipant(data.From)
			to := this.findParticipant(data.To)

			var k int

			if to.Id > from.Id {
				if to.Id == (from.Id + 1) {
					buf.Writeln("    obj%d_step_%d -> obj%d_step_%d [label=\"%s\", arrowhead=\"normal\"];",
						from.Id, data.Id, from.Id+1, data.Id, data.Name)
				} else {
					buf.Writeln("    obj%d_step_%d -> obj%d_step_%d [label=\"%s\", color=\"%s\"];",
						from.Id, data.Id, from.Id+1, data.Id, data.Name, config.CrossNeighborMsgColor.RGBString())
					for k = from.Id + 1; k < to.Id; k++ {
						if k != (to.Id - 1) {
							buf.Writeln("    obj%d_step_%d -> obj%d_step_%d [color=\"%s\"];",
								k, data.Id, k+1, data.Id, config.CrossNeighborMsgColor.RGBString())
						} else {
							buf.Writeln("    obj%d_step_%d -> obj%d_step_%d [arrowhead=\"normal\", color=\"%s\"];",
								k, data.Id, k+1, data.Id, config.CrossNeighborMsgColor.RGBString())
						}
					}
				}

			} else {
				if to.Id == (from.Id - 1) {
					buf.Writeln("    obj%d_step_%d -> obj%d_step_%d [label=\"%s\", arrowhead=\"normal\"];",
						from.Id, data.Id, from.Id-1, data.Id, data.Name)
				} else {
					buf.Writeln("    obj%d_step_%d -> obj%d_step_%d [label=\"%s\", color=\"%s\", ];",
						from.Id, data.Id, from.Id-1, data.Id, data.Name, config.CrossNeighborMsgColor.RGBString())
					for k = from.Id - 1; k > to.Id; k-- {
						if k != (to.Id + 1) {
							buf.Writeln("    obj%d_step_%d -> obj%d_step_%d [color=\"%s\"];",
								k, data.Id, k-1, data.Id, config.CrossNeighborMsgColor.RGBString())
						} else {
							buf.Writeln("    obj%d_step_%d -> obj%d_step_%d [arrowhead=\"normal\", color=\"%s\"];",
								k, data.Id, k-1, data.Id, config.CrossNeighborMsgColor.RGBString())
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
		steps := &Steps{}
		id := 0
		for _, a := range this.actions {
			switch data := a.(type) {
			case *Note:
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

func (this *Sequence) BuildPlantUml() string {

	return ""
}
