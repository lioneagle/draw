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
