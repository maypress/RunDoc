// Вывод результатов проверки.
// Форматирует и выводит в консоль информацию о прошедших и проваленных блоках.
// Поддерживает обычный, подробный (verbose) и JSON форматы.

package reporter

import (
	"fmt"
	"strings"
	"time"

	"github.com/maypress/RunDoc/internal/parser"
)

type ReportResult struct {
	Block     parser.CodeBlock
	Output    string
	ExitCode  int
	Error     error
	Duration  time.Duration
}

func Print(results []ReportResult, filePath string) {
	passed := 0
	failed := 0

	fmt.Printf("\n📄 %s\n", filePath)

	for _, result := range results {
		// Определяем язык и название
		lang := result.Block.Language
		
		// Получаем первую строку кода как описание (или используем номер)
		description := getDescription(result.Block.Code)
		
		if result.Error == nil {
			passed++
			fmt.Printf("  ✓ %s (%s) — %v\n", description, lang, result.Duration)
		} else {
			failed++
			fmt.Printf("  ✗ %s (%s) — %v\n", description, lang, result.Duration)
			fmt.Printf("    💡 %v\n", result.Error)
			if len(result.Output) > 0 {
				fmt.Printf("    📤 Фактический вывод:\n%s\n", indent(result.Output, "      "))
			}
		}
	}

	fmt.Printf("\n📊 Результат: %d из %d блоков пройдено\n", passed, passed+failed)
	
	if failed > 0 {
		fmt.Printf("💡 Обновите документацию: rundoc %s --update\n", filePath)
	}
}

func getDescription(code []string) string {
	if len(code) == 0 {
		return "пустой блок"
	}
	
	firstLine := strings.TrimSpace(code[0])
	if len(firstLine) > 50 {
		return firstLine[:47] + "..."
	}
	return firstLine
}

func indent(text, indentStr string) string {
	if text == "" {
		return ""
	}
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		lines[i] = indentStr + line
	}
	return strings.Join(lines, "\n")
}

func PrintVerbose(results []ReportResult, filePath string) {
	Print(results, filePath)
	
	fmt.Println("\n--- ПОДРОБНОСТИ ---")
	for i, result := range results {
		fmt.Printf("\n[Блок %d] Язык: %s\n", i+1, result.Block.Language)
		fmt.Printf("Код:\n%s\n", indent(strings.Join(result.Block.Code, "\n"), "  "))
		if result.Output != "" {
			fmt.Printf("Вывод:\n%s\n", indent(result.Output, "  "))
		}
		if result.Error != nil {
			fmt.Printf("Ошибка: %v\n", result.Error)
		}
	}
}