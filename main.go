package main

import (
	"flag"
	"fmt"

	"github.com/ericmdantas/bugz-cli/cli"
)

func main() {
	cli := cli.NewCliOptions()

	flag.StringVar(&cli.Msg, "m", "", "Deve ser a mensagem do commit: -m=conserta_algo")
	flag.StringVar(&cli.Bug, "b", "", "Deve ser o id do bug: -b=12346")
	flag.BoolVar(&cli.Close, "close", false, "Deve dizer se pode ou não fechar: -close")
	flag.Parse()

	fmt.Println(cli)
}
