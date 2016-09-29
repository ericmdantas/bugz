package main

import (
	"errors"
	"flag"
	"log"

	"github.com/ericmdantas/bugz-cli/cli"
)

func main() {
	opts := cli.NewCliOptions()

	flag.StringVar(&opts.Bug, "b", "", "Deve ser o id do bug: -b=12346")
	flag.StringVar(&opts.Msg, "m", "", "Deve ser a mensagem do commit: -m=conserta_algo")
	flag.BoolVar(&opts.Close, "close", false, "Deve dizer se pode ou não fechar: -close")
	flag.Parse()

	if err := cli.Send(opts); err != nil {
		log.Println(errors.New("There was an error while trying to send the files."))
		log.Fatalln(err)
	}
}
