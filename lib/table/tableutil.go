package table

import (
	"fmt"
	"strings"
)

func PrintResult(schema *Schema, records []*Record) {

	sizes := make([]int, len(schema.Columns()))

	for i := 0; i < len(schema.Columns()); i++ {

		sizes[i] = -1

		s := len(string(schema.Columns()[i].Name()))

		if sizes[i] < s {
			sizes[i] = s
		}

		for j := 0; j < len(records); j++ {
			s := lenStrByte(strings.TrimSpace(records[j].Values()[i].String()))
			if sizes[i] < s {
				sizes[i] = s
			}
		}
	}

	fmt.Printf("+")
	for _, s := range sizes {
		fmt.Printf(strings.Repeat("-", s+2))
		fmt.Printf("+")
	}
	fmt.Printf("\n")

	fmt.Printf("|")
	for i, s := range sizes {
		name := schema.Columns()[i].Name()
		ws := s - len(name)
		fmt.Printf(strings.Repeat(" ", 1))
		fmt.Printf(name)
		fmt.Printf(strings.Repeat(" ", ws+1))
		fmt.Printf("|")
	}
	fmt.Printf("\n")

	fmt.Printf("+")
	for _, s := range sizes {
		fmt.Printf(strings.Repeat("-", s+2))
		fmt.Printf("+")
	}
	fmt.Printf("\n")

	for _, r := range records {
		fmt.Printf("|")
		for i, s := range sizes {
			name := r.Values()[i].String()
			ws := s - lenStrByte(name)
			switch schema.Columns()[i].Type() {
			case INT64:
				fmt.Printf(strings.Repeat(" ", ws+1))
				fmt.Printf(name)
				fmt.Printf(strings.Repeat(" ", 1))
				fmt.Printf("|")
			case STRING:
				fmt.Printf(strings.Repeat(" ", 1))
				fmt.Printf(name)
				fmt.Printf(strings.Repeat(" ", ws+1))
				fmt.Printf("|")
			default:
				panic("will not reach here")
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("+")
	for _, s := range sizes {
		fmt.Printf(strings.Repeat("-", s+2))
		fmt.Printf("+")
	}
	fmt.Printf("\n")
}

func lenStrByte(s string) int {
	size := 0
	for _, b := range s {
		if int(b) == int(0) {
			continue
		}
		size++
	}
	return size
}
