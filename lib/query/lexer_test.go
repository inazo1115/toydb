package query

import (
	//"fmt"
	"testing"
)

func TestLex0(t *testing.T) {
	input := "create table table_name (name string(20), age int, tel int)"
	expected := []LexToken{
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
	actual, err := NewLexer().Lex(input)
	if err != nil {
		t.Errorf("%v", err)
	}

	// Assert.
	if len(actual) != len(expected) {
		t.Errorf("actual: %v doesn't equals expected: %v.", actual, expected)
	}
	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("actual: %v doesn't equals expected: %v.", actual, expected)
		}
	}
}

func TestLex1(t *testing.T) {
	input := "insert into table_name (name, age, tel) values (\"foofoo\", 100, 200)"
	expected := []LexToken{
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
	actual, err := NewLexer().Lex(input)
	if err != nil {
		t.Errorf("%v", err)
	}

	// Assert.
	if len(actual) != len(expected) {
		t.Errorf("actual: %v doesn't equals expected: %v.", actual, expected)
	}
	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("actual: %v doesn't equals expected: %v.", actual, expected)
		}
	}
}

func TestLex2(t *testing.T) {
	input := "select * from table_name"
	expected := []LexToken{
		LexToken{TokenSELECT, "select"},
		LexToken{TokenTIMES, "*"},
		LexToken{TokenFROM, "from"},
		LexToken{TokenKEY, "table_name"},
	}
	actual, err := NewLexer().Lex(input)
	if err != nil {
		t.Errorf("%v", err)
	}

	// Assert.
	if len(actual) != len(expected) {
		t.Errorf("actual: %v doesn't equals expected: %v.", actual, expected)
	}
	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("actual: %v doesn't equals expected: %v.", actual, expected)
		}
	}
}
