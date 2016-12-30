package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/inazo1115/toydb/lib/client"
)

func main() {
	fmt.Printf("version: %s\n", client.Version())

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("toydb> ")
		query, err := reader.ReadString(';')
		if err != nil {
			// panic(err)
			fmt.Println("exit.")
			break
		}
		query = strings.TrimRight(query, "\n")
		result, err := client.Query(query)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	}
}
