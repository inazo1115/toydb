package util

import (
	"errors"
)

func CopyStrToByte(s string, buf []byte) error {

	if len(buf) < len(s) {
		return errors.New("buffer size is not enough")
	}

	for i := 0; i < len(s); i++ {
		buf[i] = s[i]
	}

	return nil
}
