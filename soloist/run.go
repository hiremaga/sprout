package soloist

import (
	"os"

	"github.com/hiremaga/sprout/ruby"
)

func Run() error {
	return ruby.Exec(`
	require 'rubygems'

	gem 'soloist', '>= 1.0.3'
	$0 = 'sprout'
	load Gem.bin_path('soloist', 'soloist', '>= 1.0.3')
	`, os.Args[1:]...)
}
