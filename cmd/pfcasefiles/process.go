package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func ProcessKit(kitDir string) (Outputs, error) {
	metaPath := filepath.Join(kitDir, "kit.json")
	b, err := os.ReadFile(metaPath)
	if err != nil {
		return Outputs{}, fmt.Errorf("missing kit.json")
	}

	var meta KitMeta
	if err := json.Unmarshal(b, &meta); err != nil {
		return Outputs{}, fmt.Errorf("bad kit.json")
	}
	if strings.TrimSpace(meta.Title) == "" || strings.TrimSpace(meta.Summary) == "" || strings.TrimSpace(meta.Category) == "" {
		return Outputs{}, fmt.Errorf("kit.json missing required fields (title, summary, category)")
	}
	if len(meta.Inputs) == 0 {
		return Outputs{}, fmt.Errorf("kit.json must list inputs")
	}

	// Secret rule (env keys).
	for k := range meta.Env {
		if forbiddenKey(k) {
			return Outputs{}, fmt.Errorf("forbidden env key: %s (no secrets rule)", k)
		}
	}

	// Required input files exist.
	for _, in := range meta.Inputs {
		p := filepath.Join(kitDir, filepath.FromSlash(in.Path))
		if _, err := os.Stat(p); err != nil {
			return Outputs{}, fmt.Errorf("missing required input: %s", in.Path)
		}
	}

	// Walk kit file tree deterministically.
	files, err := listKitFiles(kitDir)
	if err != nil {
		return Outputs{}, err
	}

	entries := make([]FileEntry, 0, len(files))
	var total int64
	hashLines := make([]string, 0, len(files))

	for _, rel := range files {
		abs := filepath.Join(kitDir, filepath.FromSlash(rel))
		data, err := os.ReadFile(abs)
		if err != nil {
			return Outputs{}, err
		}

		// LF-only rule for text-like extensions.
		if isTextLike(rel) && bytesHasCRLF(data) {
			return Outputs{}, fmt.Errorf("CRLF line endings detected: %s", rel)
		}

		sum := sha256.Sum256(data)
		hexsum := hex.EncodeToString(sum[:])
		n := int64(len(data))

		entries = append(entries, FileEntry{Path: rel, Sha256: hexsum, Bytes: n})
		total += n
		hashLines = append(hashLines, fmt.Sprintf("%s  %s", hexsum, rel))
	}

	// entries already in sorted path order (listFiles), but keep explicit.
	sort.Slice(entries, func(i, j int) bool { return entries[i].Path < entries[j].Path })
	sort.Strings(hashLines)

	manifest := strings.Join(hashLines, "\n") + "\n"

	idx := KitIndex{
		Title:      meta.Title,
		Summary:    meta.Summary,
		Category:   meta.Category,
		Files:      entries,
		TotalBytes: total,
	}

	return Outputs{Index: idx, Manifest: manifest}, nil
}

func forbiddenKey(k string) bool {
	u := strings.ToUpper(k)
	for _, sub := range []string{"SECRET", "TOKEN", "PASSWORD", "KEY"} {
		if strings.Contains(u, sub) {
			return true
		}
	}
	return false
}

func isTextLike(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".md", ".json", ".csv", ".txt", ".yaml", ".yml":
		return true
	default:
		return false
	}
}

func bytesHasCRLF(b []byte) bool {
	// simple scan for \r\n
	for i := 0; i+1 < len(b); i++ {
		if b[i] == '\r' && b[i+1] == '\n' {
			return true
		}
	}
	return false
}

func listKitFiles(root string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(root, p)
		if err != nil {
			return err
		}
		files = append(files, filepath.ToSlash(rel))
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Strings(files)
	return files, nil
}
