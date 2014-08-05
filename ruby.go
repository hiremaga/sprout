package sprout

type Ruby struct{}

func (r Ruby) Exec(code string, args ...string) error {
	cmd := Command("/usr/bin/ruby", append([]string{"-e", code}, args...)...)
	cmd.AttachAllStreams()
	cmd.EnsureSystemRuby()
	return cmd.Run()
}

func (r Ruby) GemInstall(name, version string) error {
	check := Command("/usr/bin/gem", "list", "--install", name, "--version", version)
	check.EnsureSystemRuby()

	if err := check.Run(); err != nil {
		install := Command("/usr/bin/sudo", "/usr/bin/gem", "install", "--no-ri", "--no-rdoc", name, "-v", version)
		install.AttachAllStreams()
		install.EnsureSystemRuby()

		if err = install.Run(); err != nil {
			return err
		}
	}

	return nil
}
