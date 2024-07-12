package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	// Create a logger
	logger := log.New(os.Stdout, "INFO: ", log.LstdFlags)

	// Define the regex pattern for lowercase English and Russian letters
	pattern := `^[a-zа-я]`

	// Compile the regex
	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatalf("Failed to compile regex: %v", err)
	}

	// Walk through the current directory and its subdirectories
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if it's a file, has .md extension, and matches the regex
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") && re.MatchString(info.Name()) {
			logger.Println(path)
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the path: %v", err)
	}
}
