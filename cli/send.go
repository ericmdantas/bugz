package cli

import (
	"errors"
	"log"
	_ "net/http"
)

func Send(opts *CliOptions) error {
	if !opts.IsValid() {
		return errors.New("Invalid cli params. Bug (-b) and Message (-m) are obligatory.")
	}

	c := NewContent()
	c.Transform()

	if c.IsEmpty() {
		log.Fatalln(errors.New("Content is empty."))
	}

	return nil
}
