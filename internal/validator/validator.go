package validator

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/maypress/RunDoc/internal/parser"
	"github.com/maypress/RunDoc/internal/runner/extensions"
)

// normalizeOutput убирает лишние символы из вывода
func normalizeOutput(output string) string {
	// Убираем экранирование кавычек
	output = strings.ReplaceAll(output, `\"`, `"`)
	// Убираем лишние пробелы в начале и конце
	output = strings.TrimSpace(output)
	// Убираем символы возврата каретки
	output = strings.ReplaceAll(output, "\r", "")
	return output
}

// Validate проверяет результат выполнения на соответствие ожиданиям
func Validate(block parser.CodeBlock, result extensions.Result) error {
	// Проверка exit code
	if block.ExpectExit != 0 {
		if result.ExitCode != block.ExpectExit {
			return fmt.Errorf("exit code mismatch: expected %d, got %d", block.ExpectExit, result.ExitCode)
		}
	}

	// Нормализуем фактический вывод
	actual := normalizeOutput(result.Output)

	// Проверка регулярного выражения
	if block.ExpectRegex != "" {
		matched, err := regexp.MatchString(block.ExpectRegex, actual)
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
		expected = normalizeOutput(expected)
		
		if expected != actual {
			return fmt.Errorf("output mismatch:\nExpected:\n%s\nGot:\n%s", expected, actual)
		}
	}

	return nil
}