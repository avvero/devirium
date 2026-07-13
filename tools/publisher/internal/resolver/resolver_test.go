package resolver

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeFile(t *testing.T, path, body string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

func setupRepo(t *testing.T) string {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "2024", "Note One.md"), "body")
	writeFile(t, filepath.Join(root, "2024", "Sub", "Deep Note.md"), "body")
	writeFile(t, filepath.Join(root, "img", "pic one.png"), "x")
	writeFile(t, filepath.Join(root, "img", "shot.jpg"), "x")
	return root
}

func TestResolveLinks(t *testing.T) {
	root := setupRepo(t)
	idx, err := BuildIndex(root)
	if err != nil {
		t.Fatal(err)
	}
	links, err := idx.ResolveLinks("see [[Note One]] and [[Deep Note]]")
	if err != nil {
		t.Fatal(err)
	}
	if links["Note One"] != "2024/Note-One" {
		t.Errorf("got %q", links["Note One"])
	}
	if links["Deep Note"] != "2024/Sub/Deep-Note" {
		t.Errorf("got %q", links["Deep Note"])
	}
}

func TestResolveLinksMissing(t *testing.T) {
	root := setupRepo(t)
	idx, _ := BuildIndex(root)
	_, err := idx.ResolveLinks("see [[Missing]]")
	if err == nil || !strings.Contains(err.Error(), "not found") {
		t.Fatalf("want not-found error, got %v", err)
	}
}

func TestResolveLinksCaseMismatch(t *testing.T) {
	root := setupRepo(t)
	idx, _ := BuildIndex(root)
	_, err := idx.ResolveLinks("see [[note one]]")
	if err == nil || !strings.Contains(err.Error(), "case mismatch") {
		t.Fatalf("want case-mismatch error, got %v", err)
	}
}

func TestResolveImages(t *testing.T) {
	root := setupRepo(t)
	idx, _ := BuildIndex(root)
	imgs, err := idx.ResolveImages("hi ![](pic one.png) and ![alt](shot.jpg)")
	if err != nil {
		t.Fatal(err)
	}
	if imgs["pic one.png"] != "img/pic-one.png" {
		t.Errorf("got %q", imgs["pic one.png"])
	}
	if imgs["shot.jpg"] != "img/shot.jpg" {
		t.Errorf("got %q", imgs["shot.jpg"])
	}
}

func TestResolveImagesEmpty(t *testing.T) {
	root := setupRepo(t)
	idx, _ := BuildIndex(root)
	imgs, err := idx.ResolveImages("no images here")
	if err != nil {
		t.Fatal(err)
	}
	if len(imgs) != 0 {
		t.Errorf("want empty, got %v", imgs)
	}
}
