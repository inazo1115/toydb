package table

// Record

type Record struct {
	values []*Value
}

func NewRecord(values []*Value) *Record {
	return &Record{values}
}

func (r *Record) Values() []*Value {
	return r.values
}

// Value

type Value struct {
	type_   ToyDBType
	vInt64  int64
	vString string
}

func NewValueInt64(v int64) *Value {
	return &Value{INT64, v, ""}
}

func NewValueString(v string) *Value {
	return &Value{STRING, -1, v}
}

func (v *Value) Type() ToyDBType {
	return v.type_
}
