package main

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func main() {
	tags := make(map[string]struct{})
	tagRegex := regexp.MustCompile(`#\w+`)
	urlRegex := regexp.MustCompile(`https?://[^\s]+`)

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".md" {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				text := scanner.Text()
				// Проверяем, не является ли строка URL
				if !urlRegex.MatchString(text) {
					matches := tagRegex.FindAllString(text, -1)
					for _, match := range matches {
						// Добавляем тег, если он не часть URL
						tag := strings.ToLower(match)
						tags[tag] = struct{}{}
					}
				}
			}
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	var tagsSlice []string
	for tag := range tags {
		tagsSlice = append(tagsSlice, tag)
	}
	sort.Strings(tagsSlice)

	file, err := os.OpenFile("_Тэги.md", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString("\n\n")
	if err != nil {
		panic(err)
	}

	for _, tag := range tagsSlice {
		_, err = file.WriteString(tag + "\n")
		if err != nil {
			panic(err)
		}
	}
}
