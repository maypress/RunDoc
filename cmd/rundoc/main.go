package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/maypress/RunDoc/internal/parser"
)

const version = "1.0.0"

func main() {
	// Определяем флаги
	update := flag.Bool("update", false, "Обновить ожидаемый вывод в markdown файле")
	verbose := flag.Bool("verbose", false, "Подробный вывод с деталями выполнения")
	flag.Parse()

	// Проверяем что передан файл
	args := flag.Args()
	if len(args) < 1 {
		printHelp()
		os.Exit(1)
	}

	filePath := args[0]

	// Показываем шапку в подробном режиме
	if *verbose {
		printHeader()
	}

	// Выводим информацию о запуске
	fmt.Printf("📄 Файл: %s\n", filePath)
	if *update {
		fmt.Println("🔧 Режим: ОБНОВЛЕНИЕ")
	}
	if *verbose {
		fmt.Println("🔧 Режим: ПОДРОБНЫЙ")
	}
	
	fmt.Println("\n⏳ Запуск проверки...")
	
	fmt.Println("\n✅ RunDoc готов к работе!")
	
	// TODO: Здесь будет основная логика
	// var testFilePath string = "../../testdata/sample.md";
	parser.Parse(filePath)
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