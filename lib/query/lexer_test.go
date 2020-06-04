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
	util.Assert(t, lexer.tokens[0].ID(), TokenKEY)
	util.Assert(t, lexer.tokens[0].Val(), "foo")
}

func TestLexKey1(t *testing.T) {
	lexer := NewLexer("foo2000")
	lexer.lexKey()
	util.Assert(t, lexer.carsor, 7)
	util.Assert(t, lexer.pos, 7)
	util.Assert(t, lexer.line, 0)
	util.Assert(t, len(lexer.tokens), 1)
	util.Assert(t, lexer.tokens[0].ID(), TokenKEY)
	util.Assert(t, lexer.tokens[0].Val(), "foo2000")
}

func TestLexKey2(t *testing.T) {
	lexer := NewLexer("alter")
	lexer.lexKey()
	util.Assert(t, lexer.carsor, 5)
	util.Assert(t, lexer.pos, 5)
	util.Assert(t, lexer.line, 0)
	util.Assert(t, len(lexer.tokens), 1)
	util.Assert(t, lexer.tokens[0].ID(), TokenALTER)
	util.Assert(t, lexer.tokens[0].Val(), "alter")
}

func TestLexNumber(t *testing.T) {
	lexer := NewLexer("1000")
	lexer.lexNumber()
	util.Assert(t, lexer.carsor, 4)
	util.Assert(t, lexer.pos, 4)
	util.Assert(t, lexer.line, 0)
	util.Assert(t, len(lexer.tokens), 1)
	util.Assert(t, lexer.tokens[0].ID(), TokenVALUE)
	util.Assert(t, lexer.tokens[0].Val(), "1000")
}

func TestLexString0(t *testing.T) {
	lexer := NewLexer("'foo'")
	lexer.lexString()
	util.Assert(t, lexer.carsor, 5)
	util.Assert(t, lexer.pos, 5)
	util.Assert(t, lexer.line, 0)
	util.Assert(t, len(lexer.tokens), 1)
	util.Assert(t, lexer.tokens[0].ID(), TokenVALUE)
	util.Assert(t, lexer.tokens[0].Val(), "foo")
}

func TestLexString1(t *testing.T) {
	lexer := NewLexer("\"dosen\\'t\"")
	lexer.lexString()
	util.Assert(t, lexer.carsor, 10)
	util.Assert(t, lexer.pos, 10)
	util.Assert(t, lexer.line, 0)
	util.Assert(t, len(lexer.tokens), 1)
	util.Assert(t, lexer.tokens[0].ID(), TokenVALUE)
	util.Assert(t, lexer.tokens[0].Val(), "dosen't")
}

func TestLexString2(t *testing.T) {
	lexer := NewLexer("'foo2000foo'")
	lexer.lexString()
	util.Assert(t, lexer.carsor, 12)
	util.Assert(t, lexer.pos, 12)
	util.Assert(t, lexer.line, 0)
	util.Assert(t, len(lexer.tokens), 1)
	util.Assert(t, lexer.tokens[0].ID(), TokenVALUE)
	util.Assert(t, lexer.tokens[0].Val(), "foo2000foo")
}

func TestLexSymbol0(t *testing.T) {
	lexer := NewLexer("@")
	lexer.lexSymbol()
	util.Assert(t, lexer.carsor, 1)
	util.Assert(t, lexer.pos, 1)
	util.Assert(t, lexer.line, 0)
	util.Assert(t, len(lexer.tokens), 1)
	util.Assert(t, lexer.tokens[0].ID(), TokenAT)
	util.Assert(t, lexer.tokens[0].Val(), "@")
}

func TestLexSymbol1(t *testing.T) {
	lexer := NewLexer(">=")
	lexer.lexSymbol()
	util.Assert(t, lexer.carsor, 2)
	util.Assert(t, lexer.pos, 2)
	util.Assert(t, lexer.line, 0)
	util.Assert(t, len(lexer.tokens), 1)
	util.Assert(t, lexer.tokens[0].ID(), TokenGEQ)
	util.Assert(t, lexer.tokens[0].Val(), ">=")
}

func TestLex0(t *testing.T) {
	input := "create table table_name (name string(20), age int, tel int)"
	expected := []*LexToken{
		&LexToken{TokenCREATE, "create", 0, 0},
		&LexToken{TokenTABLE, "table", 7, 0},
		&LexToken{TokenKEY, "table_name", 13, 0},
		&LexToken{TokenLPAREN, "(", 24, 0},
		&LexToken{TokenKEY, "name", 25, 0},
		&LexToken{TokenSTRING, "string", 30, 0},
		&LexToken{TokenLPAREN, "(", 36, 0},
		&LexToken{TokenVALUE, "20", 37, 0},
		&LexToken{TokenRPAREN, ")", 39, 0},
		&LexToken{TokenCOMMA, ",", 40, 0},
		&LexToken{TokenKEY, "age", 42, 0},
		&LexToken{TokenINT, "int", 46, 0},
		&LexToken{TokenCOMMA, ",", 49, 0},
		&LexToken{TokenKEY, "tel", 51, 0},
		&LexToken{TokenINT, "int", 55, 0},
		&LexToken{TokenRPAREN, ")", 58, 0},
	}
	actual, err := NewLexer(input).Lex()
	if err != nil {
		t.Errorf("%v", err)
	}

	// Assert.
	util.Assert(t, len(actual), len(expected))
	for i := 0; i < len(actual); i++ {
		util.Assert(t, actual[i].ID(), expected[i].ID())
		util.Assert(t, actual[i].Val(), expected[i].Val())
		util.Assert(t, actual[i].Pos(), expected[i].Pos())
		util.Assert(t, actual[i].Line(), expected[i].Line())
	}
}

func TestLex1(t *testing.T) {
	input := "insert into table_name (name, age, tel) values (\"foofoo\", 100, 200)"
	expected := []*LexToken{
		&LexToken{TokenINSERT, "insert", 0, 0},
		&LexToken{TokenINTO, "into", 7, 0},
		&LexToken{TokenKEY, "table_name", 12, 0},
		&LexToken{TokenLPAREN, "(", 23, 0},
		&LexToken{TokenKEY, "name", 24, 0},
		&LexToken{TokenCOMMA, ",", 28, 0},
		&LexToken{TokenKEY, "age", 30, 0},
		&LexToken{TokenCOMMA, ",", 33, 0},
		&LexToken{TokenKEY, "tel", 35, 0},
		&LexToken{TokenRPAREN, ")", 38, 0},
		&LexToken{TokenVALUES, "values", 40, 0},
		&LexToken{TokenLPAREN, "(", 47, 0},
		&LexToken{TokenVALUE, "foofoo", 48, 0},
		&LexToken{TokenCOMMA, ",", 56, 0},
		&LexToken{TokenVALUE, "100", 58, 0},
		&LexToken{TokenCOMMA, ",", 61, 0},
		&LexToken{TokenVALUE, "200", 63, 0},
		&LexToken{TokenRPAREN, ")", 66, 0},
	}
	actual, err := NewLexer(input).Lex()
	if err != nil {
		t.Errorf("%v", err)
	}

	// Assert.
	util.Assert(t, len(actual), len(expected))
	for i := 0; i < len(actual); i++ {
		util.Assert(t, actual[i].ID(), expected[i].ID())
		util.Assert(t, actual[i].Val(), expected[i].Val())
		util.Assert(t, actual[i].Pos(), expected[i].Pos())
		util.Assert(t, actual[i].Line(), expected[i].Line())
	}
}

func TestLex2(t *testing.T) {
	input := "select * from table_name"
	expected := []*LexToken{
		&LexToken{TokenSELECT, "select", 0, 0},
		&LexToken{TokenTIMES, "*", 7, 0},
		&LexToken{TokenFROM, "from", 9, 0},
		&LexToken{TokenKEY, "table_name", 14, 0},
	}
	actual, err := NewLexer(input).Lex()
	if err != nil {
		t.Errorf("%v", err)
	}

	// Assert.
	util.Assert(t, len(actual), len(expected))
	for i := 0; i < len(actual); i++ {
		util.Assert(t, actual[i].ID(), expected[i].ID())
		util.Assert(t, actual[i].Val(), expected[i].Val())
		util.Assert(t, actual[i].Pos(), expected[i].Pos())
		util.Assert(t, actual[i].Line(), expected[i].Line())
	}
}
