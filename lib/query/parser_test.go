package query

import (
	"testing"

	"github.com/inazo1115/toydb/lib/util"
)

func TestParse0(t *testing.T) {
	input := []*LexToken{
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

	actual, err := NewParser().Parse(input)
	if err != nil {
		t.Errorf("%v", err)
	}

	util.Assert(t, actual.Token.ID, TokenCREATE)
	util.Assert(t, actual.Token.Val, "create")
	util.Assert(t, actual.Children[0].Token.ID, TokenKEY)
	util.Assert(t, actual.Children[0].Token.Val, "table_name")
	util.Assert(t, actual.Children[1].Token.ID, TokenKEY)
	util.Assert(t, actual.Children[1].Token.Val, "name")
	util.Assert(t, actual.Children[2].Token.ID, TokenSTRING)
	util.Assert(t, actual.Children[2].Token.Val, "string")
	util.Assert(t, actual.Children[2].Children[0].Token.ID, TokenVALUE)
	util.Assert(t, actual.Children[2].Children[0].Token.Val, "20")
	util.Assert(t, actual.Children[3].Token.ID, TokenKEY)
	util.Assert(t, actual.Children[3].Token.Val, "age")
	util.Assert(t, actual.Children[4].Token.ID, TokenINT)
	util.Assert(t, actual.Children[4].Token.Val, "int")
	util.Assert(t, actual.Children[5].Token.ID, TokenKEY)
	util.Assert(t, actual.Children[5].Token.Val, "tel")
	util.Assert(t, actual.Children[6].Token.ID, TokenINT)
	util.Assert(t, actual.Children[6].Token.Val, "int")
}

func TestParse1(t *testing.T) {
	input := []*LexToken{
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
		&LexToken{TokenVALUE, "\"foofoo\""},
		&LexToken{TokenCOMMA, ","},
		&LexToken{TokenVALUE, "100"},
		&LexToken{TokenCOMMA, ","},
		&LexToken{TokenVALUE, "200"},
		&LexToken{TokenRPAREN, ")"},
	}

	actual, err := NewParser().Parse(input)
	if err != nil {
		t.Errorf("%v", err)
	}

	util.Assert(t, actual.Token.ID, TokenINSERT)
	util.Assert(t, actual.Token.Val, "insert")
	util.Assert(t, actual.Children[0].Token.ID, TokenKEY)
	util.Assert(t, actual.Children[0].Token.Val, "table_name")
	util.Assert(t, actual.Children[1].Token.ID, TokenKEY)
	util.Assert(t, actual.Children[1].Token.Val, "name")
	util.Assert(t, actual.Children[1].Children[0].Token.ID, TokenVALUE)
	util.Assert(t, actual.Children[1].Children[0].Token.Val, "\"foofoo\"")
	util.Assert(t, actual.Children[2].Token.ID, TokenKEY)
	util.Assert(t, actual.Children[2].Token.Val, "age")
	util.Assert(t, actual.Children[2].Children[0].Token.ID, TokenVALUE)
	util.Assert(t, actual.Children[2].Children[0].Token.Val, "100")
	util.Assert(t, actual.Children[3].Token.ID, TokenKEY)
	util.Assert(t, actual.Children[3].Token.Val, "tel")
	util.Assert(t, actual.Children[3].Children[0].Token.ID, TokenVALUE)
	util.Assert(t, actual.Children[3].Children[0].Token.Val, "200")
}

func TestParse2(t *testing.T) {
	input := []*LexToken{
		&LexToken{TokenSELECT, "select"},
		&LexToken{TokenTIMES, "*"},
		&LexToken{TokenFROM, "from"},
		&LexToken{TokenKEY, "table_name"},
	}

	actual, err := NewParser().Parse(input)
	if err != nil {
		t.Errorf("%v", err)
	}

	util.Assert(t, actual.Token.ID, TokenSELECT)
	util.Assert(t, actual.Token.Val, "select")
	util.Assert(t, actual.Children[0].Token.ID, TokenKEY)
	util.Assert(t, actual.Children[0].Token.Val, "table_name")
	util.Assert(t, actual.Children[1].Token.ID, TokenTIMES)
	util.Assert(t, actual.Children[1].Token.Val, "*")
}
