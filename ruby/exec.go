package ruby

func Exec(code string, args ...string) error {
	cmd := Command("/usr/bin/ruby", append([]string{"-e", code}, args...)...)
	cmd.AttachAllStreams()
	cmd.EnsureSystemRuby()
	return cmd.Run()
}
