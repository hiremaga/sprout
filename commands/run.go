package commands

import (
	"github.com/codegangsta/cli"

	"github.com/hiremaga/sprout/prereqs"
	"github.com/hiremaga/sprout/recipe"
)

var Run = cli.Command{
	Name:  "run",
	Usage: "Run's the specified recipe",
	Action: func(c *cli.Context) {
		prereqs.Install()
		recipe.Run(c.Args().First())
	},
}
