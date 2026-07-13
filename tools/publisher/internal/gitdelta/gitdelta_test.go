package gitdelta

import (
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"testing"
)

func gitInit(t *testing.T) string {
	t.Helper()
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git not available")
	}
	dir := t.TempDir()
	runIn(t, dir, "git", "init", "-q", "-b", "main")
	runIn(t, dir, "git", "config", "user.email", "t@t")
	runIn(t, dir, "git", "config", "user.name", "t")
	runIn(t, dir, "git", "config", "commit.gpgsign", "false")
	return dir
}

func runIn(t *testing.T, dir, name string, args ...string) {
	t.Helper()
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("%s %v: %v: %s", name, args, err, string(out))
	}
}

func write(t *testing.T, root, rel, body string) {
	t.Helper()
	full := filepath.Join(root, rel)
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(full, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestChangedMarkdown(t *testing.T) {
	repo := gitInit(t)
	write(t, repo, "a.md", "one")
	write(t, repo, "keep.txt", "x")
	runIn(t, repo, "git", "add", ".")
	runIn(t, repo, "git", "commit", "-q", "-m", "init")

	write(t, repo, "b.md", "two")
	write(t, repo, "a.md", "one v2")
	runIn(t, repo, "git", "add", ".")
	runIn(t, repo, "git", "commit", "-q", "-m", "second")

	files, err := ChangedMarkdown(repo, "HEAD~1", "HEAD")
	if err != nil {
		t.Fatal(err)
	}
	sort.Strings(files)
	if len(files) != 2 || files[0] != "a.md" || files[1] != "b.md" {
		t.Errorf("got %v", files)
	}
}

func TestFileAt(t *testing.T) {
	repo := gitInit(t)
	write(t, repo, "a.md", "hello")
	runIn(t, repo, "git", "add", ".")
	runIn(t, repo, "git", "commit", "-q", "-m", "init")
	got, err := FileAt(repo, "HEAD", "a.md")
	if err != nil {
		t.Fatal(err)
	}
	if got != "hello" {
		t.Errorf("got %q", got)
	}
}

func TestChangedMarkdownEmptyBase(t *testing.T) {
	repo := gitInit(t)
	write(t, repo, "a.md", "one")
	write(t, repo, "sub/b.md", "two")
	runIn(t, repo, "git", "add", ".")
	runIn(t, repo, "git", "commit", "-q", "-m", "init")

	files, err := ChangedMarkdown(repo, "", "HEAD")
	if err != nil {
		t.Fatal(err)
	}
	sort.Strings(files)
	if len(files) != 2 {
		t.Errorf("got %v", files)
	}
}
