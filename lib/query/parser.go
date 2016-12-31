package query

import (
	"fmt"
//	"regexp"
	"strings"
)

type QueryAST struct {
	Method string
	Table  string
}

func NewSelectAST(table string) QueryAST {
	return QueryAST{"select", table}
}

func Parse(query string) (QueryAST, error) {
	query = strings.ToLower(query)

	switch cmd := strings.Split(query, " ")[0]; cmd {
	case "select":
		return NewSelectAST("bar"), nil
	default:
		return QueryAST{}, fmt.Errorf("can't parse: %s", query)
	}
}
