package cli

import (
	"testing"
)

func TestGetFilesModified(t *testing.T) {
	f, r := GetFilesModified()

	if len(f.files) == 0 {
		t.Fatalf("Expected files to be filled, but it's not.")
	}

	if r != nil {
		t.Fatalf("Expected no error, but got one: %v", r)
	}
}
