package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ofco secret share"
	app.Version = "0.1"
	app.Usage = "launch a quick and simple http server to share encrypted secrets"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE` (defaults to conf.json)",
		},
	}
	app.Action = runApp

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
