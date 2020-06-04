package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/inazo1115/toydb/lib/client"
)

func main() {

	c := client.NewClient()
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("version: %s\n", c.Version())
	for {
		fmt.Printf("toydb> ")
		q, err := reader.ReadString(';')
		if err != nil {
			fmt.Println("exit.")
			break
		}

		if err := c.Query(q); err != nil {
			fmt.Printf("Error occured: %v\n", err)
		}
	}
}
