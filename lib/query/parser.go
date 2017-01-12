package query

import (
	"errors"
	//"fmt"
)

var (
	ParseError = errors.New("ParseError")
)

//
// ASTNode
//

type ASTNode struct {
	Token    LexToken
	Children []*ASTNode
}

func NewASTNode(token LexToken) *ASTNode {
	return &ASTNode{token, make([]*ASTNode, 0)}
}

func (n *ASTNode) AppendChild(node *ASTNode) {
	n.Children = append(n.Children, node)
}

//
// Parser
//

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(tokens []LexToken) (*ASTNode, error) {
	switch {
	case tokens[0].ID == TokenCREATE && tokens[1].ID == TokenTABLE:
		return p.parseCreateTable(tokens)
	case tokens[0].ID == TokenINSERT:
		return p.parseInsert(tokens)
	case tokens[0].ID == TokenSELECT:
		return p.parseSelect(tokens)
	default:
		return nil, ParseError
	}
}

func (p *Parser) parseCreateTable(tokens []LexToken) (*ASTNode, error) {

	root := NewASTNode(tokens[0])

	if tokens[2].ID != TokenKEY {
		return nil, ParseError
	}

	table := NewASTNode(tokens[2])
	root.AppendChild(table)

	if tokens[3].ID != TokenLPAREN {
		return nil, ParseError
	}

	for i := 4; i < len(tokens); i++ {

		if tokens[i].ID == TokenRPAREN {
			break
		}

		if tokens[i].ID == TokenCOMMA {
			continue
		}

		if tokens[i].ID != TokenKEY {
			return nil, ParseError
		}

		root.AppendChild(NewASTNode(tokens[i]))

		i++

		switch tokens[i].ID {
		case TokenINT:
			root.AppendChild(NewASTNode(tokens[i]))
		case TokenSTRING:
			if !(tokens[i+1].ID == TokenLPAREN &&
				tokens[i+2].ID == TokenVALUE &&
				tokens[i+3].ID == TokenRPAREN) {
				return nil, ParseError
			}
			str := NewASTNode(tokens[i])
			size := NewASTNode(tokens[i+2])
			str.AppendChild(size)
			root.AppendChild(str)

			i += 3
		}
	}

	return root, nil
}

func (p *Parser) parseInsert(tokens []LexToken) (*ASTNode, error) {

	root := NewASTNode(tokens[0])

	if tokens[1].ID != TokenINTO {
		return nil, ParseError
	}

	table := NewASTNode(tokens[2])
	root.AppendChild(table)

	if tokens[3].ID != TokenLPAREN {
		return nil, ParseError
	}

	size := 0
	for i := 4; i < len(tokens); i++ {

		if tokens[i].ID == TokenRPAREN {
			break
		}

		if tokens[i].ID == TokenCOMMA {
			continue
		}

		if tokens[i].ID != TokenKEY {
			return nil, ParseError
		}

		size++
	}

	for i := 0; i < size; i++ {

		col := NewASTNode(tokens[(i*2)+4])
		val := NewASTNode(tokens[(i*2)+4+(size*2)+2])

		if col.Token.ID != TokenKEY {
			return nil, ParseError
		}
		if val.Token.ID != TokenVALUE {
			return nil, ParseError
		}

		col.AppendChild(val)
		root.AppendChild(col)
	}

	return root, nil
}

func (p *Parser) parseSelect(tokens []LexToken) (*ASTNode, error) {

	root := NewASTNode(tokens[0])

	table := NewASTNode(tokens[len(tokens)-1])
	root.AppendChild(table)

	if tokens[len(tokens)-2].ID != TokenFROM {
		return nil, ParseError
	}

	for i := 1; i < len(tokens)-2; i++ {

		if tokens[i].ID == TokenCOMMA {
			continue
		}

		if !(tokens[i].ID == TokenTIMES || tokens[i].ID == TokenKEY) {
			return nil, ParseError
		}
		root.AppendChild(NewASTNode(tokens[i]))
	}

	return root, nil
}
