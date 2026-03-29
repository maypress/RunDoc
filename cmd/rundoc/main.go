// Package main - точка входа в приложение RunDoc
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/maypress/RunDoc/internal/parser"
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

	result, err := parser.Parse(filePath)
	if err != nil {
		fmt.Printf("❌ Ошибка: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ Найдено блоков кода: %d\n", len(result.Blocks))

	for i, block := range result.Blocks {
		fmt.Printf("  Блок %d: язык=%s, строк кода=%d\n", i+1, block.Language, len(block.Code))
	}
}

// printHelp выводит справочную информацию об использовании программы
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

// printHeader выводит шапку программы в verbose режиме
func printHeader() {
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Printf("║     RunDoc - Executable Documentation v%s     ║\n", version)
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Println()
}