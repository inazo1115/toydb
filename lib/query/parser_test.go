package query

import (
	//"fmt"
	"testing"
)

func assert(t *testing.T, actual interface{}, expected interface{}) {
	if actual != expected {
		t.Errorf("actual: %v doesn't equals expected: %v.", actual, expected)
	}
}

func TestParse0(t *testing.T) {

	input := []LexToken{
		LexToken{TokenCREATE, "create"},
		LexToken{TokenTABLE, "table"},
		LexToken{TokenKEY, "table_name"},
		LexToken{TokenLPAREN, "("},
		LexToken{TokenKEY, "name"},
		LexToken{TokenSTRING, "string"},
		LexToken{TokenLPAREN, "("},
		LexToken{TokenVALUE, "20"},
		LexToken{TokenRPAREN, ")"},
		LexToken{TokenCOMMA, ","},
		LexToken{TokenKEY, "age"},
		LexToken{TokenINT, "int"},
		LexToken{TokenCOMMA, ","},
		LexToken{TokenKEY, "tel"},
		LexToken{TokenINT, "int"},
		LexToken{TokenRPAREN, ")"},
	}

	actual, err := Parse(input)
	if err != nil {
		t.Errorf("%v", err)
	}

	assert(t, actual.Token.ID, TokenCREATE)
	assert(t, actual.Token.Val, "create")
	assert(t, actual.Children[0].Token.ID, TokenKEY)
	assert(t, actual.Children[0].Token.Val, "table_name")
	assert(t, actual.Children[1].Token.ID, TokenKEY)
	assert(t, actual.Children[1].Token.Val, "name")
	assert(t, actual.Children[2].Token.ID, TokenSTRING)
	assert(t, actual.Children[2].Token.Val, "string")
	assert(t, actual.Children[2].Children[0].Token.ID, TokenVALUE)
	assert(t, actual.Children[2].Children[0].Token.Val, "20")
	assert(t, actual.Children[3].Token.ID, TokenKEY)
	assert(t, actual.Children[3].Token.Val, "age")
	assert(t, actual.Children[4].Token.ID, TokenINT)
	assert(t, actual.Children[4].Token.Val, "int")
	assert(t, actual.Children[5].Token.ID, TokenKEY)
	assert(t, actual.Children[5].Token.Val, "tel")
	assert(t, actual.Children[6].Token.ID, TokenINT)
	assert(t, actual.Children[6].Token.Val, "int")
}

func TestParse1(t *testing.T) {

	input := []LexToken{
		LexToken{TokenINSERT, "insert"},
		LexToken{TokenINTO, "into"},
		LexToken{TokenKEY, "table_name"},
		LexToken{TokenLPAREN, "("},
		LexToken{TokenKEY, "name"},
		LexToken{TokenCOMMA, ","},
		LexToken{TokenKEY, "age"},
		LexToken{TokenCOMMA, ","},
		LexToken{TokenKEY, "tel"},
		LexToken{TokenRPAREN, ")"},
		LexToken{TokenVALUES, "values"},
		LexToken{TokenLPAREN, "("},
		LexToken{TokenVALUE, "\"foofoo\""},
		LexToken{TokenCOMMA, ","},
		LexToken{TokenVALUE, "100"},
		LexToken{TokenCOMMA, ","},
		LexToken{TokenVALUE, "200"},
		LexToken{TokenRPAREN, ")"},
	}

	actual, err := Parse(input)
	if err != nil {
		t.Errorf("%v", err)
	}

	assert(t, actual.Token.ID, TokenINSERT)
	assert(t, actual.Token.Val, "insert")
	assert(t, actual.Children[0].Token.ID, TokenKEY)
	assert(t, actual.Children[0].Token.Val, "table_name")
	assert(t, actual.Children[1].Token.ID, TokenKEY)
	assert(t, actual.Children[1].Token.Val, "name")
	assert(t, actual.Children[1].Children[0].Token.ID, TokenVALUE)
	assert(t, actual.Children[1].Children[0].Token.Val, "\"foofoo\"")
	assert(t, actual.Children[2].Token.ID, TokenKEY)
	assert(t, actual.Children[2].Token.Val, "age")
	assert(t, actual.Children[2].Children[0].Token.ID, TokenVALUE)
	assert(t, actual.Children[2].Children[0].Token.Val, "100")
	assert(t, actual.Children[3].Token.ID, TokenKEY)
	assert(t, actual.Children[3].Token.Val, "tel")
	assert(t, actual.Children[3].Children[0].Token.ID, TokenVALUE)
	assert(t, actual.Children[3].Children[0].Token.Val, "200")
}

func TestParse2(t *testing.T) {

	input := []LexToken{
		LexToken{TokenSELECT, "select"},
		LexToken{TokenTIMES, "*"},
		LexToken{TokenFROM, "from"},
		LexToken{TokenKEY, "table_name"},
	}

	actual, err := Parse(input)
	if err != nil {
		t.Errorf("%v", err)
	}

	assert(t, actual.Token.ID, TokenSELECT)
	assert(t, actual.Token.Val, "select")
	assert(t, actual.Children[0].Token.ID, TokenKEY)
	assert(t, actual.Children[0].Token.Val, "table_name")
	assert(t, actual.Children[1].Token.ID, TokenTIMES)
	assert(t, actual.Children[1].Token.Val, "*")
}
