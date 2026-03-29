// Package main предоставляет RunDoc - инструмент исполняемой документации
// Парсит markdown файлы, выполняет блоки кода и проверяет результат
package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	appName = "RunDoc"
	version = "1.0.0"
)

// Config содержит настройки приложения из аргументов командной строки
type Config struct {
	Update   bool
	Verbose  bool
	FilePath string
}

// FlagParser отвечает за парсинг флагов командной строки
type FlagParser struct{}

// Parse парсит аргументы командной строки и возвращает конфигурацию
func (fp *FlagParser) Parse() (*Config, error) {
	var config Config

	// Определяем флаги
	flag.BoolVar(&config.Update, "update", false, "Обновить ожидаемый вывод в markdown файле")
	flag.BoolVar(&config.Verbose, "verbose", false, "Подробный вывод с деталями выполнения")

	// Парсим флаги
	flag.Parse()

	// Проверяем обязательные аргументы
	args := flag.Args()
	if len(args) < 1 {
		return nil, fmt.Errorf("не указан путь к файлу")
	}

	config.FilePath = args[0]
	return &config, nil
}

// HelpPrinter отвечает за вывод справочной информации
type HelpPrinter struct{}

// PrintUsage выводит информацию об использовании программы
func (hp *HelpPrinter) PrintUsage() {
	fmt.Printf("Использование: %s [--update] [--verbose] <file.md>\n", appName)
	fmt.Println("\nФлаги:")
	fmt.Println("  --update    Обновить ожидаемый вывод в markdown файле")
	fmt.Println("  --verbose   Подробный вывод с деталями выполнения")
	fmt.Println("\nПримеры:")
	fmt.Println("  rundoc README.md                    # Проверить документацию")
	fmt.Println("  rundoc --update README.md           # Обновить ожидаемые выводы")
	fmt.Println("  rundoc --verbose README.md          # Подробный вывод")
	fmt.Println("  rundoc --update --verbose README.md # Обновить с подробным логом")
}

// Logger отвечает за вывод сообщений в зависимости от режима
type Logger struct {
	verbose bool
}

// NewLogger создает новый логгер
func NewLogger(verbose bool) *Logger {
	return &Logger{verbose: verbose}
}

// Info выводит информационное сообщение
func (l *Logger) Info(message string) {
	if l.verbose {
		fmt.Println(message)
	}
}

// VerboseHeader выводит заголовок в подробном режиме
func (l *Logger) VerboseHeader() {
	if l.verbose {
		fmt.Println("╔════════════════════════════════════════════════════════════╗")
		fmt.Printf("║                 %s - Executable Documentation v%s      ║\n", appName, version)
		fmt.Println("╚════════════════════════════════════════════════════════════╝")
		fmt.Println()
	}
}

// Runner отвечает за выполнение основной логики приложения
type Runner struct {
	logger *Logger
}

// NewRunner создает новый раннер
func NewRunner(logger *Logger) *Runner {
	return &Runner{logger: logger}
}

// Run выполняет основную логику приложения
func (r *Runner) Run(config *Config) {
	r.logger.Info(fmt.Sprintf("📄 Файл: %s", config.FilePath))
	r.logger.Info(fmt.Sprintf("🔧 Режимы: %s %s",
		getModeText(config.Update, "UPDATE"),
		getModeText(config.Verbose, "VERBOSE")))
	r.logger.Info("")

	r.logger.Info("⏳ Запуск проверки...")

	// TODO: Здесь будет основная логика
	// 1. Вызов парсера: parser.Parse(config.FilePath)
	// 2. Для каждого блока: runner.GetRunner(language)
	// 3. Выполнение: runner.Run(code)
	// 4. Валидация: validator.Validate(block, result)
	// 5. Вывод: reporter.Report(results, config.Verbose)
	// 6. Если config.Update: обновить файл с новыми ожиданиями
}

// getModeText возвращает текст режима если он активен
func getModeText(active bool, text string) string {
	if active {
		return "[" + text + "]"
	}
	return ""
}

// Application главный класс приложения
type Application struct {
	flagParser  *FlagParser
	helpPrinter *HelpPrinter
}

// NewApplication создает новое приложение
func NewApplication() *Application {
	return &Application{
		flagParser:  &FlagParser{},
		helpPrinter: &HelpPrinter{},
	}
}

// Execute запускает приложение
func (app *Application) Execute() int {
	// Парсим флаги
	config, err := app.flagParser.Parse()
	if err != nil {
		app.helpPrinter.PrintUsage()
		return 1
	}

	// Создаем логгер
	logger := NewLogger(config.Verbose)
	logger.VerboseHeader()

	// Выводим информацию о запуске
	if !config.Verbose {
		fmt.Printf("Processing: %s", config.FilePath)
		if config.Update {
			fmt.Print(" (update mode)")
		}
		fmt.Println()
	}

	// Создаем и запускаем раннер
	runner := NewRunner(logger)
	runner.Run(config)

	// Временный вывод
	fmt.Println("✅ RunDoc готов к работе!")
	fmt.Println("📝 Реализация парсера, раннеров и валидатора в процессе...")

	logger.Info("✨ Для подробной информации используйте --verbose")

	return 0
}

// main точка входа в программу
func main() {
	app := NewApplication()
	os.Exit(app.Execute())
}
