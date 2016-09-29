package cli

import (
	"testing"
)

func TestNewCliOptions(t *testing.T) {
	c := NewCliOptions()
	
	correctInfo := "yo"
	
	if c.Info != correctInfo {
		t.Fatalf("Expected info to equal %s, but got: %s", correctInfo, c.Info)
	}
}