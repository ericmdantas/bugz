package cli

import (
	"testing"
)

var tableCliValidation = []struct {
	in  CliOptions
	out bool
}{
	{
		in: CliOptions{
			Bug:   "",
			Msg:   "",
			Close: false,
		},
		out: false,
	},
	{
		in: CliOptions{
			Bug:   "1",
			Msg:   "",
			Close: false,
		},
		out: false,
	},
	{
		in: CliOptions{
			Bug:   "1",
			Msg:   "2",
			Close: false,
		},
		out: true,
	},
	{
		in: CliOptions{
			Bug:   "1",
			Msg:   "2",
			Close: true,
		},
		out: true,
	},
	{
		in: CliOptions{
			Bug:   "",
			Msg:   "2",
			Close: true,
		},
		out: false,
	},
	{
		in: CliOptions{
			Bug:   "",
			Msg:   "",
			Close: true,
		},
		out: false,
	},
}

func TestNewCliOptions(t *testing.T) {
	c := NewCliOptions()

	defaultBug := ""
	defaultMsg := ""
	defaultClose := false

	if c.Bug != defaultBug {
		t.Fatalf("Expected bug to equal %s, but got: %s", defaultBug, c.Bug)
	}

	if c.Msg != defaultMsg {
		t.Fatalf("Expected Msg to equal %s, but got: %s", defaultMsg, c.Msg)
	}

	if c.Close != defaultClose {
		t.Fatalf("Expected Close to equal %v, but got: %s", defaultClose, c.Close)
	}
}

func TestCliValidation(t *testing.T) {
	for _, v := range tableCliValidation {
		cli := NewCliOptions()

		cli.Bug = v.in.Bug
		cli.Msg = v.in.Msg
		cli.Close = v.in.Close

		if cli.IsValid() != v.out {
			t.Fatalf("Expected cli to be %t, but got %t", v.out, cli.IsValid())
		}
	}
}
