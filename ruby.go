package sprout

import (
	"os"
	"os/exec"
)

type Ruby struct{}

func (r Ruby) Exec(code string, args ...string) error {
	cmd := rubyCommand("/usr/bin/ruby", append([]string{"-e", code}, args...)...)
	cmd.AttachAllStreams()
	cmd.EnsureSystemRuby()
	return cmd.Run()
}

func (r Ruby) GemInstall(name, version string) error {
	check := rubyCommand("/usr/bin/gem", "list", "--install", name, "--version", version)
	check.EnsureSystemRuby()

	if err := check.Run(); err != nil {
		install := rubyCommand("/usr/bin/sudo", "/usr/bin/gem", "install", "--no-ri", "--no-rdoc", name, "-v", version)
		install.AttachAllStreams()
		install.EnsureSystemRuby()

		if err = install.Run(); err != nil {
			return err
		}
	}

	return nil
}

type rubyCmd struct {
	*exec.Cmd
}

func rubyCommand(path string, args ...string) rubyCmd {
	return rubyCmd{exec.Command(path, args...)}
}

func (cmd *rubyCmd) EnsureSystemRuby() {
	cmd.Env = []string{
		os.ExpandEnv("PATH=${PATH}"),
	}
}

func (cmd *rubyCmd) AttachAllStreams() {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}
