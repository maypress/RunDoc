// Парсер Markdown файлов.
// Читает .md файл, находит блоки кода с аннотацией "run",
// извлекает код и ожидаемый вывод (# expect:, # expect-regex:, # expect-exit:).
// Возвращает список CodeBlock для дальнейшего выполнения.

package parser

import (
	"fmt"
	"path/filepath"
)

// В Go НЕЛЬЗЯ сделать константу-массив, поэтому используем переменную
var ALLOWED_FILE_EXTENTIONS = []string{".md", ".txt", ".json"}

func isExtensionAllowed(ext string) bool {
	// Используем range вместо C-стиля цикла
	for _, allowedExt := range ALLOWED_FILE_EXTENTIONS {
		if ext == allowedExt {
			return true
		}
	}
	return false
}

// Основная функция, определяющая последовательность
func Parse(file string) error {
	// Убираем fmt.Println, парсер не должен ничего выводить
	// fmt.Println("parse: file=" + file);

	fileExtension := filepath.Ext(file)
	isFileExtensionAllowed := isExtensionAllowed(fileExtension)
	
	// Убираем отладочный вывод
	// fmt.Printf("parse: fileExtension= %s isFileExtentisonAllowed=%t", 
	// 				fileExtension, isFileExtensionAllowed);

	if !isFileExtensionAllowed {
		return fmt.Errorf("unsupported file extension: %s", fileExtension)
	}

	switch fileExtension {
		case ".md":
				fmt.Println("fileExtensions md parser")
				return nil

		default:
				fmt.Printf("Other file type: %s", fileExtension)
				return nil
	}

	// Реализация парсирования данных исходя из формата, через функции из других файлов. Условно - /types/shell.go, types/go.go, types/python.go, просто вызов функции + проверка расширения файлов через switch
	// Parse(file) 
	// -> определяет тип 
	// -> вызывает конкретный парсер 
	// -> получает структуру с блоками кода 
	// -> возвращает ее вызывающему коду
	return nil
}