package cli

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	empty := NewContent()
	filled := NewContent()
	filled.Files["a"] = newFileInfo([]byte("a"), []byte("b"))

	if !empty.IsEmpty() {
		t.Fatalf("Expected content to be empty, but it wasn't.")
	}

	if filled.IsEmpty() {
		t.Fatalf("Expected content NOT to be empty, but it was.")
	}
}
