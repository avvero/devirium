package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	root := "./" // Path to the root directory

	// Read the contents of the specified file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %v: %v\n", filePath, err)
		os.Exit(1)
	}

	// Find all image links
	linkPattern := regexp.MustCompile(`!\[.*?\]\((.+?\.(png|jpg|jpeg))\)`)
	matches := linkPattern.FindAllStringSubmatch(string(content), -1)

	if len(matches) == 0 {
		fmt.Println("{}")
		return
	}

	// Collect all files in the current directory, including subdirectories
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

	results := make(map[string]string)

	// Check for the presence of all links and collect results
	for _, match := range matches {
		imageName := match[1]
		lowerImageName := strings.ToLower(imageName)
		if path, exists := files[lowerImageName]; exists {
			actualFileName := filepath.Base(path)
			if actualFileName != imageName {
				fmt.Printf("Case mismatch for link ![Alt text](%s): expected %s but found %s\n", match[1], imageName, actualFileName)
				os.Exit(1)
			}
			formattedPath := strings.ReplaceAll(path, " ", "-")
			results[match[1]] = formattedPath
		} else {
			log.Fatalf("[[%s]] - File not found\n", match[1])
		}
	}

	jsonResult, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonResult))
}
