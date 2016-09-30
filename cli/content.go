package cli

import (
	"errors"
	"io/ioutil"
	"log"
)

type fileInfo struct {
	BodyRaw        []byte
	BodyCompressed []byte
}

type Content struct {
	Files map[string]*fileInfo
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

		if err != nil {
			log.Fatalf(errors.New("There was an error trying to read the file content: %s").Error(), v)
		}

		contentCompressed, err := Compress(content)

		if err != nil {
			log.Fatalf(errors.New("There was an error trying to compress the file. Error: %s").Error(), err)
		}

		c.Files[v] = &fileInfo{
			BodyRaw:        content,
			BodyCompressed: contentCompressed,
		}
	}

	return nil
}

func (c *Content) IsEmpty() bool {
	return len(c.Files) == 0
}

func NewContent() *Content {
	return &Content{
		Files: make(map[string]*fileInfo),
	}
}

func newFileInfo(bodyRaw, bodyCompressed []byte) *fileInfo {
	return &fileInfo{
		BodyRaw:        bodyRaw,
		BodyCompressed: bodyCompressed,
	}
}
