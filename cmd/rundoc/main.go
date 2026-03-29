// Package main - точка входа в приложение RunDoc
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/maypress/RunDoc/internal/parser"
	"github.com/maypress/RunDoc/internal/reporter"
	"github.com/maypress/RunDoc/internal/runner"
	"github.com/maypress/RunDoc/internal/validator"
)

const version = "1.0.0"

func main() {
	update := flag.Bool("update", false, "Обновить ожидаемый вывод в markdown файле")
	verbose := flag.Bool("verbose", false, "Подробный вывод с деталями выполнения")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		printHelp()
		os.Exit(1)
	}

	filePath := args[0]

	if *verbose {
		printHeader()
	}

	fmt.Printf("📄 Файл: %s\n", filePath)
	if *update {
		fmt.Println("🔧 Режим: ОБНОВЛЕНИЕ")
	}
	if *verbose {
		fmt.Println("🔧 Режим: ПОДРОБНЫЙ")
	}

	fmt.Println("\n⏳ Запуск проверки...")

	// Парсинг файла
	result, err := parser.Parse(filePath)
	if err != nil {
		fmt.Printf("❌ Ошибка парсинга: %v\n", err)
		os.Exit(1)
	}

	if len(result.Blocks) == 0 {
		fmt.Println("⚠️ Исполняемых блоков кода не найдено")
		return
	}

	// Выполнение и валидация блоков
	var reportResults []reporter.ReportResult

	for _, block := range result.Blocks {
		start := time.Now()
		
		// Получение runner для языка
		runnerImpl, err := runner.GetRunner(block.Language)
		if err != nil {
			reportResults = append(reportResults, reporter.ReportResult{
				Block:    block,
				Error:    err,
				Duration: time.Since(start),
			})
			continue
		}

		// Выполнение кода
		runResult := runnerImpl.Run(block.Code)

		// Валидация результата
		err = validator.Validate(block, runResult)

		reportResults = append(reportResults, reporter.ReportResult{
			Block:    block,
			Output:   runResult.Output,
			ExitCode: runResult.ExitCode,
			Error:    err,
			Duration: time.Since(start),
		})
	}

	// Вывод отчета
	if *verbose {
		reporter.PrintVerbose(reportResults, filePath)
	} else {
		reporter.Print(reportResults, filePath)
	}

	if *update {
		fmt.Println("\n⚠️ Режим обновления: функционал в разработке")
	}

	// Выход с правильным кодом
	failedCount := 0
	for _, r := range reportResults {
		if r.Error != nil {
			failedCount++
		}
	}
	if failedCount > 0 {
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("Использование: rundoc [--update] [--verbose] <file.md>")
	fmt.Println("\nФлаги:")
	fmt.Println("  --update    Обновить ожидаемый вывод в markdown файле")
	fmt.Println("  --verbose   Подробный вывод с деталями выполнения")
	fmt.Println("\nПримеры:")
	fmt.Println("  rundoc README.md")
	fmt.Println("  rundoc --update README.md")
	fmt.Println("  rundoc --verbose README.md")
}

func printHeader() {
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Printf("║     RunDoc - Executable Documentation v%s     ║\n", version)
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Println()
}