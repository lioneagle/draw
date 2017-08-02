package core

import (
	"bytes"
	"fmt"
)

type Writer struct {
	buf bytes.Buffer
}

func (this *Writer) Write(format string, args ...interface{}) {
	this.buf.WriteString(fmt.Sprintf(format, args...))
}

func (this *Writer) Writeln(format string, args ...interface{}) {
	this.buf.WriteString(fmt.Sprintf(format, args...))
	this.buf.WriteString(fmt.Sprintln())
}
