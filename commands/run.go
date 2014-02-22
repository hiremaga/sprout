package commands

import (
	"github.com/codegangsta/cli"
)

var Run = cli.Command{
	Name:  "run",
	Usage: "Run's the specified recipe",
	Action: func(c *cli.Context) {
		println("TODO: Run pre-requisites (e.g. librarian-chef/berkshelf)")
		println("TODO: Run recipe with chef-solo: ", c.Args().First())
	},
}
