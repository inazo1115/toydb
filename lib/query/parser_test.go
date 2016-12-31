package query_test

import (
	"testing"

	"github.com/inazo1115/toydb/lib/query"
)

func TestParse(t *testing.T) {
	actual, err := query.Parse("select * from table;")

	if err != nil {
		t.Errorf("%s", err)
	}

	expected := query.QueryAST{"select", "bar"}

	if actual != expected {
		t.Errorf("error")
	}
}
