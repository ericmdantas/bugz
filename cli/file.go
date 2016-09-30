package cli

type File struct {
	filesPath []string
}

func GetFilesModifiedPaths() (*File, error) {
	f := File{}

	f.filesPath = append(f.filesPath, []string{
		"___test1.txt",
		"___test2.txt",
		"___test3.txt",
	}...)

	return &f, nil
}

func GetFilesModifiedContent() {

}
