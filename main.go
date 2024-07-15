package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// TODO:
// - use interface for calling models (abstraction)
// - text templates for ticket proposal printout
// - sending ticket to Jira

func main() {

	cli.AppHelpTemplate = helpTemplate
	app := appConfig

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
