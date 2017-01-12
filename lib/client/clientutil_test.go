package client

import (
	"testing"

	"github.com/inazo1115/toydb/lib/util"
)

func TestWashInput(t *testing.T) {
	util.Assert(t, WashInput("foobar;"), "foobar")
	util.Assert(t, WashInput("foobar;;"), "foobar")
	util.Assert(t, WashInput("foo\nbar;"), "foobar")
	util.Assert(t, WashInput("\nfoo\nbar\n;"), "foobar")
}
