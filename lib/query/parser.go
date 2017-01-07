package query

import (
	"errors"
	"fmt"
)

type AST interface {
	Eval() error
	Inspect(key string) string
}

type ASTCreateTable struct {
	table string
	keys  []LexToken
	types []ColumnType
}

type ASTInsert struct {
	table  string
	cols   []string
	values []string
}

type ASTSelect struct {
	name  string
	isAll bool
	cols  []string
	limit int
}

type ColumnType struct {
	id   LexTokenID
	meta map[string]string
}

func (ast ASTCreateTable) Eval() error {
	return errors.New("")
}

func (ast ASTCreateTable) Inspect(key string) string {
	switch key {
	case "table":
		return fmt.Sprintf("%v", ast.table)
	case "keys":
		return fmt.Sprintf("%v", ast.keys)
	case "types":
		return fmt.Sprintf("%v", ast.types)
	default:
		panic("foo")
	}
}

func Parse(tokens []LexToken) (AST, error) {
	switch {
	case tokens[0].ID == TokenCREATE && tokens[1].ID == TokenTABLE:
		return ParseCreateTable(tokens)
	case tokens[0].ID == TokenINSERT:
		return ASTCreateTable{}, ParseError
		//return ParseInsert(tokens)
	case tokens[0].ID == TokenSELECT:
		return ASTCreateTable{}, ParseError
		//return ParseSelect(tokens)
	default:
		return ASTCreateTable{}, ParseError
	}
}

func ParseCreateTable(tokens []LexToken) (ASTCreateTable, error) {

	table := tokens[2].Val
	keys := make([]LexToken, 0)
	types := make([]ColumnType, 0)

	for i := 4; i < len(tokens); i++ {

		if tokens[i].ID == TokenLPAREN || tokens[i].ID == TokenCOMMA {
			continue
		}

		if tokens[i].ID == TokenRPAREN {
			break
		}

		if tokens[i].ID != TokenKEY {
			return ASTCreateTable{}, ParseError
		}
		keys = append(keys, tokens[i])

		i++

		switch tokens[i].ID {
		case TokenINT:
			types = append(types, ColumnType{tokens[i].ID, map[string]string{}})
		case TokenSTRING:
			if !(tokens[i+1].ID == TokenLPAREN &&
				tokens[i+2].ID == TokenVALUE &&
				tokens[i+3].ID == TokenRPAREN) {
				return ASTCreateTable{}, ParseError
			}
			meta := map[string]string{
				"size": tokens[i+2].Val,
			}
			types = append(types, ColumnType{tokens[i].ID, meta})
			i += 3
		}
	}

	return ASTCreateTable{table, keys, types}, nil
}

var (
	ParseError = errors.New("ParseError")
)
