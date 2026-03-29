// python runner

package extensions

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/maypress/RunDoc/internal/runner"
)

type PythonRunner struct{}

func (r PythonRunner) Run(code []string) runner.Result {
	cmd := exec.Command("python", "-c", strings.Join(code, "\n"))

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else {
			return runner.Result{Error: err, ExitCode: -1}
		}
	}

	output := stdout.String()
	if stderr.Len() > 0 {
		output += stderr.String()
	}

	return runner.Result{Output: output, ExitCode: exitCode}
}