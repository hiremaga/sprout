package ruby

func GemInstall(name, version string) error {
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
