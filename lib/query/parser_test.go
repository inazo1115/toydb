package query

import (
	//"fmt"
	"testing"
)

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

	expected := "table_name"
	if actual.Inspect("table") != expected {
		t.Errorf("actual: %v doesn't equals expected: %v.", actual, expected)
	}

	expected = "[{0 name} {0 age} {0 tel}]"
	if actual.Inspect("keys") != expected {
		t.Errorf("actual: %v doesn't equals expected: %v.", actual, expected)
	}

	expected = "[{59 map[size:20]} {60 map[]} {60 map[]}]"
	if actual.Inspect("types") != expected {
		t.Errorf("actual: %v doesn't equals expected: %v.", actual, expected)
	}

}
