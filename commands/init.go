package commands

import (
	"github.com/codegangsta/cli"
)

var Init = cli.Command{
	Name:  "init",
	Usage: "Creates the various files that sprout needs in the current directory",
	Action: func(c *cli.Context) {
		println("TODO: Create .gitignore")
		println("TODO: Create site-cookbooks/ directory")
		println("TODO: Create site-cookbooks/example cookbook")
		println("TODO: Create Sproutfile")
	},
}
