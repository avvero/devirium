package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"net/http"
)

func main() {
	root := "./"
	files := make(map[string]string)

	// Collect all image files in the current directory, including subdirectories
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			contentType := http.DetectContentType(content)
			if strings.HasPrefix(contentType, "image/") {
				files[strings.ToLower(filepath.Base(path))] = path
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %v: %v\n", root, err)
		return
	}

	// Process all markdown files and check image links
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
	linkPattern := regexp.MustCompile(`!\[.*?\]\((.+?\.(png|jpg|jpeg))\)`)
	matches := linkPattern.FindAllStringSubmatch(content, -1)
	updatedContent := content

	for _, match := range matches {
		imageName := match[1]
		lowerImageName := strings.ToLower(imageName)
		if actualPath, exists := files[lowerImageName]; exists {
			// Fix links if necessary
			actualFileName := filepath.Base(actualPath)
			if actualFileName != imageName {
				updatedContent = strings.ReplaceAll(updatedContent, match[0], fmt.Sprintf("![Alt text](%s)", actualFileName))
			}
		} else {
			log.Fatalf("Missing file for link in %s: ![Alt text](%s)", filePath, match[1])
		}
	}

	return updatedContent
}
