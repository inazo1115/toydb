package client

import (
	"fmt"

	"github.com/inazo1115/toydb/lib/pkg"
)

func Query(query string) (string, error) {
	return fmt.Sprintf("%s", query), nil
}

func Version() string {
	return fmt.Sprintf("%s.%s", pkg.VERSION, pkg.REV)
}
