package query

import (
	"fmt"
)

//
// LexToken
//

type LexToken struct {
	id   LexTokenID
	val  string
	pos  int
	line int
}

func NewLexToken(id LexTokenID, val string, pos int, line int) *LexToken {
	return &LexToken{id, val, pos, line}
}

func (t *LexToken) ID() LexTokenID {
	return t.id
}

func (t *LexToken) Val() string {
	return t.val
}

func (t *LexToken) Pos() int {
	return t.pos
}

func (t *LexToken) Line() int {
	return t.line
}

var keywordMap = map[string]LexTokenID{
	"alter":    TokenALTER,
	"and":      TokenAND,
	"as":       TokenAS,
	"between":  TokenBETWEEN,
	"by":       TokenBY,
	"create":   TokenCREATE,
	"database": TokenDATABASE,
	"delete":   TokenDELETE,
	"distinct": TokenDISTINCT,
	"drop":     TokenDROP,
	"exists":   TokenEXISTS,
	"from":     TokenFROM,
	"group":    TokenGROUP,
	"having":   TokenHAVING,
	"if":       TokenIF,
	"in":       TokenIN,
	"index":    TokenINDEX,
	"inner":    TokenINNER,
	"insert":   TokenINSERT,
	"into":     TokenINTO,
	"join":     TokenJOIN,
	"left":     TokenLEFT,
	"like":     TokenLIKE,
	"limit":    TokenLIMIT,
	"not":      TokenNOT,
	"order":    TokenORDER,
	"pasword":  TokenPASSWORD,
	"right":    TokenRIGHT,
	"select":   TokenSELECT,
	"set":      TokenSET,
	"show":     TokenSHOW,
	"table":    TokenTABLE,
	"truncate": TokenTRUNCATE,
	"union":    TokenUNION,
	"update":   TokenUPDATE,
	"use":      TokenUSE,
	"user":     TokenUSER,
	"values":   TokenVALUES,
	"where":    TokenWHERE,

	// Type
	"string": TokenSTRING,
	"int":    TokenINT,
}

var symbolMap = map[string]LexTokenID{
	"@":  TokenAT,
	">=": TokenGEQ,
	"<=": TokenLEQ,
	"!=": TokenNEQ,
	"=":  TokenEQ,
	">":  TokenGT,
	"<":  TokenLT,
	"(":  TokenLPAREN,
	")":  TokenRPAREN,
	"[":  TokenLBRACK,
	"]":  TokenRBRACK,
	",":  TokenCOMMA,
	"+":  TokenPLUS,
	"-":  TokenMINUS,
	"*":  TokenTIMES,
	"/":  TokenDIV,
	"//": TokenDIVINT,
	"%":  TokenMODINT,
}

// Characters which are used as Symbol.
var symbolChars = []byte{
	'@', '>', '<', '!', '=', '>', '<', '(', ')',
	'[', ']', ',', '+', '-', '*', '/', '/', '%',
}

//
// Lexer
//

type Lexer struct {
	input  []rune
	carsor int         // Current index of input
	pos    int         // Current rune pointer
	line   int         // Current line pointer
	tokens []*LexToken // Result of lex process
}

func NewLexer(input string) *Lexer {
	return &Lexer{[]rune(input), 0, 0, 0, make([]*LexToken, 0)}
}

func (l *Lexer) Lex() ([]*LexToken, error) {
	for {
		if still := l.skipWhiteSpace(); !still {
			break
		}
		if err := l.lexToken(); err != nil {
			return nil, err
		}
	}
	return l.tokens, nil
}

func (l *Lexer) skipWhiteSpace() bool {

	for l.carsor < len(l.input) {

		r := l.input[l.carsor]

		switch r {
		case ' ':
			l.pos++
		case '\n':
			l.line++
			l.pos = 0
		default:
			return true
		}
		l.carsor++
	}

	return false
}

func (l *Lexer) lexToken() error {

	r := l.input[l.carsor]

	switch {
	case isString(r):
		return l.lexKey()

	case isNumber(r):
		return l.lexNumber()

	case isQuoto(r):
		return l.lexString()

	case isSymbolChar(r):
		return l.lexSymbol()

	default:
		return fmt.Errorf("lexToken: can't read token '%s' from %s", r, l.input)
	}
}

// lexXxx

func (l *Lexer) lexKey() error {

	start := l.carsor
	pos := l.pos
	line := l.line

	for l.carsor < len(l.input) {
		r := l.input[l.carsor]
		if !(isString(r) || isNumber(r)) {
			break
		}
		l.carsor++
		l.pos++
	}

	token := string(l.input[start:l.carsor])

	if isKeyword(token) {
		l.tokens = append(l.tokens, NewLexToken(keywordMap[token], token, pos, line))
		return nil
	}

	l.tokens = append(l.tokens, NewLexToken(TokenKEY, token, pos, line))
	return nil
}

func (l *Lexer) lexNumber() error {

	start := l.carsor
	pos := l.pos
	line := l.line

	for l.carsor < len(l.input) {
		r := l.input[l.carsor]
		if !isNumber(r) {
			break
		}
		l.carsor++
		l.pos++
	}

	token := string(l.input[start:l.carsor])

	l.tokens = append(l.tokens, NewLexToken(TokenVALUE, token, pos, line))
	return nil
}

func (l *Lexer) lexString() error {

	pos := l.pos
	line := l.line

	quota := l.input[l.carsor]
	l.carsor++
	l.pos++
	start := l.carsor
	skip := make([]int, 0)

	for l.carsor < len(l.input) {
		r := l.input[l.carsor]

		if r == '\\' {
			skip = append(skip, l.carsor)
			l.carsor += 2
			l.pos += 2
			continue
		}

		if r == quota {
			l.carsor++
			l.pos++
			break
		}

		if !(isString(r) || isNumber(r)) {
			break
		}

		l.carsor++
		l.pos++
	}

	skip = append(skip, l.carsor-1)
	tmp := make([]rune, 0)
	for _, i := range skip {
		tmp = append(tmp, l.input[start:i]...)
		start = i + 1
	}
	token := string(tmp)

	l.tokens = append(l.tokens, NewLexToken(TokenVALUE, token, pos, line))
	return nil
}

func (l *Lexer) lexSymbol() error {

	start := l.carsor
	pos := l.pos
	line := l.line

	for l.carsor < len(l.input) {

		r := l.input[l.carsor]

		if !isSymbolChar(r) {
			break
		}
		if !isSymbol(string(l.input[start : l.carsor+1])) {
			break
		}

		l.carsor++
		l.pos++
	}

	token := string(l.input[start:l.carsor])

	l.tokens = append(l.tokens, NewLexToken(symbolMap[token], token, pos, line))
	return nil
}

// isXxx(byte)

func isString(r rune) bool {
	return (rune('a') <= r && r <= rune('z')) || r == rune('_')
}

func isNumber(r rune) bool {
	return rune('0') <= r && r <= rune('9')
}

func isQuoto(r rune) bool {
	return rune('"') == r || r == rune('\'')
}

func isSymbolChar(r rune) bool {
	for _, s := range symbolChars {
		if r == rune(s) {
			return true
		}
	}
	return false
}

// isXxx(string)

func isKeyword(s string) bool {
	_, ok := keywordMap[s]
	return ok
}

func isSymbol(s string) bool {
	_, ok := symbolMap[s]
	return ok
}
