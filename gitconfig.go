package gitconfig

type GitConfig struct {
}

var (
	cmd commandRunner
)

func (g GitConfig) Email() (string, error) {
	return getGitConfigUserEmail(cmd)
}
func (g GitConfig) SetEmail(email string) (string, error) {
	return setGitConfigUserEmail(cmd, email)
}

func (g GitConfig) UserName() (string, error) {
	return getGitConfigUserName(cmd)
}
func (g GitConfig) SetUserName(name string) (string, error) {
	return setGitConfigUserName(cmd, name)
}
