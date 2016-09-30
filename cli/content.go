package cli

import (
	"io/ioutil"
)

type Content struct {
	Files map[string]struct {
		BodyRaw        []byte
		BodyCompressed []byte
	}
}

func (c *Content) Transform() error {
	var paths []string

	paths = append(paths, []string{
		"___test1.txt",
		"___test2.txt",
		"___test3.txt",
	}...)

	for _, v := range paths {
		content, err := ioutil.ReadFile(v)

		c.Files[v].BodyRaw = content
		if c.Files[v].BodyCompressed, err = Compress(content); err != nil {
			panic(err)
		}
	}

	return nil
}

func (c *Content) IsEmpty() bool {
	return len(c.Files) > 0
}

func NewContent() *Content {
	return &Content{
		Files: make(map[string]struct {
			BodyRaw        []byte
			BodyCompressed []byte
		}),
	}
}
