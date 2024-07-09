package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run script.go <path to file>")
		os.Exit(1)
	}
	filePath := os.Args[1]
	root := "./" // Путь к корневой директории

	// Считываем содержимое указанного файла
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %v: %v\n", filePath, err)
		os.Exit(1)
	}

	// Находим все ссылки вида [[имя файла]]
	linkPattern := regexp.MustCompile(`\[\[(.+?)\]\]`)
	matches := linkPattern.FindAllStringSubmatch(string(content), -1)

	if len(matches) == 0 {
		fmt.Println("null")
		return
	}

	// Собираем все файлы в текущей директории, включая поддиректории
	files := make(map[string]string)
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files[strings.ToLower(filepath.Base(path))] = path
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %v: %v\n", root, err)
		os.Exit(1)
	}

	var results []string

	// Проверка наличия всех ссылок и сбор результатов
	for _, match := range matches {
		noteName := match[1] + ".md"
		lowerNoteName := strings.ToLower(noteName)
		if path, exists := files[lowerNoteName]; exists {
			actualFileName := filepath.Base(path)
			if actualFileName != noteName {
				fmt.Printf("Case mismatch for link [[%s]]: expected %s but found %s\n", match[1], noteName, actualFileName)
				os.Exit(1)
			}
			formattedPath := strings.ReplaceAll(strings.TrimSuffix(path, filepath.Ext(path)), " ", "-")
			results = append(results, fmt.Sprintf("[[%s]] - %s", match[1], formattedPath))
		} else {
			fmt.Printf("[[%s]] - File not found\n", match[1])
			os.Exit(1)
		}
	}

	if len(results) == 0 {
		fmt.Println("null")
	} else {
		fmt.Println(strings.Join(results, "\n"))
	}
}
