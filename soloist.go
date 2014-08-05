package sprout

import "os"

type Soloist struct {
	ruby Ruby
}

func (s Soloist) Run() error {
	return s.ruby.Exec(`
	require 'rubygems'

	gem 'soloist', '>= 1.0.3'
	$0 = 'sprout'
	load Gem.bin_path('soloist', 'soloist', '>= 1.0.3')
	`, os.Args[1:]...)
}
