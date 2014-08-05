package main

import (
	"os"

	"github.com/hiremaga/sprout/ruby"
	"github.com/hiremaga/sprout/soloist"
)

func main() {
	ruby.GemInstall("soloist", ">= 1.0.3")
	ruby.GemInstall("librarian-chef", ">= 0.0.4")

	err := soloist.Run()

	if err != nil {
		os.Exit(1)
	}
}
