package core

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func FileEqual(filename1, filename2 string) bool {
	file1, err := ioutil.ReadFile(filename1)
	if err != nil {
		fmt.Printf("ERROR: cannot open file %s\r\n", filename1)
		return false
	}

	file2, err := ioutil.ReadFile(filename2)
	if err != nil {
		fmt.Printf("ERROR: cannot open file %s\r\n", filename2)
		return false
	}

	return bytes.Equal(file1, file2)
}

func ReplaceFileSuffix(filename, newSuffix string) string {
	base := filepath.Base(filename)
	ext := filepath.Ext(base)
	newName := strings.TrimSuffix(filename, ext)
	return fmt.Sprintf("%s.%s", newName, newSuffix)
}
