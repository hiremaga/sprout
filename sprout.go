package main

import (
	"os"
	"os/exec"
)

func main() {
	installGem("soloist", ">= 1.0.3")
	installGem("librarian-chef", ">= 0.0.4")

	soloist := `
	require 'rubygems'

	gem 'soloist', '>= 1.0.3'
	$0 = 'sprout' 
	load Gem.bin_path('soloist', 'soloist', '>= 1.0.3')
	`

	cmd := exec.Command("/usr/bin/ruby", append([]string{"-e", soloist}, os.Args[1:]...)...)
	attachAllStreams(cmd)
	ensureSystemRuby(cmd)

	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
}

func installGem(name, version string) error {
	check := exec.Command("/usr/bin/gem", "list", "--install", name, "--version", version)
	ensureSystemRuby(check)

	if err := check.Run(); err != nil {
		install := exec.Command("/usr/bin/sudo", "/usr/bin/gem", "install", "--no-ri", "--no-rdoc", name, "-v", version)
		attachAllStreams(install)
		ensureSystemRuby(install)

		if err = install.Run(); err != nil {
			return err
		}
	}

	return nil
}

func ensureSystemRuby(cmd *exec.Cmd) {
	cmd.Env = []string{
		os.ExpandEnv("PATH=${PATH}"),
	}
}

func attachAllStreams(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}
