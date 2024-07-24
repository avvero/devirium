package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	root := "./"
	files := make(map[string]string)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".md" {
			files[strings.ToLower(filepath.Base(path))] = path
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %v: %v\n", root, err)
		return
	}

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".md" {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			updatedContent := findAndFixLinks(path, string(content), files)
			if updatedContent != string(content) {
				err = ioutil.WriteFile(path, []byte(updatedContent), 0644)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %v: %v\n", root, err)
	}
}

func findAndFixLinks(filePath, content string, files map[string]string) string {
	linkPattern := regexp.MustCompile(`\[\[(.+?)\]\]`)
	matches := linkPattern.FindAllStringSubmatch(content, -1)
	updatedContent := content

	for _, match := range matches {
		noteName := match[1] + ".md"
		lowerNoteName := strings.ToLower(noteName)
		if actualPath, exists := files[lowerNoteName]; exists {
			// Исправляем ссылки, если необходимо
			actualFileName := filepath.Base(actualPath)
			if actualFileName != noteName {
				updatedContent = strings.ReplaceAll(updatedContent, match[0], fmt.Sprintf("[[%s]]", strings.TrimSuffix(actualFileName, ".md")))
			}
		} else {
			log.Fatalf("Missing file for link in %s: [[%s]]", filePath, match[1])

			// Создаем недостающий файл в папке "missing"
			// newFilePath := filepath.Join(missingDir, noteName)
			// err := ioutil.WriteFile(newFilePath, []byte("# "+strings.TrimSuffix(noteName, ".md")), 0644)
			// if err != nil {
			// 	fmt.Printf("Error creating file %v: %v\n", newFilePath, err)
			// 	continue
			// }
			// files[lowerNoteName] = newFilePath
			// fmt.Printf("Created missing file: %s\n", newFilePath)
		}
	}

	return updatedContent
}
