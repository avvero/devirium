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

func main() {
	root := "."
	var files []fileInfo

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == ".md" && info.Name() != "index.md" {
			files = append(files, fileInfo{path, info})
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
			newContent.WriteString(generateNewSection(files, root))
		} else if inSection && strings.HasPrefix(line, "- [[") {
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

func generateNewSection(files []fileInfo, root string) string {
	var section strings.Builder
	for _, file := range files {
		relPath, err := filepath.Rel(root, file.path)
		if err != nil {
			log.Fatalf("Error getting relative path: %v\n", err)
		}
		section.WriteString(fmt.Sprintf("- [[%s]]\n", relPath))
	}
	return section.String()
}
