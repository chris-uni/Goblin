package lexer

import (
	"fmt"
	"unicode"
)

const lexerErrorOne = "lexer error on"
const lexerErrorTwo = "message "
const lexerErrorThree = "code "

type LexerError struct {
	Code    int
	Message string
	Line    int
	Col     int
}

func (e *LexerError) Error() string {
	return fmt.Sprintf("%v line %v col %v\n%v%v\n%v %v\n", lexerErrorOne, e.Line, e.Col, lexerErrorTwo, e.Message, lexerErrorThree, e.Code)
}

/*
Experminental. Formats a length dependant line and arrow, pointing to specific part of error in source.
*/
func ErrorUnderlineFormatter(message string, subject string) string {

	underline := ""

	lenPrefixAndMessage := len(lexerErrorTwo) + len(message)

	for i := 0; i < lenPrefixAndMessage; i++ {
		underline += " "
	}

	for i := 0; i < len(subject); i++ {
		underline += "~"
	}

	underline += "^"

	return underline
}

var lexerErrorCodes = map[int]string{

	// Lexer struct codes.
	100: "peekNext cannot overflow past end of file",
	101: "peekNextN cannot overflow past end of file",

	// Unexpected character/rune codes
	200: "unsupported rune found",

	// Unterminated string codes
	300: "unterminated string detected",

	// Invalid number codes
	// 400

	// Invalid Identifier codes
	// 500

	// Malformed multi-character token codes
	// 600
}

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

		return 0, &LexerError{100, lexerErrorCodes[100], l.Line, l.Col}
	}

	return l.Source[newPos], nil
}

/*
Non-consuming. Returns the next n'th rune (if available) in the Source list.
*/
func (l *Lexer) peekNextN(n int) (string, error) {

	newPos := l.Pos + n

	if newPos > len(l.Source) {

		return "", &LexerError{101, lexerErrorCodes[101], l.Line, l.Col}
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
Non-comsuming. Returns true if any of the argument values match the value of __peek__.
*/
func (l *Lexer) matchAny(chs ...rune) bool {
	for _, ch := range chs {

		if l.peek() == ch {
			return true
		}
	}

	return false
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

		subject := string(l.Source[start:l.Pos])
		return &LexerError{300, fmt.Sprintf("%v \"%v\n%v", lexerErrorCodes[300], subject, ErrorUnderlineFormatter(lexerErrorCodes[300], subject)), l.Line, l.Col}
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

	var candidate string
	var err error

	for _, length := range availableTokenLengths {

		// Prevents the lexer looking for tokens larger than the source of the file.
		if length > len(l.Source) {
			continue
		}

		candidate, err = l.peekNextN(length)
		if err != nil {
			return err
		}

		tokenType, ok := tokenListByLength[length][string(candidate)]
		if ok {
			l.Tokens = append(l.Tokens, l.createToken(tokenType, string(candidate), l.Line, l.Pos))
			l.consumeN(length)
			return nil
		}
	}

	subject := string(candidate)
	return &LexerError{200, fmt.Sprintf("%v %v\n%v", lexerErrorCodes[200], subject, ErrorUnderlineFormatter(lexerErrorCodes[200], subject)), l.Line, l.Col}
}

func (l *Lexer) processComment() error {

	nextRune, err := l.peekNext()
	if err != nil {
		return err
	}

	var err_ error

	switch nextRune {
	case '/':
		err_ = l.processSingleLineComment()
	case '*':
		err_ = l.processMultiLineComment()
	}

	return err_
}

func (l *Lexer) processSingleLineComment() error {

	nextRune, err := l.peekNext()
	if err != nil {
		return err
	}

	if l.peek() == '/' && nextRune == '/' {

		// Loop until EOF or a new line starts.
		for !l.EOF() && !l.match('\n') {
			l.consume()
		}
	}

	return nil
}

func (l *Lexer) processMultiLineComment() error {

	nextRune, err := l.peekNext()
	if err != nil {
		return err
	}

	if l.peek() == '/' && nextRune == '*' {

		// Move past '/'
		l.consume()

		// Move past '*'
		l.consume()

		for !l.EOF() {

			nextPeek, err := l.peekNext()
			if err != nil {
				return err
			}

			if l.peek() == '*' && nextPeek == '/' {
				l.consume()
				l.consume()
				break
			}

			l.consume()
		}

	}

	return nil
}

func Lex(source string) ([]Token, map[int]string, error) {

	lexer := Lexer{
		Source: []rune(source),

		Pos:  0,
		Line: 1,
		Col:  0,

		Audit:  make(map[int]string),
		Tokens: make([]Token, 0),
	}

	for !lexer.EOF() {

		switch ch := lexer.peek(); {

		case lexer.matchAny(' ', '\n', '\t', '\r'):
			lexer.consume()
			continue
		case lexer.match(';'):
			lexer.Tokens = append(lexer.Tokens, lexer.createToken(EOL, ";", lexer.Line, lexer.Col))
			lexer.consume()
			continue
		case lexer.match('/'):
			err := lexer.processComment()
			if err != nil {
				return nil, nil, err
			}
			continue
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
			err := lexer.processSymbol()
			if err != nil {
				return nil, nil, err
			}
			continue
		}
	}

	// Add EOF to token list.
	lexer.Tokens = append(lexer.Tokens, lexer.createToken(EOF, "EOF", lexer.Line, lexer.Col))

	return lexer.Tokens, lexer.Audit, nil
}
