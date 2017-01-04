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

	printBorder(sizes)

	fmt.Printf("|")
	for i, s := range sizes {
		name := schema.Columns()[i].Name()
		ws := s - len(name)
		printLeftAligned(name, ws)
		fmt.Printf("|")
	}
	fmt.Printf("\n")

	printBorder(sizes)

	for _, r := range records {
		fmt.Printf("|")
		for i, s := range sizes {
			value := r.Values()[i].String()
			ws := s - lenStrByte(value)

			switch schema.Columns()[i].Type() {
			case INT64:
				printRightAligned(value, ws)
			case STRING:
				printLeftAligned(value, ws)
			default:
				panic("will not reach here")
			}

			fmt.Printf("|")
		}
		fmt.Printf("\n")
	}

	printBorder(sizes)
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

func printBorder(sizes []int) {
	fmt.Printf("+")
	for _, s := range sizes {
		fmt.Printf(strings.Repeat("-", s+2))
		fmt.Printf("+")
	}
	fmt.Printf("\n")
}

func printLeftAligned(s string, ws int) {
	fmt.Printf(strings.Repeat(" ", 1))
	fmt.Printf(s)
	fmt.Printf(strings.Repeat(" ", ws+1))
}

func printRightAligned(s string, ws int) {
	fmt.Printf(strings.Repeat(" ", ws+1))
	fmt.Printf(s)
	fmt.Printf(strings.Repeat(" ", 1))
}
