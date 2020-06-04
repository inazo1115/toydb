package query

import (
	"testing"

	"github.com/inazo1115/toydb/lib/util"
)

func TestParse0(t *testing.T) {
	input := []*LexToken{
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

	actual, err := NewParser().Parse(input)
	if err != nil {
		t.Errorf("%v", err)
	}

	util.Assert(t, actual.Token.ID(), TokenCREATE)
	util.Assert(t, actual.Token.Val(), "create")
	util.Assert(t, actual.Children[0].Token.ID(), TokenKEY)
	util.Assert(t, actual.Children[0].Token.Val(), "table_name")
	util.Assert(t, actual.Children[1].Token.ID(), TokenKEY)
	util.Assert(t, actual.Children[1].Token.Val(), "name")
	util.Assert(t, actual.Children[2].Token.ID(), TokenSTRING)
	util.Assert(t, actual.Children[2].Token.Val(), "string")
	util.Assert(t, actual.Children[2].Children[0].Token.ID(), TokenVALUE)
	util.Assert(t, actual.Children[2].Children[0].Token.Val(), "20")
	util.Assert(t, actual.Children[3].Token.ID(), TokenKEY)
	util.Assert(t, actual.Children[3].Token.Val(), "age")
	util.Assert(t, actual.Children[4].Token.ID(), TokenINT)
	util.Assert(t, actual.Children[4].Token.Val(), "int")
	util.Assert(t, actual.Children[5].Token.ID(), TokenKEY)
	util.Assert(t, actual.Children[5].Token.Val(), "tel")
	util.Assert(t, actual.Children[6].Token.ID(), TokenINT)
	util.Assert(t, actual.Children[6].Token.Val(), "int")
}

func TestParse1(t *testing.T) {
	input := []*LexToken{
		&LexToken{TokenINSERT, "insert", 0, 0},
		&LexToken{TokenINTO, "into", 8, 0},
		&LexToken{TokenKEY, "table_name", 13, 0},
		&LexToken{TokenLPAREN, "(", 24, 0},
		&LexToken{TokenKEY, "name", 25, 0},
		&LexToken{TokenCOMMA, ",", 29, 0},
		&LexToken{TokenKEY, "age", 31, 0},
		&LexToken{TokenCOMMA, ",", 34, 0},
		&LexToken{TokenKEY, "tel", 36, 0},
		&LexToken{TokenRPAREN, ")", 39, 0},
		&LexToken{TokenVALUES, "values", 41, 0},
		&LexToken{TokenLPAREN, "(", 48, 0},
		&LexToken{TokenVALUE, "foofoo", 49, 0},
		&LexToken{TokenCOMMA, ",", 57, 0},
		&LexToken{TokenVALUE, "100", 59, 0},
		&LexToken{TokenCOMMA, ",", 62, 0},
		&LexToken{TokenVALUE, "200", 64, 0},
		&LexToken{TokenRPAREN, ")", 67, 0},
	}

	actual, err := NewParser().Parse(input)
	if err != nil {
		t.Errorf("%v", err)
	}

	util.Assert(t, actual.Token.ID(), TokenINSERT)
	util.Assert(t, actual.Token.Val(), "insert")
	util.Assert(t, actual.Children[0].Token.ID(), TokenKEY)
	util.Assert(t, actual.Children[0].Token.Val(), "table_name")
	util.Assert(t, actual.Children[1].Token.ID(), TokenKEY)
	util.Assert(t, actual.Children[1].Token.Val(), "name")
	util.Assert(t, actual.Children[1].Children[0].Token.ID(), TokenVALUE)
	util.Assert(t, actual.Children[1].Children[0].Token.Val(), "foofoo")
	util.Assert(t, actual.Children[2].Token.ID(), TokenKEY)
	util.Assert(t, actual.Children[2].Token.Val(), "age")
	util.Assert(t, actual.Children[2].Children[0].Token.ID(), TokenVALUE)
	util.Assert(t, actual.Children[2].Children[0].Token.Val(), "100")
	util.Assert(t, actual.Children[3].Token.ID(), TokenKEY)
	util.Assert(t, actual.Children[3].Token.Val(), "tel")
	util.Assert(t, actual.Children[3].Children[0].Token.ID(), TokenVALUE)
	util.Assert(t, actual.Children[3].Children[0].Token.Val(), "200")
}

func TestParse2(t *testing.T) {
	input := []*LexToken{
		&LexToken{TokenSELECT, "select", 0, 0},
		&LexToken{TokenTIMES, "*", 7, 0},
		&LexToken{TokenFROM, "from", 9, 0},
		&LexToken{TokenKEY, "table_name", 14, 0},
	}

	actual, err := NewParser().Parse(input)
	if err != nil {
		t.Errorf("%v", err)
	}

	util.Assert(t, actual.Token.ID(), TokenSELECT)
	util.Assert(t, actual.Token.Val(), "select")
	util.Assert(t, actual.Children[0].Token.ID(), TokenKEY)
	util.Assert(t, actual.Children[0].Token.Val(), "table_name")
	util.Assert(t, actual.Children[1].Token.ID(), TokenTIMES)
	util.Assert(t, actual.Children[1].Token.Val(), "*")
}
