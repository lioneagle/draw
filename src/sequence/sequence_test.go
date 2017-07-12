package sequence

import (
	"core"
	"fmt"
	"testing"
)

func TestSequenceBuildDot(t *testing.T) {
	seq := &Sequence{}

	color := core.NewColor()
	objectFont := &core.Font{Name: "ObjectFont", Style: "FangSong", Color: *color, Size: 14}
	messageFont := &core.Font{Name: "MessageFont", Style: "FangSong", Color: *color, Size: 9}
	noteFont := &core.Font{Name: "NodeFont", Style: "FangSong", Color: *color, Size: 9}

	seq.AddObject(&Object{Name: "UE", Font: objectFont})
	seq.AddObject(&Object{Name: "ZXUN B200", Font: objectFont, IsMain: true})
	seq.AddObject(&Object{Name: "I/S-CSCF", Font: objectFont})
	seq.AddObject(&Object{Name: "SCC AS", Font: objectFont})

	seq.AddMessage(&Message{FromObj: "UE", ToObj: "ZXUN B200", Name: "INVITE", Font: messageFont, Seq: 1})
	seq.AddMessage(&Message{FromObj: "ZXUN B200", ToObj: "I/S-CSCF", Name: "INVITE", Font: messageFont, Seq: 2})
	seq.AddMessage(&Message{FromObj: "I/S-CSCF", ToObj: "SCC AS", Name: "INVITE", Font: messageFont, Seq: 2})
	seq.AddMessage(&Message{FromObj: "SCC AS", ToObj: "I/S-CSCF", Name: "INVITE 180", Font: messageFont, Seq: 3})
	seq.AddMessage(&Message{FromObj: "I/S-CSCF", ToObj: "ZXUN B200", Name: "INVITE 180", Font: messageFont, Seq: 3})
	seq.AddMessage(&Message{FromObj: "ZXUN B200", ToObj: "UE", Name: "INVITE 180", Font: messageFont, Seq: 4})
	seq.AddMessage(&Message{FromObj: "SCC AS", ToObj: "UE", Name: "test1", Font: messageFont, Seq: 4})
	seq.AddMessage(&Message{FromObj: "UE", ToObj: "SCC AS", Name: "test2", Font: messageFont, Seq: 4})

	seq.AddNote(&Note{Obj: "ZXUN B200", Name: "A", Font: noteFont})

	dot := seq.BuildDot()

	fmt.Print("dot =\n", dot)
}
