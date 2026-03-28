// RunDoc - исполнимая документация
// Точка входа в программу.
// Парсит аргументы CLI, вызывает парсер, раннеры и валидатор,
// выводит результат через reporter.

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Определяем флаги командной строки
	updateFlag := flag.Bool("update", false, "Обновить ожидаемый вывод в markdown файле")
	verboseFlag := flag.Bool("verbose", false, "Подробный вывод с деталями выполнения")

	// Парсим флаги
	flag.Parse()

	// Получаем аргумент (путь к файлу)
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Использование: rundoc [--update] [--verbose] <file.md>")
		fmt.Println("\nФлаги:")
		fmt.Println("  --update    Обновить ожидаемый вывод в markdown файле")
		fmt.Println("  --verbose   Подробный вывод с деталями выполнения")
		fmt.Println("\nПримеры:")
		fmt.Println("  rundoc README.md                    # Проверить документацию")
		fmt.Println("  rundoc --update README.md           # Обновить ожидаемые выводы")
		fmt.Println("  rundoc --verbose README.md          # Подробный вывод")
		fmt.Println("  rundoc --update --verbose README.md # Обновить с подробным логом")
		os.Exit(1)
	}

	filePath := args[0]
	updateMode := *updateFlag
	verboseMode := *verboseFlag

	// Выводим информацию о режиме работы
	if verboseMode {
		fmt.Println("╔════════════════════════════════════════════════════════════╗")
		fmt.Println("║                    RunDoc - Executable Documentation      ║")
		fmt.Println("╚════════════════════════════════════════════════════════════╝")
		fmt.Println()
		fmt.Printf("📄 Файл: %s\n", filePath)
		fmt.Printf("🔧 Режимы: ")
		if updateMode {
			fmt.Printf("[UPDATE] ")
		}
		if verboseMode {
			fmt.Printf("[VERBOSE]")
		}
		fmt.Println("\n")
	} else {
		fmt.Printf("Processing: %s", filePath)
		if updateMode {
			fmt.Print(" (update mode)")
		}
		fmt.Println()
	}

	// TODO: Здесь будет основная логика
	// 1. Вызов парсера: parser.Parse(filePath)
	// 2. Для каждого блока: runner.GetRunner(language)
	// 3. Выполнение: runner.Run(code)
	// 4. Валидация: validator.Validate(block, result)
	// 5. Вывод: reporter.Report(results, verboseMode)
	// 6. Если updateMode: обновить файл с новыми ожиданиями

	if verboseMode {
		fmt.Println("\n⏳ Запуск проверки...")
		fmt.Println()
	}

	// Временный вывод (потом убрать)
	fmt.Println("✅ RunDoc готов к работе!")
	fmt.Println("📝 Реализация парсера, раннеров и валидатора в процессе...")

	if verboseMode {
		fmt.Println("\n✨ Для подробной информации используйте --verbose")
	}
}
