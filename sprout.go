package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "sprout"
	app.Usage = "Automatically configure of your Mac"
	app.Action = func(c *cli.Context) {
		println("Coming soon: much awesomeness.")
	}

	app.Run(os.Args)
}
