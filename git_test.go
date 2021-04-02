package gitconfig

import (
	"errors"
	"testing"
)

type testCommandRunner struct{}

func (c testCommandRunner) CombinedOutput(command string, args ...string) ([]byte, error) {
	if args[0] == "config" {
		if args[1] == "user.name" {
			if len(args) == 3 {
				return []byte(""), nil
			} else {
				return []byte("Some User"), nil
			}
		}
		if args[1] == "user.email" {
			if len(args) == 3 {
				return []byte(""), nil
			} else {
				return []byte("me@example.com"), nil
			}
		}
	}
	return []byte("Unknown Command"), errors.New("Unknown Command")
}

func TestGetGitConfigUserName(test *testing.T) {
	var cmdRunner testCommandRunner
	expectedUserName := "Some User"
	userName, err := getGitConfigUserName(cmdRunner)
	if userName != expectedUserName {
		test.Errorf("Expected %v got %v", expectedUserName, userName)
	}
	if err != nil {
		test.Errorf("Got error %v, expected none", err)
	}
}

func TestGetGitConfigUserEmail(test *testing.T) {
	var cmdRunner testCommandRunner
	expectedUserEmail := "me@example.com"
	userEmail, err := getGitConfigUserEmail(cmdRunner)
	if userEmail != expectedUserEmail {
		test.Errorf("Expected %v got %v", expectedUserEmail, userEmail)
	}
	if err != nil {
		test.Errorf("Got error %v, expected none", err)
	}
}

func TestSetGitConfigUserEmail(test *testing.T) {
	var cmdRunner testCommandRunner
	desiredUserEmail := "me@example.com"
	output, err := setGitConfigUserEmail(cmdRunner, desiredUserEmail)
	if output != "" {
		test.Errorf("Expected blank response, got: %v", output)
	}
	if err != nil {
		test.Errorf("Got error %v, expected none", err)
	}
}

func TestSetGitConfigUserName(test *testing.T) {
	var cmdRunner testCommandRunner
	desiredUserName := "Some User"
	output, err := setGitConfigUserName(cmdRunner, desiredUserName)
	if output != "" {
		test.Errorf("Expected blank response, got: %v", output)
	}
	if err != nil {
		test.Errorf("Got error %v, expected none", err)
	}
}
