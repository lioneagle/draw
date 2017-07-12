package sequence

import (
	"bytes"
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

type Object struct {
	Name   string
	Id     int
	Font   *core.Font
	IsMain bool
}

func (this *Object) Equal(other *Object) bool {
	return this.Name == other.Name
}

const (
	ACTION_TYPE_MESSAGE = iota
	ACTION_TYPE_NOTE
)

type Action interface {
	IsAction() bool
	Type() int
	GetRow() int
	GetId() int
}

type Message struct {
	FromObj string
	ToObj   string
	Name    string
	Font    *core.Font
	Row     int
	Seq     int
	Id      int
}

func (this *Message) IsAction() bool { return true }
func (this *Message) Type() int      { return ACTION_TYPE_MESSAGE }
func (this *Message) GetRow() int    { return this.Row }
func (this *Message) GetId() int     { return this.Id }

type Note struct {
	Obj  string
	Name string
	Font *core.Font
	Id   int
	Row  int
}

func (this *Note) IsAction() bool { return true }
func (this *Note) Type() int      { return ACTION_TYPE_NOTE }
func (this *Note) GetRow() int    { return this.Row }
func (this *Note) GetId() int     { return this.Id }

type Sequence struct {
	objects []*Object
	actions []Action
}

func (this *Sequence) AddObject(obj *Object) {
	v := this.findObject(obj.Name)
	if v == nil {
		this.objects = append(this.objects, obj)
	}
}

func (this *Sequence) findObject(name string) *Object {
	for _, v := range this.objects {
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

func (this *Sequence) BuildAndGenDotPng(pngfile string) {
	dotfile := core.ReplaceFileSuffix(pngfile, "gv")
	this.BuildDotFile(dotfile)
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

func (this *Sequence) BuildDotFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("ERROR: cannot open file %s\r\n", filename)
		return
	}
	defer file.Close()

	file.WriteString(this.BuildDot())
}

func (this *Sequence) BuildDot() string {
	var buf bytes.Buffer

	m, _ := this.getSteps()

	buf.WriteString(`digraph G {
    rankdir="LR";
    node[shape="point", width=0, height=0];
    edge[arrowhead="none", style="solid"]
	
	`)

	for i, v := range this.objects {
		v.Id = i
	}

	for _, v := range this.objects {
		buf.WriteString("\r\n")
		buf.WriteString("    {\r\n")
		buf.WriteString("        rank=\"same\";\r\n")
		buf.WriteString("        edge[style=\"solid\"];\r\n")
		if !v.IsMain {
			buf.WriteString(fmt.Sprintf("        obj%d[shape=\"record\", label=\"%s\"];\r\n", v.Id, v.Name))
		} else {
			buf.WriteString(fmt.Sprintf("        obj%d[shape=\"record\", label=\"%s\", fillcolor=\"%s\", style=filled];\r\n",
				v.Id, v.Name, core.ColorGold.String()))
		}

		steps, _ := m[v.Name]
		for j, s := range steps.Data {
			if j == (len(steps.Data) - 1) {
				buf.WriteString(fmt.Sprintf("        obj%d_step_%d[shape=\"\", width=0.5, label=\"\"];\r\n", v.Id, s.Id))
				break
			}
			if s.Type == STEP_TYPE_NOTE {
				buf.WriteString(fmt.Sprintf("        obj%d_note_%d[shape=\"circle\", label=\"%s\"];\r\n", v.Id, s.Id, s.Name))
			}
		}

		buf.WriteString(fmt.Sprintf("        obj%d", v.Id))

		for _, s := range steps.Data {
			if s.Type == STEP_TYPE_NOTE {
				buf.WriteString(fmt.Sprintf(" -> obj%d_note_%d", v.Id, s.Id))
			} else {
				buf.WriteString(fmt.Sprintf(" -> obj%d_step_%d", v.Id, s.Id))
			}
		}

		buf.WriteString(";\r\n    }\r\n")
	}

	buf.WriteString("\r\n")

	for _, a := range this.actions {
		switch data := a.(type) {
		case *Message:
			from := this.findObject(data.FromObj)
			to := this.findObject(data.ToObj)

			var k int

			if to.Id > from.Id {
				if to.Id == (from.Id + 1) {
					buf.WriteString(fmt.Sprintf("    obj%d_step_%d -> obj%d_step_%d [label=\"%s\", arrowhead=\"normal\"];\r\n",
						from.Id, data.Id, from.Id+1, data.Id, data.Name))
				} else {
					buf.WriteString(fmt.Sprintf("    obj%d_step_%d -> obj%d_step_%d [label=\"%s\", color=\"%s\", color=\"%s\"];\r\n",
						from.Id, data.Id, from.Id+1, data.Id, data.Name, core.ColorAqua.String(), core.ColorAqua.String()))
					for k = from.Id + 1; k < to.Id; k++ {
						if k != (to.Id - 1) {
							buf.WriteString(fmt.Sprintf("    obj%d_step_%d -> obj%d_step_%d [color=\"%s\"];\r\n",
								k, data.Id, k+1, data.Id, core.ColorAqua.String()))
						} else {
							buf.WriteString(fmt.Sprintf("    obj%d_step_%d -> obj%d_step_%d [arrowhead=\"normal\", color=\"%s\"];\r\n",
								k, data.Id, k+1, data.Id, core.ColorAqua.String()))
						}
					}
				}

			} else {
				if to.Id == (from.Id - 1) {
					buf.WriteString(fmt.Sprintf("    obj%d_step_%d -> obj%d_step_%d [label=\"%s\", arrowhead=\"normal\"];\r\n",
						from.Id, data.Id, from.Id-1, data.Id, data.Name))
				} else {
					buf.WriteString(fmt.Sprintf("    obj%d_step_%d -> obj%d_step_%d [label=\"%s\", color=\"%s\", color=\"%s\"];\r\n",
						from.Id, data.Id, from.Id-1, data.Id, data.Name, core.ColorAqua.String(), core.ColorAqua.String()))
					for k = from.Id - 1; k > to.Id; k-- {
						if k != (to.Id + 1) {
							buf.WriteString(fmt.Sprintf("    obj%d_step_%d -> obj%d_step_%d [color=\"%s\"];\r\n",
								k, data.Id, k-1, data.Id, core.ColorAqua.String()))
						} else {
							buf.WriteString(fmt.Sprintf("    obj%d_step_%d -> obj%d_step_%d [arrowhead=\"normal\", color=\"%s\"];\r\n",
								k, data.Id, k-1, data.Id, core.ColorAqua.String()))
						}
					}
				}
			}
		}
	}

	buf.WriteString("}\r\n")

	return buf.String()
}

func (this *Sequence) getSteps() (map[string]*Steps, error) {
	m := make(map[string]*Steps)
	totalSteps := this.getTotalStepNum()

	for _, v := range this.objects {
		steps := &Steps{}
		id := 0
		for _, a := range this.actions {
			switch data := a.(type) {
			case *Note:
				if data.Obj == v.Name {
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
