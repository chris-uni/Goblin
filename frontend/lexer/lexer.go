package lexer

import (
	"fmt"
	"unicode"
)

type Lexer struct {
	Source []rune

	Pos  int
	Line int
	Col  int

	Audit  map[int]string
	Tokens []Token
}

// API

/*
Non-consuming. Returns the current rune in the Source list.
*/
func (l *Lexer) peek() rune {

	return l.Source[l.Pos]
}

/*
Non-consuming. Returns the next rune (if available) in the Source list.
*/
func (l *Lexer) peekNext() (rune, error) {

	newPos := l.Pos + 1

	if newPos >= len(l.Source) {

		return 0, fmt.Errorf("peekNext cannot overflow past end of file.")
	}

	return l.Source[newPos], nil
}

/*
Non-consuming. Returns the next n'th rune (if available) in the Source list.
*/
func (l *Lexer) peekNextN(n int) (string, error) {

	newPos := l.Pos + n

	if newPos > len(l.Source) {

		return "", fmt.Errorf("peekNextN cannot overflow past end of file.")
	}

	return string(l.Source[l.Pos:newPos]), nil
}

/*
Consuming. Returns the current rune in the Source list and incremenets Source pointer, line and col indexes.
*/
func (l *Lexer) consume() string {

	value := l.Source[l.Pos]

	l.Pos++

	// New line detection.
	if value == '\n' {
		l.Line++
		l.Col = 0
	} else {
		l.Col++
	}

	return string(value)
}

/*
Consuming. Wrapper for __consume__, called 'n' number of times. Collects returned runes and returns as string.
*/
func (l *Lexer) consumeN(n int) string {

	s := ""
	for i := 0; i < n; i++ {

		s += l.consume()
	}

	return s
}

/*
Non-consuming. Returns true if the argument matches the value of __peek__.
*/
func (l *Lexer) match(ch rune) bool {

	return l.peek() == ch
}

/*
Non-consuming. Returns true if pointer __Pos__ is at the end of the Source list.
*/
func (l *Lexer) EOF() bool {

	return l.Pos >= len(l.Source)
}

/*
Non-consuming. Returns a new Token object.
*/
func (l *Lexer) createToken(tkntyp TokenType, val string, line int, col int) Token {

	l.Audit[l.Line] += val

	return Token{
		Type:  tkntyp,
		Value: val,
		Line:  line,
		Col:   col,
	}
}

// Methods.

/*
Characters that are effectively ignored in Goblin.
*/
func (l *Lexer) isSkippable(ch rune) bool {
	return ch == ' ' || ch == '\n' || ch == '\t' || ch == '\r'
}

func (l *Lexer) processNumber() {

	start := l.Pos

	for !l.EOF() && unicode.IsDigit(l.peek()) {
		l.consume()
	}

	value := l.Source[start:l.Pos]

	l.Tokens = append(l.Tokens, l.createToken(Number, string(value), l.Line, l.Col))
}

func (l *Lexer) processString() error {

	// Move past initial '"'.
	l.consume()

	start := l.Pos

	for !l.EOF() && !l.match('"') {
		l.consume()
	}

	if l.EOF() {
		return fmt.Errorf("incomplete string detected, line %v col %v", l.Line, l.Col)
	}

	value := l.Source[start:l.Pos]
	l.Tokens = append(l.Tokens, l.createToken(String, string(value), l.Line, l.Col))

	// Move past closing '"'.
	l.consume()

	return nil
}

/*
Does the target rune belong to the valid set of chars that form the start of an iden/keyword?
*/
func (l *Lexer) isIdentifierStart(ch rune) bool {

	return unicode.IsLetter(ch) || ch == '_'
}

/*
Does the target rune belong to the valid set of chars that form the rest of an iden/keyword?
*/
func (l *Lexer) isIdentifierPart(ch rune) bool {

	return unicode.IsLetter(ch) ||
		unicode.IsDigit(ch) ||
		ch == '_'
}

func (l *Lexer) processIdentifier() {

	start := l.Pos

	for !l.EOF() && l.isIdentifierPart(l.peek()) {
		l.consume()
	}

	value := l.Source[start:l.Pos]

	if tokenType, ok := Keywords[string(value)]; ok {

		// Keyword found.
		l.Tokens = append(l.Tokens, l.createToken(tokenType, string(value), l.Line, l.Col))
	} else {
		// Identifier found.
		l.Tokens = append(l.Tokens, l.createToken(Identifier, string(value), l.Line, l.Col))
	}

}

func (l *Lexer) processSymbol() error {

	for _, length := range availableTokenLengths {

		candidate, err := l.peekNextN(length)
		if err != nil {
			return err
		}

		tokenType, ok := tokenListByLength[length][string(candidate)]
		if !ok {
			continue
		}

		l.Tokens = append(l.Tokens, l.createToken(tokenType, string(candidate), l.Line, l.Pos))
		l.consumeN(length)
	}

	return nil
}

func Lex(source string) ([]Token, map[int]string, error) {

	lexer := Lexer{
		Source: []rune(source),

		Pos:  0,
		Line: 0,
		Col:  0,

		Audit:  make(map[int]string),
		Tokens: make([]Token, 0),
	}

	for !lexer.EOF() {

		switch ch := lexer.peek(); {

		case lexer.isSkippable(ch):
			lexer.consume()
			continue
		case lexer.match(';'):
			lexer.Tokens = append(lexer.Tokens, lexer.createToken(EOL, ";", lexer.Line, lexer.Col))
			lexer.consume()
		case unicode.IsDigit(ch):
			lexer.processNumber()
			continue
		case lexer.isIdentifierStart(ch):
			lexer.processIdentifier()
			continue
		case lexer.match('"'):
			err := lexer.processString()
			if err != nil {
				return nil, nil, err
			}
			continue
		default:
			lexer.processSymbol()
			continue
		}
	}

	// Add EOF to token list.
	lexer.Tokens = append(lexer.Tokens, lexer.createToken(EOF, "EOF", lexer.Line, lexer.Col))

	return lexer.Tokens, lexer.Audit, nil
}
