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

func getMdFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".md" {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func containsDraftTag(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "#draft") {
			return true, nil
		}
	}
	return false, scanner.Err()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	root := "."

	mdFiles, err := getMdFiles(root)
	if err != nil {
		log.Fatalf("Failed to get .md files: %v", err)
	}
	if len(mdFiles) == 0 {
		log.Fatalf("No .md files found")
	}

	for {
		if len(mdFiles) == 0 {
			log.Fatalf("No valid .md files found without #draft tag")
		}
		randomIndex := rand.Intn(len(mdFiles))
		randomFile := mdFiles[randomIndex]
		isDraft, err := containsDraftTag(randomFile)
		if err != nil {
			log.Printf("Error reading file %s: %v", randomFile, err)
			mdFiles = append(mdFiles[:randomIndex], mdFiles[randomIndex+1:]...)
			continue
		}
		if !isDraft {
			fmt.Println(randomFile)
			break
		}
		mdFiles = append(mdFiles[:randomIndex], mdFiles[randomIndex+1:]...)
	}
}
