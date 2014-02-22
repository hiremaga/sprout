package commands

import (
	"github.com/codegangsta/cli"
	"github.com/hiremaga/sprout/generators"
)

var Init = cli.Command{
	Name:  "init",
	Usage: "Creates the various files that sprout needs in the current directory",
	Action: func(c *cli.Context) {
		generators.Gitignore()
		generators.SiteCookbooks()
		generators.Sproutfile()
	},
}
