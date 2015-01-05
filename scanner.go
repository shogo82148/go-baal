package baal

import (
	"errors"
	"fmt"
)

const (
	EOF = -1   // End of file.
	EOL = '\n' // End of line.
)

type Scanner struct {
	src      []rune
	offset   int
	lineHead int
	line     int
}

var keywords = map[string]int{
	"namespace":  NAMESPACE,
	"import":     IMPORT,
	"abstract":   ABSTRACT,
	"entity":     ENTITY,
	"service":    SERVICE,
	"includes":   INCLUDES,
	"required":   REQUIRED,
	"nullable":   NULLABLE,
	"list":       LIST_OF,
	"dictionary": DICTIONARY_OF,
	"hash":       DICTIONARY_OF,
	"map":        DICTIONARY_OF,
	"accepts":    ACCEPTS,
	"returns":    RETURNS,
	"as":         AS,

	"boolean":   BOOLEAN,
	"bool":      BOOLEAN,
	"int8":      INTEGER8,
	"sbyte":     INTEGER8,
	"byte":      INTEGER8,
	"int16":     INTEGER16,
	"short":     INTEGER16,
	"int32":     INTEGER32,
	"integer":   INTEGER32,
	"int":       INTEGER32,
	"int64":     INTEGER64,
	"long":      INTEGER64,
	"float32":   FLOAT32,
	"float":     FLOAT32,
	"float64":   FLOAT64,
	"double":    FLOAT64,
	"real":      FLOAT64,
	"number":    FLOAT64,
	"decimal32": DECIMAL32,
	"decimal64": DECIMAL64,
	"decimal":   DECIMAL64,
	"numeric":   DECIMAL64,
	"money":     DECIMAL64,
	"date":      DATE,
	"time":      TIME,
	"timestamp": TIMESTAMP,
	"datetime":  TIMESTAMP,
	"string":    STRING,
	"str":       STRING,
	"binary":    BINARY,
	"bin":       BINARY,
}

// Init reset code to scan.
func (s *Scanner) Init(src string) {
	s.src = []rune(src)
}

func (s *Scanner) Scan() (tok int, lit string, pos Position, err error) {
retry:
	s.skipBlank()
	switch ch := s.peek(); {
	case isIdentifierFirstCharacter(ch):
		tok = IDENTIFIER
		lit, err = s.scanIdentifier()
		return
	case isKeywordFirstCharacter(ch):
		lit, err = s.scanIdentifier()
		if err != nil {
			return
		}
		if name, ok := keywords[lit]; ok {
			tok = name
		} else {
			err = errors.New("unknown keyword")
			return
		}
		if tok == LIST_OF || tok == DICTIONARY_OF {
			s.skipBlank()
			if keywordOf, err := s.scanIdentifier(); err != nil || keywordOf != "of" {
				err = errors.New("missing `of'")
			}
		}
		return
	default:
		switch ch {
		case -1:
			tok = EOF
		case ':':
			tok = AS
			lit = ":"
		case '!':
			tok = REQUIRED
			lit = "!"
		case '?':
			tok = NULLABLE
			lit = "?"
		case '@':
			tok = LIST_OF
			lit = "@"
		case '%':
			tok = DICTIONARY_OF
			lit = "%"
		case '+':
			s.next()
			if s.peek() != '=' {
				err = fmt.Errorf(`syntax error "%s"`, string(ch))
				return
			}
			tok = APPEND
			lit = "+="
		case '-':
			s.next()
			if s.peek() != '=' {
				err = fmt.Errorf(`syntax error "%s"`, string(ch))
				return
			}
			tok = REMOVE
			lit = "-="
		case '<':
			s.next()
			if s.peek() != '=' {
				err = fmt.Errorf(`syntax error "%s"`, string(ch))
				return
			}
			tok = ACCEPTS
			lit = "<="
		case '=':
			s.next()
			if s.peek() != '>' {
				err = fmt.Errorf(`syntax error "%s"`, string(ch))
				return
			}
			tok = RETURNS
			lit = "=>"
		case '/':
			s.next()
			switch s.peek() {
			case '/':
				for !isEOL(s.peek()) {
					s.next()
				}
				goto retry
			case '*':
				_, err = s.scanComment('*')
				if err != nil {
					return
				}
				goto retry
			case '#':
				tok = DOCUMENT_COMMENT
				lit, err = s.scanComment('#')
			default:
				err = fmt.Errorf(`syntax error "%s"`, string(ch))
				return
			}
		case '{':
			tok = BEGIN
			lit = "{"
		case '}':
			tok = END
			lit = "}"
		case ';':
			tok = CLOSE
			lit = ";"
		case '.', '*':
			tok = int(ch)
			lit = string(ch)
		default:
			err = fmt.Errorf(`syntax error "%s"`, string(ch))
		}
		s.next()
	}
	return
}

// next move offset to next.
func (s *Scanner) next() {
	if !s.reachEOF() {
		if s.peek() == '\n' {
			s.lineHead = s.offset + 1
			s.line++
		}
		s.offset++
	}
}

// current return the current offset.
func (s *Scanner) current() int {
	return s.offset
}

// offset set the offset value.
func (s *Scanner) set(o int) {
	s.offset = o
}

// back move back offset once to top.
func (s *Scanner) back() {
	s.offset--
}

// reachEOF return true if offset is at end-of-file.
func (s *Scanner) reachEOF() bool {
	return len(s.src) <= s.offset
}

// isLetter return true if the rune is a letter for identity.
func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

// isDigit return true if the rune is a number.
func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

// isIdentifierFirstCharacter returns true if the rune can be a first character of an identifier.
func isIdentifierFirstCharacter(ch rune) bool {
	return 'A' <= ch && ch <= 'Z'
}

// isKeywordFirstCharacter returns true if the rune can be a first character of an identifier.
func isKeywordFirstCharacter(ch rune) bool {
	return 'a' <= ch && ch <= 'z'
}

// isBlank returns true if the rune is empty character.
func isBlank(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

// isEOL return true if the rune is at end-of-line or end-of-file.
func isEOL(ch rune) bool {
	return ch == '\n' || ch == -1
}

// scanIdentifier return identifier begining at current position.
func (s *Scanner) scanIdentifier() (string, error) {
	var ret []rune
	for {
		if !isLetter(s.peek()) && !isDigit(s.peek()) {
			break
		}
		ret = append(ret, s.peek())
		s.next()
	}
	return string(ret), nil
}

func (s *Scanner) scanComment(l rune) (string, error) {
	var ret []rune
eoc:
	for {
		s.next()
		switch s.peek() {
		case l:
			s.next()
			if s.peek() == '/' {
				s.next()
				break eoc
			}
			s.back()
			ret = append(ret, '#')
		default:
			ret = append(ret, s.peek())
		}
	}
	return string(ret), nil
}

// peek returns current rune in the code.
func (s *Scanner) peek() rune {
	if !s.reachEOF() {
		return s.src[s.offset]
	} else {
		return -1
	}
}

// skipBlank moves position into non-blank character.
func (s *Scanner) skipBlank() {
	for isBlank(s.peek()) {
		s.next()
	}
}
