package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// getMdFiles retrieves all .md files in the repository that contain the specified tag
func getMdFiles(root string, tag string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".md" {
			containsTag, err := containsTag(path, tag)
			if err != nil {
				log.Printf("Error reading file %s: %v", path, err)
				return nil
			}
			if containsTag {
				files = append(files, path)
			}
		}
		return nil
	})
	return files, err
}

// containsTag checks if a file contains the specified tag
func containsTag(filePath string, tag string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), tag) {
			return true, nil
		}
	}
	return false, scanner.Err()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	root := "."
	excludeTag := "#debug"

	mdFiles, err := getMdFiles(root, excludeTag)
	if err != nil {
		log.Fatalf("Failed to get .md files: %v", err)
	}
	if len(mdFiles) == 0 {
		log.Fatalf("No valid .md files found without %s tag", excludeTag)
	}

	randomFile := mdFiles[rand.Intn(len(mdFiles))]
	fmt.Println(randomFile)
}
