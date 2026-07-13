package gitdelta

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// ChangedMarkdown returns list of *.md files changed between two refs.
// If baseRef is empty, falls back to the working tree state via git ls-tree.
func ChangedMarkdown(repoRoot, baseRef, headRef string) ([]string, error) {
	if baseRef == "" {
		return listAll(repoRoot, headRef)
	}
	out, err := run(repoRoot, "git", "diff", "--name-only", "-z",
		"--diff-filter=ACM", baseRef, headRef, "--", "*.md")
	if err != nil {
		return nil, err
	}
	return splitZ(out), nil
}

func listAll(repoRoot, headRef string) ([]string, error) {
	if headRef == "" {
		headRef = "HEAD"
	}
	out, err := run(repoRoot, "git", "ls-tree", "--name-only", "-r", "-z", headRef)
	if err != nil {
		return nil, err
	}
	all := splitZ(out)
	filtered := make([]string, 0, len(all))
	for _, p := range all {
		if strings.HasSuffix(p, ".md") {
			filtered = append(filtered, p)
		}
	}
	return filtered, nil
}

// FileAt returns file content at the given git ref.
func FileAt(repoRoot, ref, path string) (string, error) {
	out, err := run(repoRoot, "git", "show", ref+":"+path)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func run(dir, name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("%s %s: %v: %s", name, strings.Join(args, " "), err, stderr.String())
	}
	return stdout.Bytes(), nil
}

func splitZ(b []byte) []string {
	parts := bytes.Split(b, []byte{0})
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		s := strings.TrimSpace(string(p))
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
