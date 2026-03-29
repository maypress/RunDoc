package extensions

import (
	"bytes"
	"os/exec"
	"strings"
)

type GoRunner struct{}

func (r GoRunner) Run(code []string) Result {
	cmd := exec.Command("go", "run", "-")
	cmd.Stdin = strings.NewReader(strings.Join(code, "\n"))

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else {
			return Result{Error: err, ExitCode: -1}
		}
	}

	output := stdout.String()
	if stderr.Len() > 0 {
		output += stderr.String()
	}

	return Result{Output: output, ExitCode: exitCode}
}