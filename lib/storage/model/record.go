package record

/*import (
	"fmt"
)*/

type Record struct {
	name string
	age int
	message string
}

func NewRecord(name string, age int, message string) {
	return Record{name, age, message}
}
