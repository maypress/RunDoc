package extensions

import (
	"bytes"
	"os/exec"
	"runtime"
	"strings"
)

type BashRunner struct{}

type Result struct {
	Output   string
	ExitCode int
	Error    error
}

func (r BashRunner) Run(code []string) Result {
	script := strings.Join(code, "\n")
	
	var cmd *exec.Cmd
	
	if runtime.GOOS == "windows" {
		// Для Windows используем echo без кавычек
		script = strings.ReplaceAll(script, `"`, "")
		cmd = exec.Command("cmd", "/c", script)
	} else {
		cmd = exec.Command("bash", "-c", script)
	}

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
	
	// Нормализуем вывод
	output = strings.TrimSpace(output)
	output = strings.ReplaceAll(output, "\r", "")

	return Result{Output: output, ExitCode: exitCode}
}