package cli

import (
	"testing"
)

func TestCompress(t *testing.T) {
	r := Compress()

	if r != nil {
		t.Fatal("Expected no error to occur, but there was one.")
	}
}
