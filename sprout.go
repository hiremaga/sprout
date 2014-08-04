package main

import (
	"os"
	"os/exec"
)

func main() {
	install("soloist", ">= 1.0.3")
	install("librarian-chef", ">= 0.0.4")

	soloist := `
	require 'rubygems'

	gem 'soloist', '>= 1.0.3'
	$0 = 'sprout' 
	load Gem.bin_path('soloist', 'soloist', '>= 1.0.3')
	`

	cmd := exec.Command("/usr/bin/ruby", append([]string{"-e", soloist}, os.Args[1:]...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = []string{} // clean env to ensure system ruby

	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
}

func install(name, version string) error {
	check := exec.Command("/usr/bin/gem", "list", "--installed", name, "--version", version)
	check.Env = []string{}

	if err := check.Run(); err != nil {
		install := exec.Command("/usr/bin/sudo", "/usr/bin/gem", "install", "--no-ri", "--no-rdoc", name, "-v", version)
		install.Stdin = os.Stdin
		install.Stdout = os.Stdout
		install.Stderr = os.Stderr
		install.Env = []string{} // clean env to ensure system ruby

		if err = install.Run(); err != nil {
			return err
		}
	}

	return nil
}
