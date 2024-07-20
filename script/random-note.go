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

// getMdFiles retrieves all .md files in the repository that contain the specified tag (or any tag if none specified)
// and do not contain the draft tag
func getMdFiles(root string, tag string, excludeTag string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".md" {
			containsExcludeTag, err := fileContainsTag(path, excludeTag)
			if err != nil {
				log.Printf("Error reading file %s: %v", path, err)
				return nil
			}
			if containsExcludeTag {
				return nil
			}

			if tag == "" {
				// If no tag is specified, include any file without the exclude tag
				files = append(files, path)
				return nil
			}

			containsTag, err := fileContainsTag(path, tag)
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

// fileContainsTag checks if a file contains the specified tag
func fileContainsTag(filePath string, tag string) (bool, error) {
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
	var tag string
	if len(os.Args) > 1 {
		tag = os.Args[1]
	}
	excludeTag := "#draft"
	rand.Seed(time.Now().UnixNano())
	root := "."

	mdFiles, err := getMdFiles(root, tag, excludeTag)
	if err != nil {
		log.Fatalf("Failed to get .md files: %v", err)
	}
	if len(mdFiles) == 0 {
		log.Fatalf("No valid .md files found with %s tag and without %s tag", tag, excludeTag)
	}

	randomFile := mdFiles[rand.Intn(len(mdFiles))]
	fmt.Println(randomFile)
}
