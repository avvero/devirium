package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type fileInfo struct {
	path string
	info os.FileInfo
}

type tagCount struct {
	tag   string
	count int
}

func main() {
	root := "."
	var files []fileInfo
	tags := make(map[string]int)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == ".md" {
			files = append(files, fileInfo{path, info})
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the path %q: %v\n", root, err)
	}

	tagRegex := regexp.MustCompile(`#\w+`)

	for _, file := range files {
		content, err := os.ReadFile(file.path)
		if err != nil {
			log.Printf("Error reading file %q: %v\n", file.path, err)
			continue
		}

		matches := tagRegex.FindAllString(string(content), -1)
		for _, tag := range matches {
			if tag == "#draft" {
				continue
			}
			tags[tag]++
		}
	}

	var tagCounts []tagCount
	for tag, count := range tags {
		tagCounts = append(tagCounts, tagCount{tag, count})
	}

	sort.Slice(tagCounts, func(i, j int) bool {
		return tagCounts[i].count > tagCounts[j].count
	})

	topTags := tagCounts
	if len(tagCounts) > 5 {
		topTags = tagCounts[:5]
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
		if strings.TrimSpace(line) == "## Популярные тэги" {
			inSection = true
			newContent.WriteString(line + "\n")
			newContent.WriteString(generatePopularTagsSection(topTags))
		} else if inSection && strings.HasPrefix(line, "- #") {
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

func generatePopularTagsSection(topTags []tagCount) string {
	var section strings.Builder
	for _, tag := range topTags {
		section.WriteString(fmt.Sprintf("- %s\n", tag.tag))
	}
	return section.String()
}
