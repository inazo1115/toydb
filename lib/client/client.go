package client

import (
	"fmt"
	"strings"

	"github.com/inazo1115/toydb/lib/pkg"
	"github.com/inazo1115/toydb/lib/query"
	"github.com/inazo1115/toydb/lib/table"
)

type Client struct {
	ev *query.Evaluator
}

func NewClient() *Client {
	return &Client{query.NewEvaluator()}
}

func CleanInput(s string) string {
	s = strings.Replace(s, "\n", "", -1)
	s = strings.TrimRight(s, ";")
	return s
}

func (c *Client) Query(q string) error {

	q = CleanInput(q)

	tokens, err := query.Lex(q)
	if err != nil {
		return err
	}

	ast, err := query.Parse(tokens)
	if err != nil {
		return err
	}

	res, err := c.ev.Eval(ast)
	if err != nil {
		return err
	}

	if len(res) == 0 {
		fmt.Println("ok")
		return nil
	}

	table.PrintResult(c.ev.Runtime().CurrentSchema(), res)
	return nil
}

func (c *Client) Version() string {
	return fmt.Sprintf("%s.%s", pkg.VERSION, pkg.REV)
}
