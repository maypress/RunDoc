// Парсер Markdown файлов.
// Читает .md файл, находит блоки кода с аннотацией "run",
// извлекает код и ожидаемый вывод (# expect:, # expect-regex:, # expect-exit:).
// Возвращает список CodeBlock для дальнейшего выполнения.

package parser

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var allowedExtensions = []string{".md", ".txt", ".json"}

// isExtensionAllowed проверяет, разрешено ли расширение файла
func isExtensionAllowed(ext string) bool {
	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			return true
		}
	}
	return false
}

// CodeBlock представляет один исполняемый блок кода из документации
type CodeBlock struct {
	Language     string
	Code         []string
	ExpectOutput []string
	ExpectRegex  string
	ExpectExit   int
}

// ParseResult содержит результат парсинга файла
type ParseResult struct {
	Blocks   []CodeBlock
	FilePath string
}

// Parse анализирует markdown файл и извлекает исполняемые блоки кода
func Parse(filePath string) (*ParseResult, error) {
	ext := filepath.Ext(filePath)
	if !isExtensionAllowed(ext) {
		return nil, fmt.Errorf("unsupported file extension: %s", ext)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var blocks []CodeBlock
	scanner := bufio.NewScanner(file)

	var currentBlock *CodeBlock
	var inCodeBlock bool

	for scanner.Scan() {
		line := scanner.Text()

		// Начало блока кода
		if strings.HasPrefix(line, "```") && !inCodeBlock {
			parts := strings.Fields(strings.TrimPrefix(line, "```"))
			if len(parts) >= 2 && parts[1] == "run" {
				currentBlock = &CodeBlock{
					Language: parts[0],
					Code:     []string{},
				}
				inCodeBlock = true
			}
			continue
		}

		// Конец блока кода
		if strings.HasPrefix(line, "```") && inCodeBlock {
			blocks = append(blocks, *currentBlock)
			currentBlock = nil
			inCodeBlock = false
			continue
		}

		// Сбор кода внутри блока
		if inCodeBlock && currentBlock != nil {
			// Проверка на ожидаемый вывод
			if strings.HasPrefix(line, "# expect:") {
				expected := strings.TrimPrefix(line, "# expect:")
				currentBlock.ExpectOutput = append(currentBlock.ExpectOutput, strings.TrimSpace(expected))
			} else if strings.HasPrefix(line, "# expect-regex:") {
				currentBlock.ExpectRegex = strings.TrimPrefix(line, "# expect-regex:")
			} else if strings.HasPrefix(line, "# expect-exit:") {
				var exitCode int
				fmt.Sscanf(strings.TrimPrefix(line, "# expect-exit:"), "%d", &exitCode)
				currentBlock.ExpectExit = exitCode
			} else {
				currentBlock.Code = append(currentBlock.Code, line)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return &ParseResult{
		Blocks:   blocks,
		FilePath: filePath,
	}, nil
}