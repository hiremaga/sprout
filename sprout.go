package main

import (
	"os"
	"os/exec"
)

func main() {
	soloist := `
	require 'rubygems'

	gem 'soloist', '>= 1.0.3'
	$0 = 'sprout' 
	load Gem.bin_path('soloist', 'soloist', '>= 1.0.3')
	`
	cmd := exec.Command("/usr/bin/ruby", append([]string{"-e", soloist, "--"}, os.Args[1:]...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = []string{} // clean env to ensure system ruby
	cmd.Run()
}
