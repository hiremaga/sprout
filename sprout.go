package main

import (
	"github.com/codegangsta/cli"
	"github.com/hiremaga/sprout/commands"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "sprout"
	app.Usage = "Configure your Mac"

	app.Commands = []cli.Command{
		commands.Init,
		commands.Run,
	}

	app.Run(os.Args)
}
