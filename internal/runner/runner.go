package runner

import (
	"fmt"

	"github.com/maypress/RunDoc/internal/runner/extensions"
)

type Result struct {
	Output   string
	ExitCode int
	Error    error
}

type Runner interface {
	Run(code []string) Result
}

func GetRunner(language string) (Runner, error) {
	switch language {
	case "bash", "sh":
		return extensions.BashRunner{}, nil
	case "go":
		return extensions.GoRunner{}, nil
	case "python", "py":
		return extensions.PythonRunner{}, nil
	default:
		return nil, fmt.Errorf("unsupported language: %s", language)
	}
}