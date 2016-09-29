package cli

import (
	"errors"
	_ "net/http"
)

func Send(opts *CliOptions) error {
	if !opts.IsValid() {
		return errors.New("Invalid cli params. Bug (-b) and Message (-m) are obligatory.")
	}

	return nil
}
