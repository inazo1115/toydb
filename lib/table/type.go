package table

type ToyDBType int

const (
	INT64 ToyDBType = iota
	STRING
)

func (t ToyDBType) String() string {
	switch t {
	case INT64:
		return "INT64"
	case STRING:
		return "STRING"
	default:
		return "Unknown"
	}
}

const (
	INT64Size = 8
)
