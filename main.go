package main

import (
	"errors"
	"flag"
	"log"

	"github.com/ericmdantas/bugz-cli/cli"
)

func main() {
	opts := cli.NewCliOptions()

	flag.StringVar(&opts.Bug, "b", "", "Bug id -> Usage: -b=12346")
	flag.StringVar(&opts.Msg, "m", "", "Commit message -> Usage: -m=conserta_algo")
	flag.BoolVar(&opts.Close, "close", false, "Should close or not -> Usage: -close")
	flag.Parse()

	if err := cli.Send(opts); err != nil {
		log.Println(errors.New("There was an error while trying to send the files."))
		log.Fatalln(err)
	}
}
