package cli

type File struct {
	files []string
}

func GetFilesModified() (*File, error) {
	f := File{}
	var fs []string

	fs = append(fs, []string{
		"___test1.txt",
		"___test2.txt",
		"___test3.txt",
	}...)

	f.files = fs

	return &f, nil
}
