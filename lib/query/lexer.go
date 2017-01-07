package query

import (
	"bytes"
	"fmt"
	//"unicode/utf8"
)

//
// LexToken
//

type LexToken struct {
	ID  LexTokenID
	Val string
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

func Lex(query string) ([]LexToken, error) {

	var token LexToken
	var err error
	rest := query
	ret := make([]LexToken, 0)

	for {
		rest = skipWhiteSpace(rest)
		if len(rest) == 0 {
			break
		}

		token, err = lexToken(rest)
		if err != nil {
			return []LexToken{}, err
		}

		rest = rest[len(token.Val):]
		ret = append(ret, token)
	}

	return ret, nil
}

func lexToken(s string) (LexToken, error) {

	switch {
	case isString(s[0]):
		return lexKey(s), nil

	case isNumber(s[0]):
		return lexNumber(s), nil

	case isQuoto(s[0]):
		return lexString(s), nil

	case isSymbolChar(s[0]):
		return lexSymbol(s), nil

	default:
		return LexToken{}, fmt.Errorf("lexToken: can't read token from %s, %s", s, s[0])
	}
}

// isXxx(byte)

func isString(b byte) bool {
	return ('a' <= b && b <= 'z') || b == '_'
}

func isNumber(b byte) bool {
	return '0' <= b && b <= '9'
}

func isQuoto(b byte) bool {
	return '"' == b || b == '\''
}

func isSymbolChar(b byte) bool {
	for _, s := range symbolChars {
		if b == s {
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

// lexXxx

func lexKey(s string) LexToken {

	buf := make([]byte, 0)
	for _, r := range s {
		b := byte(r)
		if !isString(b) {
			break
		}
		buf = append(buf, b)
	}

	token := string(buf)

	if isKeyword(token) {
		return LexToken{keywordMap[token], token}
	}
	return LexToken{TokenKEY, token}
}

func lexNumber(s string) LexToken {

	buf := make([]byte, 0)
	for _, r := range s {
		b := byte(r)
		if !isNumber(b) {
			break
		}
		buf = append(buf, b)
	}

	token := string(buf)

	return LexToken{TokenVALUE, token}
}

func lexString(s string) LexToken {

	quota := s[0]
	buf := make([]byte, 1)
	buf = append(buf, quota)

	for i := 1; i < len(s); i++ {

		b := s[i]

		if b == quota {
			buf = append(buf, b)
			break
		}

		if b == '\\' {
			buf = append(buf, b)
			buf = append(buf, s[i+1])
			i++
			break
		}

		if !isString(b) {
			break
		}

		buf = append(buf, b)
	}

	// I don't know why this code is required.
	buf = bytes.Trim(buf, string(0))

	token := string(buf)

	return LexToken{TokenVALUE, token}
}

func lexSymbol(s string) LexToken {

	buf := make([]byte, 0)

	for _, r := range s {
		b := byte(r)
		if !isSymbolChar(b) {
			break
		}
		if !isSymbol(string(append(buf, b))) {
			break
		}
		buf = append(buf, b)
	}

	token := string(buf)

	return LexToken{symbolMap[token], token}
}

// Unitily functions

func skipWhiteSpace(s string) string {

	if len(s) == 0 {
		return ""
	}

	if s[0] == ' ' {
		return skipWhiteSpace(s[1:])
	}

	return s
}
