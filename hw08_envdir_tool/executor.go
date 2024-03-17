package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	if len(cmd) == 0 {
		return 0
	}

	command := exec.Command(cmd[0], cmd[1:]...) //nolint:gosec
	command.Env = os.Environ()

	for key, value := range env {
		if !value.NeedRemove {
			environ := fmt.Sprintf("%v=%v", key, value.Value)
			command.Env = append(command.Env, environ)
		} else {
			command.Env = append(command.Env, fmt.Sprintf("%v=", key))
		}
	}

	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		var ee *exec.ExitError
		var e *exec.Error
		if errors.As(err, &ee) {
			return ee.ExitCode()
		}
		if errors.As(err, &e) {
			return 127
		}
	}

	return 0
}
