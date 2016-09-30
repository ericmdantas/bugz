package cli

import (
	"testing"
)

func TestGetFilesModifiedPath(t *testing.T) {
	f, r := GetFilesModifiedPaths()

	if len(f.filesPath) == 0 {
		t.Fatalf("Expected filesPath to be filled, but it's not.")
	}

	if r != nil {
		t.Fatalf("Expected no error, but got one: %v", r)
	}
}
