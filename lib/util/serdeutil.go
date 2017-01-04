package util

import (
	"encoding/binary"
)

func SerializeInt64(v int64) []byte {
	buf := make([]byte, 8)
	binary.PutVarint(buf, v)
	return buf
}

func DeserializeInt64(b []byte) int64 {
	v, x := binary.Varint(b)
	if x <= 0 {
		panic("DeserializeInt64 failed.")
	}
	return v
}

func SerializeString(v string) []byte {
	return []byte(v)
}

func DeserializeString(b []byte) string {
	return string(b)
}
