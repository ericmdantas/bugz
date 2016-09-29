package cli

type CliOptions struct {
	Bug   string
	Msg   string
	Close bool
}

func (c *CliOptions) IsValid() bool {
	if c.Bug != "" && c.Msg != "" {
		return true
	}

	return false
}

func NewCliOptions() *CliOptions {
	return &CliOptions{
		Bug:   "",
		Msg:   "",
		Close: false,
	}
}
