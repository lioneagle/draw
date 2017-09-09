package parser

import (
	//"fmt"
	"testing"
)

func TestLexerScanNumber(t *testing.T) {
	testdata := []struct {
		str    string
		token  Token
		offset int
		lit    string
	}{
		{"123.a", FLOAT, 4, "123."},
		{"123.", FLOAT, 4, "123."},
		{".1", FLOAT, 2, ".1"},
		{"5.11", FLOAT, 4, "5.11"},
		{"1e+11", FLOAT, 5, "1e+11"},
		{"1e0", FLOAT, 3, "1e0"},
		{".45e-11", FLOAT, 7, ".45e-11"},
	}

	prefix := "TestLexerScanNumber"

	for i, v := range testdata {
		var lexer Lexer

		file := &File{}
		file.size = len(v.str)

		lexer.Init(file, []byte(v.str))
		pos, token, lit := lexer.Scan()

		if token != v.token {
			t.Errorf("%s[%d] failed: token = %s, wanted = %s\n", prefix, i, token, v.token)
		}

		if lit != v.lit {
			t.Errorf("%s[%d] failed: lit = %s, wanted = %s\n", prefix, i, lit, v.lit)
		}

		if pos != 0 {
			t.Errorf("%s[%d] failed: pos = %d, wanted = 0\n", prefix, i, pos)
		}

		if lexer.offset != v.offset {
			t.Errorf("%s[%d] failed: offset = %d, wanted = %d\n", prefix, i, lexer.offset, v.offset)
		}
	}

}
