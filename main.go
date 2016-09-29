package main

import (
	"fmt"
	"flag"
	"github.com/ericmdantas/bugz-cli/cli"
)

func main() {
	cli := cli.NewCliOptions()

	flag.StringVar(&cli.Info, "info", "", "-info=wtf")
	flag.Parse()

	fmt.Println(cli)
}