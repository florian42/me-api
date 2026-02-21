package cmd

import "os/exec"

type CommandRunner interface {
	Run(name string, args ...string) ([]byte, error)
}
type realRunner struct{}

func (r realRunner) Run(name string, args ...string) ([]byte, error) {
	return exec.Command(name, args...).Output()
}

func Runner() CommandRunner {
	return realRunner{}
}
