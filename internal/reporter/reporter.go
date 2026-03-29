// Вывод результатов проверки.
// Форматирует и выводит в консоль информацию о прошедших и проваленных блоках.
// Поддерживает обычный, подробный (verbose) и JSON форматы.

package reporter

import (
	"fmt"
	"strings"

	"github.com/maypress/RunDoc/internal/parser"
)

// ReportResult содержит результаты выполнения блоков кода
type ReportResult struct {
	Block    parser.CodeBlock
	Output   string
	ExitCode int
	Error    error
}

// Print выводит отчет о проверке всех блоков
func Print(results []ReportResult, verbose bool) {
	passed := 0
	failed := 0

	for i, result := range results {
		if result.Error == nil {
			passed++
			if verbose {
				fmt.Printf("✅ Блок %d: УСПЕШНО\n", i+1)
			}
		} else {
			failed++
			fmt.Printf("❌ Блок %d: ОШИБКА - %v\n", i+1, result.Error)
		}
	}

	fmt.Printf("\n📊 ИТОГО: Успешно=%d, Ошибок=%d\n", passed, failed)
}

// PrintVerbose выводит детальный отчет о проверке
func PrintVerbose(results []ReportResult) {
	Print(results, true)

	for i, result := range results {
		fmt.Printf("\n--- Детали блока %d ---\n", i+1)
		fmt.Printf("Язык: %s\n", result.Block.Language)
		fmt.Printf("Код:\n%s\n", strings.Join(result.Block.Code, "\n"))
		fmt.Printf("Вывод:\n%s\n", result.Output)
		fmt.Printf("Exit code: %d\n", result.ExitCode)
	}
}