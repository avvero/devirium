package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	rootDir := "./" // Путь к корневой директории с заметками

	err := filepath.Walk(rootDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			processFile(path)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
}

func processFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fileContent []string
	containsDraft := false

	for scanner.Scan() {
		line := scanner.Text()
		fileContent = append(fileContent, line)
		if strings.Contains(line, "#draft") {
			containsDraft = true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if containsDraft {
		fmt.Printf("Skipping file (contains #draft): %s\n", path)
	} else {
		fileContent = append(fileContent, "#draft")
		err := ioutil.WriteFile(path, []byte(strings.Join(fileContent, "\n")), 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
		} else {
			fmt.Printf("Added #draft to file: %s\n", path)
		}
	}
}
