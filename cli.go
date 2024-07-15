package main

import (
	"errors"
	"fmt"

	"github.com/urfave/cli/v2"
	"golang.design/x/clipboard"
)

var model string
var clipboardRead bool

type Input struct {
	Model string
	Prompt string
}

var helpTemplate = `
NAME:
{{.Name}} - {{.Usage}}
USAGE:
	{{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
	{{if len .Authors}}
AUTHOR:
	{{range .Authors}}{{ . }}{{end}}
	{{end}}{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
GLOBAL OPTIONS:
	{{range .VisibleFlags}}{{.}}
	{{end}}{{end}}{{if .Copyright }}
COPYRIGHT:
	{{.Copyright}}
	{{end}}{{if .Version}}
VERSION:
	{{.Version}}
	{{end}}
`

var appConfig = &cli.App{
	Name:  "jspark",
	Version: "v0.0.1",
	Usage: "Jspark generates jira ticket TITLE nad SUMMARY based on given TEXT. It relays on local LLM models to accomplish this task.",
	Action: func(cCtx *cli.Context) error {
		fmt.Println("No arguments provided")
		return errors.New("no arguments provided, use 'help' to see available options")
	},
	Commands: []*cli.Command{
		{
			Name:    "ollama",
			Aliases: []string{"o"},
			Usage:   "use local Ollama model",
			Action: func(cCtx *cli.Context) error {
				var prompt string
				if clipboardRead {
					err := clipboard.Init()
					if err != nil {
     				panic(err)
					}
					prompt = string(clipboard.Read(clipboard.FmtText))
				} else {
					prompt = cCtx.Args().Get(0)
				}
				input := Input {
					Model: model,
					Prompt: prompt,
				}
					co, err := callOllama(input)
					if err != nil {
						return err
					}

					co, err = processCompletion(co)
					if err != nil {
						return err
					}
					co.PrintObject()
					return nil

			},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "model",
				Aliases: []string{"m"},
				Value:   "ticketer:latest",
				Usage:   "Local Ollama model to use.",
				EnvVars: []string{"JSPARK_OLLAMA_MODEL"},
				Destination: &model,
			},
			clipboardCliFlag,
		},
	},
	{
			Name:    "gpt",
			Aliases: []string{"g"},
			Usage:   "use ChatGPT model",
			Action: func(cCtx *cli.Context) error {
				var prompt string
				if clipboardRead {
					err := clipboard.Init()
					if err != nil {
     				panic(err)
					}
					prompt = string(clipboard.Read(clipboard.FmtText))
				} else {
					prompt = cCtx.Args().Get(0)
				}
				input := Input {
					Model: model,
					Prompt: prompt,
				}
				co, err := callGPT(input)
				if err != nil {
					return err
				}

				co, err = processCompletion(co)
				if err != nil {
					return err
				}
				co.PrintObject()
				return nil

			},
			Flags: []cli.Flag{
				clipboardCliFlag,
			},
	},
	},
}


var clipboardCliFlag = &cli.BoolFlag{
	Name:    "clipboard",
	Aliases: []string{"c"},
	Usage:   "Read prompt form clipboard",
	Destination: &clipboardRead,
}
