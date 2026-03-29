package runner

import (
	"fmt"

	"github.com/maypress/RunDoc/internal/runner/extensions"
)

type Runner interface {
	Run(code []string) extensions.Result
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