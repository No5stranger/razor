package main

import (
	"github.com/No5stranger/razor/commands"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		commands.ShorterKeyFlag,
	}

	app.Commands = []cli.Command{
		commands.ShortUrlCommand,
		commands.FormateGithubEnterpriseCommand,
		commands.FormatTraceUrlCommand,
		commands.FormatSearchTraceCommand,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
