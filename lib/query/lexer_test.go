package query

import (
	//"fmt"
	"testing"

	"github.com/inazo1115/toydb/lib/util"
)

func TestSkipWhiteSpace0(t *testing.T) {
	lexer := NewLexer("   a")
	lexer.skipWhiteSpace()
	util.Assert(t, lexer.carsor, 3)
	util.Assert(t, lexer.pos, 3)
	util.Assert(t, lexer.line, 0)
}

func TestSkipWhiteSpace1(t *testing.T) {
	lexer := NewLexer("   \na")
	lexer.skipWhiteSpace()
	util.Assert(t, lexer.carsor, 4)
	util.Assert(t, lexer.pos, 0)
	util.Assert(t, lexer.line, 1)
}

func TestLexKey0(t *testing.T) {
	lexer := NewLexer("foo")
	lexer.lexKey()
	util.Assert(t, lexer.carsor, 3)
	util.Assert(t, lexer.pos, 3)
	util.Assert(t, lexer.line, 0)
	util.Assert(t, len(lexer.tokens), 1)
	util.Assert(t, lexer.tokens[0].ID, TokenKEY)
	util.Assert(t, lexer.tokens[0].Val, "foo")
}

func TestLexKey1(t *testing.T) {
	lexer := NewLexer("foo2000")
	lexer.lexKey()
	util.Assert(t, lexer.carsor, 7)
	util.Assert(t, lexer.pos, 7)
	util.Assert(t, lexer.line, 0)
	util.Assert(t, len(lexer.tokens), 1)
	util.Assert(t, lexer.tokens[0].ID, TokenKEY)
	util.Assert(t, lexer.tokens[0].Val, "foo2000")
}

func TestLexKey2(t *testing.T) {
	lexer := NewLexer("alter")
	lexer.lexKey()
	util.Assert(t, lexer.carsor, 5)
	util.Assert(t, lexer.pos, 5)
	util.Assert(t, lexer.line, 0)
	util.Assert(t, len(lexer.tokens), 1)
	util.Assert(t, lexer.tokens[0].ID, TokenALTER)
	util.Assert(t, lexer.tokens[0].Val, "alter")
}

func TestLexNumber(t *testing.T) {
	lexer := NewLexer("1000")
	lexer.lexNumber()
	util.Assert(t, lexer.carsor, 4)
	util.Assert(t, lexer.pos, 4)
	util.Assert(t, lexer.line, 0)
	util.Assert(t, len(lexer.tokens), 1)
	util.Assert(t, lexer.tokens[0].ID, TokenVALUE)
	util.Assert(t, lexer.tokens[0].Val, "1000")
}

func TestLexString0(t *testing.T) {
	lexer := NewLexer("'foo'")
	lexer.lexString()
	util.Assert(t, lexer.carsor, 5)
	util.Assert(t, lexer.pos, 5)
	util.Assert(t, lexer.line, 0)
	util.Assert(t, len(lexer.tokens), 1)
	util.Assert(t, lexer.tokens[0].ID, TokenVALUE)
	util.Assert(t, lexer.tokens[0].Val, "foo")
}

func TestLexString1(t *testing.T) {
	lexer := NewLexer("\"dosen\\'t\"")
	lexer.lexString()
	util.Assert(t, lexer.carsor, 10)
	util.Assert(t, lexer.pos, 10)
	util.Assert(t, lexer.line, 0)
	util.Assert(t, len(lexer.tokens), 1)
	util.Assert(t, lexer.tokens[0].ID, TokenVALUE)
	util.Assert(t, lexer.tokens[0].Val, "dosen't")
}

func TestLexString2(t *testing.T) {
	lexer := NewLexer("'foo2000foo'")
	lexer.lexString()
	util.Assert(t, lexer.carsor, 12)
	util.Assert(t, lexer.pos, 12)
	util.Assert(t, lexer.line, 0)
	util.Assert(t, len(lexer.tokens), 1)
	util.Assert(t, lexer.tokens[0].ID, TokenVALUE)
	util.Assert(t, lexer.tokens[0].Val, "foo2000foo")
}

func TestLexSymbol0(t *testing.T) {
	lexer := NewLexer("@")
	lexer.lexSymbol()
	util.Assert(t, lexer.carsor, 1)
	util.Assert(t, lexer.pos, 1)
	util.Assert(t, lexer.line, 0)
	util.Assert(t, len(lexer.tokens), 1)
	util.Assert(t, lexer.tokens[0].ID, TokenAT)
	util.Assert(t, lexer.tokens[0].Val, "@")
}

func TestLexSymbol1(t *testing.T) {
	lexer := NewLexer(">=")
	lexer.lexSymbol()
	util.Assert(t, lexer.carsor, 2)
	util.Assert(t, lexer.pos, 2)
	util.Assert(t, lexer.line, 0)
	util.Assert(t, len(lexer.tokens), 1)
	util.Assert(t, lexer.tokens[0].ID, TokenGEQ)
	util.Assert(t, lexer.tokens[0].Val, ">=")
}

func TestLex0(t *testing.T) {
	input := "create table table_name (name string(20), age int, tel int)"
	expected := []*LexToken{
		&LexToken{TokenCREATE, "create"},
		&LexToken{TokenTABLE, "table"},
		&LexToken{TokenKEY, "table_name"},
		&LexToken{TokenLPAREN, "("},
		&LexToken{TokenKEY, "name"},
		&LexToken{TokenSTRING, "string"},
		&LexToken{TokenLPAREN, "("},
		&LexToken{TokenVALUE, "20"},
		&LexToken{TokenRPAREN, ")"},
		&LexToken{TokenCOMMA, ","},
		&LexToken{TokenKEY, "age"},
		&LexToken{TokenINT, "int"},
		&LexToken{TokenCOMMA, ","},
		&LexToken{TokenKEY, "tel"},
		&LexToken{TokenINT, "int"},
		&LexToken{TokenRPAREN, ")"},
	}
	actual, err := NewLexer(input).Lex()
	if err != nil {
		t.Errorf("%v", err)
	}

	// Assert.
	util.Assert(t, len(actual), len(expected))
	for i := 0; i < len(actual); i++ {
		util.Assert(t, actual[i].ID, expected[i].ID)
		util.Assert(t, actual[i].Val, expected[i].Val)
	}
}

func TestLex1(t *testing.T) {
	input := "insert into table_name (name, age, tel) values (\"foofoo\", 100, 200)"
	expected := []*LexToken{
		&LexToken{TokenINSERT, "insert"},
		&LexToken{TokenINTO, "into"},
		&LexToken{TokenKEY, "table_name"},
		&LexToken{TokenLPAREN, "("},
		&LexToken{TokenKEY, "name"},
		&LexToken{TokenCOMMA, ","},
		&LexToken{TokenKEY, "age"},
		&LexToken{TokenCOMMA, ","},
		&LexToken{TokenKEY, "tel"},
		&LexToken{TokenRPAREN, ")"},
		&LexToken{TokenVALUES, "values"},
		&LexToken{TokenLPAREN, "("},
		&LexToken{TokenVALUE, "foofoo"},
		&LexToken{TokenCOMMA, ","},
		&LexToken{TokenVALUE, "100"},
		&LexToken{TokenCOMMA, ","},
		&LexToken{TokenVALUE, "200"},
		&LexToken{TokenRPAREN, ")"},
	}
	actual, err := NewLexer(input).Lex()
	if err != nil {
		t.Errorf("%v", err)
	}

	// Assert.
	util.Assert(t, len(actual), len(expected))
	for i := 0; i < len(actual); i++ {
		util.Assert(t, actual[i].ID, expected[i].ID)
		util.Assert(t, actual[i].Val, expected[i].Val)
	}
}

func TestLex2(t *testing.T) {
	input := "select * from table_name"
	expected := []*LexToken{
		&LexToken{TokenSELECT, "select"},
		&LexToken{TokenTIMES, "*"},
		&LexToken{TokenFROM, "from"},
		&LexToken{TokenKEY, "table_name"},
	}
	actual, err := NewLexer(input).Lex()
	if err != nil {
		t.Errorf("%v", err)
	}

	// Assert.
	util.Assert(t, len(actual), len(expected))
	for i := 0; i < len(actual); i++ {
		util.Assert(t, actual[i].ID, expected[i].ID)
		util.Assert(t, actual[i].Val, expected[i].Val)
	}
}
