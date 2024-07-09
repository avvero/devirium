package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	root := "./" // Путь к корневой директории с заметками
	files := make(map[string]bool)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".md" {
			files[filepath.Base(path)] = true
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
			findAndCheckLinks(path, string(content), files)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %v: %v\n", root, err)
	}
}

func findAndCheckLinks(filePath, content string, files map[string]bool) {
	linkPattern := regexp.MustCompile(`\[\[(.+?)\]\]`)
	matches := linkPattern.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		noteName := match[1] + ".md"
		if _, exists := files[noteName]; !exists {
			fmt.Printf("Broken link in file %s: [[%s]]\n", filePath, match[1])
		}
	}
}
