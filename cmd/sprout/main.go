package main

import (
	"os"

	"github.com/hiremaga/sprout"
)

func main() {
	ruby := sprout.Ruby{}
	ruby.GemInstall("soloist", ">= 1.0.3")
	ruby.GemInstall("librarian-chef", ">= 0.0.4")

	soloist := sprout.Soloist{}
	err := soloist.Run()

	if err != nil {
		os.Exit(1)
	}
}
