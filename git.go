package gitconfig

import (
	"os/exec"
)

type commander interface {
	CombinedOutput(string, ...string) ([]byte, error)
}

type commandRunner struct{}

func (c commandRunner) CombinedOutput(command string, args ...string) ([]byte, error) {
	return exec.Command(command, args...).CombinedOutput()
}

func getGitConfigUserName(g commander) (string, error) {
	out, err := g.CombinedOutput("git", "config", "user.name")
	return string(out), err
}

func getGitConfigUserEmail(g commander) (string, error) {
	out, err := g.CombinedOutput("git", "config", "user.email")
	return string(out), err
}

func setGitConfigUserEmail(g commander, userEmail string) (string, error) {
	out, err := g.CombinedOutput("git", "config", "user.email", userEmail)
	return string(out), err
}

func setGitConfigUserName(g commander, userName string) (string, error) {
	out, err := g.CombinedOutput("git", "config", "user.name", userName)
	return string(out), err
}
