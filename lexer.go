package baal

import (
	"errors"
)

// Lexer provide inteface to parse codes.
type Lexer struct {
	s          *Scanner
	lit        string
	pos        Position
	e          error
	namespaces []Namespace
}

// Lex scan the token and literals.
func (l *Lexer) Lex(lval *yySymType) int {
	tok, lit, pos, err := l.s.Scan()
	if err != nil {
		l.e = err
	}
	if tok == EOF {
		return 0
	}
	lval.tok = Token{Pos: pos, Tok: tok, Lit: lit}
	l.lit = lit
	l.pos = pos
	return tok
}

// Error set parse error.
func (l *Lexer) Error(msg string) {
	l.e = errors.New(msg)
}

// Parser provide way to parse the code.
func Parse(s *Scanner) ([]Namespace, error) {
	l := Lexer{s: s}
	if yyParse(&l) != 0 {
		return nil, l.e
	}
	return l.namespaces, l.e
}
