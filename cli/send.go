package cli

import (
	"errors"
	"fmt"
	_ "net/http"
)

func Send(opts *CliOptions) error {
	if !opts.IsValid() {
		return errors.New("Invalid cli params. Bug (-b) and Message (-m) are obligatory.")
	}

	f, err := GetFilesModifiedPaths()

	if err != nil {
		panic(err)
	}

	for _, v := range f.filesPath {
		fmt.Println(v)
	}

	return nil
}
