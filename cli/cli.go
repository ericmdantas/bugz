package cli

type CliOptions struct {
	Info string
}

func NewCliOptions() *CliOptions {
	return &CliOptions {
		Info: "yo",
	}
}