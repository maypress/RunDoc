// Валидация вывода.
// Сравнивает реальный вывод команды с ожидаемым.
// Поддерживает точное совпадение, регулярные выражения и проверку exit code.

package validator

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/maypress/RunDoc/internal/parser"
	"github.com/maypress/RunDoc/internal/runner"
)

// Validate проверяет результат выполнения на соответствие ожиданиям
func Validate(block parser.CodeBlock, result runner.Result) error {
	// Проверка exit code
	if block.ExpectExit != 0 && result.ExitCode != block.ExpectExit {
		return fmt.Errorf("exit code mismatch: expected %d, got %d", block.ExpectExit, result.ExitCode)
	}

	// Проверка регулярного выражения
	if block.ExpectRegex != "" {
		matched, err := regexp.MatchString(block.ExpectRegex, result.Output)
		if err != nil {
			return fmt.Errorf("invalid regex: %w", err)
		}
		if !matched {
			return fmt.Errorf("output does not match regex: %s", block.ExpectRegex)
		}
		return nil
	}

	// Проверка точного совпадения вывода
	if len(block.ExpectOutput) > 0 {
		expected := strings.Join(block.ExpectOutput, "\n")
		actual := strings.TrimRight(result.Output, "\n")
		if expected != actual {
			return fmt.Errorf("output mismatch:\nExpected:\n%s\nGot:\n%s", expected, actual)
		}
	}

	return nil
}