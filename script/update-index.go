package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type fileInfo struct {
	path string
	info os.FileInfo
}

var ignoreTags = []string{"#draft", "#ignore", "#cv", "#aboutme"}

func main() {
	root := "."
	var files []fileInfo

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// игнорируем директории, содержащие "draft" в пути
		if strings.Contains(path, string(filepath.Separator)+"draft"+string(filepath.Separator)) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == ".md" && info.Name() != "index.md" {
			if !containsIgnoredTags(path) {
				files = append(files, fileInfo{path, info})
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the path %q: %v\n", root, err)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].info.ModTime().After(files[j].info.ModTime())
	})

	if len(files) > 20 {
		files = files[:20]
	}

	indexFilePath := filepath.Join(root, "index.md")
	content, err := os.ReadFile(indexFilePath)
	if err != nil {
		log.Fatalf("Error reading index.md file: %v\n", err)
	}

	var newContent strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	inSection := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "## Последние заметки" {
			inSection = true
			newContent.WriteString(line + "\n")
			newContent.WriteString(generateNewSection(files))
		} else if inSection && strings.HasPrefix(line, "- [") {
			continue
		} else {
			newContent.WriteString(line + "\n")
			inSection = false
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning index.md file: %v\n", err)
	}

	err = os.WriteFile(indexFilePath, []byte(newContent.String()), 0644)
	if err != nil {
		log.Fatalf("Error writing to index.md file: %v\n", err)
	}

	fmt.Println("index.md updated successfully")
}

func containsIgnoredTags(filePath string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file %s: %v\n", filePath, err)
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, tag := range ignoreTags {
			if strings.Contains(scanner.Text(), tag) {
				return true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error scanning file %s: %v\n", filePath, err)
	}

	return false
}

func generateNewSection(files []fileInfo) string {
	var section strings.Builder
	for _, file := range files {
		name := strings.TrimSuffix(filepath.Base(file.path), filepath.Ext(file.path))
		path := sanitizePath(file.path)
		section.WriteString(fmt.Sprintf("- [%s](%s)\n", name, path))
	}
	return section.String()
}

func sanitizePath(path string) string {
	path = strings.ReplaceAll(path, " ", "-")
	path = strings.ReplaceAll(path, "#", "")
	return path
}
