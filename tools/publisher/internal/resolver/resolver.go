package resolver

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	linkPattern  = regexp.MustCompile(`\[\[(.+?)\]\]`)
	imagePattern = regexp.MustCompile(`!\[.*?\]\((.+?\.(?:png|jpg|jpeg))\)`)
)

// Index maps lowercase basename -> real path (relative to root).
type Index struct {
	files map[string]string
}

// BuildIndex walks the given root and indexes every file by lowercase basename.
func BuildIndex(root string) (*Index, error) {
	idx := &Index{files: map[string]string{}}
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		rel, relErr := filepath.Rel(root, path)
		if relErr != nil {
			rel = path
		}
		idx.files[strings.ToLower(filepath.Base(path))] = filepath.ToSlash(rel)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return idx, nil
}

// ResolveLinks returns wikilink name -> resolved path (extension stripped, spaces -> dashes).
// Errors on case mismatch or missing file.
func (i *Index) ResolveLinks(content string) (map[string]string, error) {
	matches := linkPattern.FindAllStringSubmatch(content, -1)
	out := map[string]string{}
	for _, m := range matches {
		name := m[1]
		note := name + ".md"
		key := strings.ToLower(note)
		path, ok := i.files[key]
		if !ok {
			return nil, fmt.Errorf("[[%s]] - file not found", name)
		}
		actual := filepath.Base(path)
		if actual != note {
			return nil, fmt.Errorf("case mismatch for link [[%s]]: expected %s but found %s", name, note, actual)
		}
		trimmed := strings.TrimSuffix(path, filepath.Ext(path))
		out[name] = strings.ReplaceAll(trimmed, " ", "-")
	}
	return out, nil
}

// ResolveImages returns image ref -> resolved path (spaces -> dashes).
func (i *Index) ResolveImages(content string) (map[string]string, error) {
	matches := imagePattern.FindAllStringSubmatch(content, -1)
	out := map[string]string{}
	for _, m := range matches {
		name := m[1]
		key := strings.ToLower(name)
		path, ok := i.files[key]
		if !ok {
			return nil, fmt.Errorf("[[%s]] - file not found", name)
		}
		actual := filepath.Base(path)
		if actual != name {
			return nil, fmt.Errorf("case mismatch for image %s: expected %s but found %s", name, name, actual)
		}
		out[name] = strings.ReplaceAll(path, " ", "-")
	}
	return out, nil
}
