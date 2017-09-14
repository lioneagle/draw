package parser

import (
	//"fmt"
	"testing"
)

func TestLexerScan(t *testing.T) {
	testdata := []struct {
		str    string
		token  Token
		offset int
		lit    string
	}{
		{"123a", INT, 3, "123"},
		{"0a", INT, 1, "0"},
		{"0123456789a", INT, 10, "0123456789"},
		{"0xabg", INT, 4, "0xab"},
		{"0x01aBg", INT, 6, "0x01aB"},
		{"123.a", FLOAT, 4, "123."},
		{"123.", FLOAT, 4, "123."},
		{".1", FLOAT, 2, ".1"},
		{"0.3", FLOAT, 3, "0.3"},
		{"5.11", FLOAT, 4, "5.11"},
		{"1e+11", FLOAT, 5, "1e+11"},
		{"1e0", FLOAT, 3, "1e0"},
		{".45e-11", FLOAT, 7, ".45e-11"},

		{`"123a"`, STRING, 6, `"123a"`},
		{`"0123a"`, STRING, 7, `"0123a"`},
		{`"0123a按时as"`, STRING, 15, `"0123a按时as"`},
		{`"_as0123a"`, STRING, 10, `"_as0123a"`},
		{`"_as\a0123a"`, STRING, 12, `"_as\a0123a"`},
		{`"_as\x0123a"`, STRING, 12, `"_as\x0123a"`},
		{`"_as\ua123a"`, STRING, 12, `"_as\ua123a"`},
		{`"_as\Ua123a"`, STRING, 12, `"_as\Ua123a"`},
		{`"_as\123a"`, STRING, 10, `"_as\123a"`},

		{"ahgc_", IDENT, 5, "ahgc_"},
		{"_0Asa_A0,", IDENT, 8, "_0Asa_A0"},

		{"= ", ASSIGN, 1, "="},
		{", ", COMMA, 1, ","},
		{"( ", LPAREN, 1, "("},
		{"[ ", LBRACK, 1, "["},
		{"{ ", LBRACE, 1, "{"},
		{". ", PERIOD, 1, "."},
		{") ", RPAREN, 1, ")"},
		{"] ", RBRACK, 1, "]"},
		{"} ", RBRACE, 1, "}"},
		{"; ", SEMICOLON, 1, ";"},
		{": ", COLON, 1, ":"},
		{"-> ", ARROW, 2, "->"},
		{"\n ", NEWLINE, 1, "\n"},

		{"sequence,", SEQUENCE, 8, "sequence"},
		{"participant ", PARTICIPANT, 11, "participant"},
		{"note ", NOTE, 4, "note"},
		{"end_note ", END_NOTE, 8, "end_note"},
		{"end_message ", END_MESSAGE, 11, "end_message"},
		{"over ", OVER, 4, "over"},
		{"left ", LEFT, 4, "left"},
		{"right ", RIGHT, 5, "right"},
		{"name ", NAME, 4, "name"},
		{"label ", LABEL, 5, "label"},
		{"config ", CONFIG, 6, "config"},
		{"font ", FONT, 4, "font"},
		{"size ", SIZE, 4, "size"},
		{"style ", STYLE, 5, "style"},
		{"bold ", BOLD, 4, "bold"},
		{"underline ", UNDERLINE, 9, "underline"},
		{"italic ", ITALIC, 6, "italic"},
		{"color ", COLOR, 5, "color"},
		{"backcolor ", BACKCOLOR, 9, "backcolor"},
		{"focus ", FOCUS, 5, "focus"},
		//*/
	}

	prefix := "TestLexerScan"

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
