package client

import (
	"strings"
)

func WashInput(s string) string {
	s = strings.Replace(s, "\n", "", -1)
	s = strings.TrimRight(s, ";")
	return s
}
