package parser

import (
	"fmt"
	"strconv"
	"unicode"
	"unicode/utf8"
)

const bom = 0xFEFF

type Token int

const (
	ILLEGAL Token = iota
	EOF
	COMMENT

	IDENT     // identifier
	INT       // 12345
	FLOAT     // 123.45
	STRING    // "abc"
	ASSIGN    // =
	COMMA     // ,
	LPAREN    // (
	LBRACK    // [
	LBRACE    // {
	PERIOD    // .
	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	SEMICOLON // ;
	COLON     // :
	ARROW     // ->
	NEWLINE   // '\n'

	keyword_beg
	SEQUENCE    // sequence
	PARTICIPANT // participant
	NOTE        // note
	END_NOTE    // end_note
	END_MESSAGE // end_message
	OVER        // over
	LEFT        // left
	RIGHT       // right
	NAME        // name
	LABEL       // label
	CONFIG      // config
	FONT        // font
	SIZE        // size
	STYLE       // style
	BOLD        // bold
	UNDERLINE   // underline
	ITALIC      // italic
	COLOR       // color
	BACKCOLOR   // backcolor
	FOCUS       // focus
	keyword_end
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	COMMENT: "COMMENT",

	IDENT:     "IDENT",
	INT:       "INT",
	FLOAT:     "FLOAT",
	STRING:    "STRING",
	ASSIGN:    "=",
	COMMA:     ",",
	LPAREN:    "(",
	LBRACK:    "[",
	LBRACE:    "{",
	PERIOD:    ".",
	RPAREN:    ")",
	RBRACK:    "]",
	RBRACE:    "}",
	SEMICOLON: ";",
	COLON:     ":",
	ARROW:     "->",
	NEWLINE:   "NEWLINE",

	SEQUENCE:    "sequence",
	PARTICIPANT: "participant",
	NOTE:        "note",
	END_NOTE:    "end_note",
	END_MESSAGE: "end_message",
	OVER:        "over",
	LEFT:        "left",
	RIGHT:       "right",
	NAME:        "name",
	LABEL:       "label",
	CONFIG:      "config",
	FONT:        "font",
	SIZE:        "size",
	STYLE:       "style",
	BOLD:        "bold",
	UNDERLINE:   "underline",
	ITALIC:      "italic",
	COLOR:       "color",
	BACKCOLOR:   "backcolor",
	FOCUS:       "focus",
}

func (tok Token) String() string {
	s := ""
	if 0 <= tok && tok < Token(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch >= utf8.RuneSelf && unicode.IsLetter(ch)
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9' || ch >= utf8.RuneSelf && unicode.IsDigit(ch)
}

var keywords map[string]Token

func init() {
	keywords = make(map[string]Token)
	for i := keyword_beg + 1; i < keyword_end; i++ {
		keywords[tokens[i]] = i
	}
}

func Lookup(ident string) Token {
	if tok, is_keyword := keywords[ident]; is_keyword {
		return tok
	}
	return IDENT
}

func (tok Token) IsKeyword() bool { return keyword_beg < tok && tok < keyword_end }

type File struct {
	name  string
	size  int
	base  int // Pos value range for this file is [base...base+size]
	lines []int
	infos []lineInfo
}

type Pos int

// Name returns the file name of file  as registered with AddFile.
func (f *File) Name() string {
	return f.name
}

// Size returns the size of file  as registered with AddFile.
func (this *File) Size() int {
	return this.size
}

// Pos returns the Pos value for the given file offset;
// the offset must be <= f.Size().
// f.Pos(f.Offset(p)) == p.
//
func (this *File) Pos(offset int) Pos {
	if offset > this.size {
		panic("illegal file offset")
	}
	return Pos(this.base + offset)
}

// Offset returns the offset for the given file position p;
// p must be a valid Pos value in that file.
// f.Offset(f.Pos(offset)) == offset.
//
func (this *File) Offset(p Pos) int {
	if int(p) < this.base || int(p) > (this.base+this.size) {
		panic("illegal Pos value")
	}
	return int(p) - this.base
}

// AddLine adds the line offset for a new line.
// The line offset must be larger than the offset for the previous line
// and smaller than the file size; otherwise the line offset is ignored.
//
func (this *File) AddLine(offset int) {
	if i := len(this.lines); (i == 0 || this.lines[i-1] < offset) && offset < this.size {
		this.lines = append(this.lines, offset)
	}
}

// A lineInfo object describes alternative file and line number
// information (such as provided via a //line comment in a .go
// file) for a given file offset.
type lineInfo struct {
	// fields are exported to make them accessible to gob
	Offset   int
	Filename string
	Line     int
}

// AddLineInfo adds alternative file and line number information for
// a given file offset. The offset must be larger than the offset for
// the previously added alternative line info and smaller than the
// file size; otherwise the information is ignored.
//
// AddLineInfo is typically used to register alternative position
// information for //line filename:line comments in source files.
//
func (this *File) AddLineInfo(offset int, filename string, line int) {
	if i := len(this.infos); i == 0 || this.infos[i-1].Offset < offset && offset < this.size {
		this.infos = append(this.infos, lineInfo{offset, filename, line})
	}
}

type Lexer struct {
	file *File
	src  []byte

	ch          rune // current character
	offset      int  // character offset
	readOffset  int  // reading offset (position after current character)
	lineOffset  int  // current line offset
	currentLine int
	ErrorCount  int
}

func (this *Lexer) Init(file *File, src []byte) {
	if file.Size() != len(src) {
		panic(fmt.Sprintf("file size (%d) does not match src len (%d)", file.Size(), len(src)))
	}

	this.file = file
	this.src = src
	this.ch = ' '
	this.offset = 0
	this.readOffset = 0
	this.lineOffset = 0
	this.currentLine = 0
	this.ErrorCount = 0

	this.next()
	if this.ch == bom {
		this.next() // ignore BOM at file beginning
	}
}

func (this *Lexer) Scan() (pos Pos, token Token, lit string) {
	this.skipWhitespace()

	pos = this.file.Pos(this.offset)

	switch ch := this.ch; {
	case isLetter(ch):
		lit := this.scanIdentifier()
		if len(lit) > 1 {
			token = Lookup(lit)
		} else {
			token = IDENT
		}
	case '0' <= ch && ch <= '9':
		token, lit = this.scanNumber(false)
	default:
		this.next() // always make progress
		switch ch {
		case -1:
			token = EOF
		case '\n':
			token = NEWLINE
		case '"':
			token = STRING
			lit = this.scanString()
		case '=':
			token = ASSIGN
		case ',':
			token = COMMA
		case '(':
			token = LPAREN
		case '[':
			token = LBRACK
		case '{':
			token = LBRACE
		case '.':
			if this.ch >= '0' && this.ch <= '9' {
				token, lit = this.scanNumber(true)
			} else {
				token = PERIOD
			}
		case ')':
			token = RPAREN
		case ']':
			token = RBRACK
		case '}':
			token = RBRACE
		case ';':
			token = SEMICOLON
		case ':':
			token = COLON
		case '-':
			if this.ch == '>' {
				token = ARROW
			} else {
				this.error(this.file.Offset(pos), fmt.Sprintf("illegal character %#U", ch))
				token = ILLEGAL
				lit = string(ch)
			}
		default:
			this.error(this.file.Offset(pos), fmt.Sprintf("illegal character %#U", ch))
			token = ILLEGAL
			lit = string(ch)
		}
	}

	return
}

func (this *Lexer) scanString() string {
	// '"' opening already consumed
	offs := this.offset - 1

	for {
		ch := this.ch
		if ch == '\n' || ch < 0 {
			this.error(offs, "string literal not terminated")
			break
		}
		this.next()
		if ch == '"' {
			break
		}
		if ch == '\\' {
			this.scanEscape('"')
		}
	}

	return string(this.src[offs:this.offset])
}

// scanEscape parses an escape sequence where rune is the accepted
// escaped quote. In case of a syntax error, it stops at the offending
// character (without consuming it) and returns false. Otherwise
// it returns true.
func (this *Lexer) scanEscape(quote rune) bool {
	offs := this.offset

	var n int
	var base, max uint32
	switch this.ch {
	case 'a', 'b', 'f', 'n', 'r', 't', 'v', '\\', quote:
		this.next()
		return true
	case '0', '1', '2', '3', '4', '5', '6', '7':
		n, base, max = 3, 8, 255
	case 'x':
		this.next()
		n, base, max = 2, 16, 255
	case 'u':
		this.next()
		n, base, max = 4, 16, unicode.MaxRune
	case 'U':
		this.next()
		n, base, max = 8, 16, unicode.MaxRune
	default:
		msg := "unknown escape sequence"
		if this.ch < 0 {
			msg = "escape sequence not terminated"
		}
		this.error(offs, msg)
		return false
	}

	var x uint32
	for n > 0 {
		d := uint32(digitVal(this.ch))
		if d >= base {
			msg := fmt.Sprintf("illegal character %#U in escape sequence", this.ch)
			if this.ch < 0 {
				msg = "escape sequence not terminated"
			}
			this.error(this.offset, msg)
			return false
		}
		x = x*base + d
		this.next()
		n--
	}

	if x > max || 0xD800 <= x && x < 0xE000 {
		this.error(offs, "escape sequence is invalid Unicode code point")
		return false
	}

	return true
}

func (this *Lexer) scanNumber(seenDecimalPoint bool) (Token, string) {
	// digitVal(s.ch) < 10
	offs := this.offset
	tok := INT

	if seenDecimalPoint {
		offs--
		tok = FLOAT
		this.scanMantissa(10)
		goto exponent
	}

	if this.ch == '0' {
		// int or float
		offs := this.offset
		this.next()
		if this.ch == 'x' || this.ch == 'X' {
			// hexadecimal int
			this.next()
			this.scanMantissa(16)
			if this.offset-offs <= 2 {
				// only scanned "0x" or "0X"
				this.error(offs, "illegal hexadecimal number")
			}
		} else {
			// octal int or float
			seenDecimalDigit := false
			this.scanMantissa(8)
			if this.ch == '8' || this.ch == '9' {
				// illegal octal int or float
				seenDecimalDigit = true
				this.scanMantissa(10)
			}
			if this.ch == '.' || this.ch == 'e' || this.ch == 'E' || this.ch == 'i' {
				goto fraction
			}
			// octal int
			if seenDecimalDigit {
				this.error(offs, "illegal octal number")
			}
		}
		goto exit
	}

	// decimal int or float
	this.scanMantissa(10)

fraction:
	if this.ch == '.' {
		tok = FLOAT
		this.next()
		this.scanMantissa(10)
	}

exponent:
	if this.ch == 'e' || this.ch == 'E' {
		tok = FLOAT
		this.next()
		if this.ch == '-' || this.ch == '+' {
			this.next()
		}
		if digitVal(this.ch) < 10 {
			this.scanMantissa(10)
		} else {
			this.error(offs, "illegal floating-point exponent")
		}
	}

exit:
	return tok, string(this.src[offs:this.offset])
}

func (this *Lexer) scanMantissa(base int) {
	for digitVal(this.ch) < base {
		this.next()
	}
}

func digitVal(ch rune) int {
	switch {
	case '0' <= ch && ch <= '9':
		return int(ch - '0')
	case 'a' <= ch && ch <= 'f':
		return int(ch - 'a' + 10)
	case 'A' <= ch && ch <= 'F':
		return int(ch - 'A' + 10)
	}
	return 16 // larger than any legal digit val
}

func (this *Lexer) scanIdentifier() string {
	begin := this.offset
	for isLetter(this.ch) || isDigit(this.ch) {
		this.next()
	}
	return string(this.src[begin:this.offset])
}

func (this *Lexer) skipWhitespace() {
	for this.ch == ' ' || this.ch == '\t' || this.ch == '\n' || this.ch == '\r' {
		this.next()
	}
}

func (this *Lexer) next() {
	if this.readOffset < len(this.src) {
		this.offset = this.readOffset
		if this.ch == '\n' {
			this.lineOffset = this.offset
			this.file.AddLine(this.offset)
		}

		r, w := rune(this.src[this.readOffset]), 1
		switch {
		case r == 0:
			this.error(this.offset, "illegal character NUL")
		case r >= utf8.RuneSelf:
			// not ASCII
			r, w = utf8.DecodeRune(this.src[this.readOffset:])
			if r == utf8.RuneError && w == 1 {
				this.error(this.offset, "illegal UTF-8 encoding")
			} else if r == bom && this.offset > 0 {
				this.error(this.offset, "illegal byte order mark")
			}
		}
		this.readOffset += w
		this.ch = r
	} else {
		this.offset = len(this.src)
		if this.ch == '\n' {
			this.lineOffset = this.offset
			this.file.AddLine(this.offset)
		}
		this.ch = -1 // eof
	}
}

func (this *Lexer) error(offset int, msg string) {
	fmt.Printf(`%s:%d:%d %s`, this.file.Name(), this.currentLine, offset-this.lineOffset, msg)
	this.ErrorCount++
}

func (this *Lexer) AddLine(offset int) {
	this.currentLine++
	this.file.AddLine(offset)
}
