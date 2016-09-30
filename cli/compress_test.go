package cli

import (
	"testing"
)

func TestCompress(t *testing.T) {
	r, err := Compress([]byte("a"))

	if err != nil {
		t.Fatal("Expected error NOT to occur, but there was one.")
	}

	if len(r) == 0 {
		t.Fatal("Expected returned info to be compressed.")
	}
}
