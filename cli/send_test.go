package cli

import (
	"errors"
	"reflect"
	"testing"
)

var tableSendCliValidation = []struct {
	in  *CliOptions
	out error
}{
	{
		in: &CliOptions{
			Bug:   "",
			Msg:   "",
			Close: true,
		},
		out: errors.New("Invalid cli params. Bug (-b) and Message (-m) are obligatory."),
	},
	{
		in: &CliOptions{
			Bug:   "1",
			Msg:   "2",
			Close: true,
		},
		out: nil,
	},
}

func TestSendCliValidation(t *testing.T) {
	for _, v := range tableSendCliValidation {
		if r := Send(v.in); !reflect.DeepEqual(r, v.out) {
			t.Fatalf(`Expected send to to return "%v", but it returned "%v".`, v.out, r)
		}
	}
}
