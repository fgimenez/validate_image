package cli

import "os/exec"

var (
	execCommand = exec.Command
)

// Commander comprises the methods required by a generic command executor
type Commander interface {
	ExecCommand(...string) (output string, err error)
}

// Executor is a concrete type for CLI execution
type Executor struct{}

// ExecCommand sends the given command to the CLI and returns the output and
// the resulting error
func (e *Executor) ExecCommand(cmds ...string) (output string, err error) {
	cmd := execCommand(cmds[0], cmds[1:]...)
	outputByte, err := cmd.CombinedOutput()
	output = string(outputByte)
	return
}
