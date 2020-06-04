package client

import (
	"fmt"

	"github.com/inazo1115/toydb/lib/pkg"
	"github.com/inazo1115/toydb/lib/query"
)

type Client struct {
	interp *query.Interpreter
}

func NewClient() *Client {
	return &Client{query.NewInterpreter()}
}

func (c *Client) Query(q string) error {

	q = WashInput(q)
	result, err := c.interp.Interpret(q)
	if err != nil {
		return err
	}

	if len(result) == 0 {
		fmt.Println("ok")
		return nil
	}

	c.interp.PrintResult(result)
	return nil
}

func (c *Client) Version() string {
	return fmt.Sprintf("%s.%s", pkg.VERSION, pkg.REV)
}
