package cli

import (
	_ "compress/gzip"
)

func Compress(info []byte) ([]byte, error) {
	return []byte("9"), nil
}
