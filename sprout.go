package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "sprout"
	app.Usage = "Automatically configure of your Mac"

	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "Creates the various files that sprout needs in the current directory",
			Action: func(c *cli.Context) {
				println("TODO: Create .gitignore")
				println("TODO: Create site-cookbooks/ directory")
				println("TODO: Create site-cookbooks/example cookbook")
				println("TODO: Create Sproutfile")
			},
		},
		{
			Name:  "run",
			Usage: "Run's the specified recipe",
			Action: func(c *cli.Context) {
				println("TODO: Run pre-requisites (e.g. librarian-chef/berkshelf)")
				println("TODO: Run recipe with chef-solo: ", c.Args().First())
			},
		},
	}

	app.Run(os.Args)
}
