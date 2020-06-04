package client

import (
	"strings"
)

func WashInput(s string) string {
	return strings.TrimRight(s, ";")
}
