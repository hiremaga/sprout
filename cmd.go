package sprout

import (
	"os"
	"os/exec"
)

type Cmd struct {
	*exec.Cmd
}

func Command(path string, args ...string) Cmd {
	return Cmd{exec.Command(path, args...)}
}

func (cmd *Cmd) EnsureSystemRuby() {
	cmd.Env = []string{
		os.ExpandEnv("PATH=${PATH}"),
	}
}

func (cmd *Cmd) AttachAllStreams() {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}
