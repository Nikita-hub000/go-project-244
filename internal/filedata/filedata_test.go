package filedata

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFileDataSuccessWithRelativePaths(t *testing.T) {
	path1 := filepath.Join("..", "..", "testdata", "fixture", "file1.json")
	path2 := filepath.Join("..", "..", "testdata", "fixture", "file2.json")

	got, err := FileData(path1, path2)
	if err != nil {
		t.Fatalf("FileData returned error: %v", err)
	}

	expected := "{\n  - follow: false\n    host: hexlet.io\n  - proxy: 123.234.53.22\n  - timeout: 50\n  + timeout: 20\n  + verbose: true\n}"
	if got != expected {
		t.Fatalf("unexpected diff\nexpected:\n%s\nactual:\n%s", expected, got)
	}
}

func TestFileDataSuccessWithAbsolutePaths(t *testing.T) {
	relPath1 := filepath.Join("..", "..", "testdata", "fixture", "file1.json")
	relPath2 := filepath.Join("..", "..", "testdata", "fixture", "file2.json")
	path1, err := filepath.Abs(relPath1)
	if err != nil {
		t.Fatalf("failed to build abs path for file1: %v", err)
	}
	path2, err := filepath.Abs(relPath2)
	if err != nil {
		t.Fatalf("failed to build abs path for file2: %v", err)
	}

	got, err := FileData(path1, path2)
	if err != nil {
		t.Fatalf("FileData returned error: %v", err)
	}

	expected := "{\n  - follow: false\n    host: hexlet.io\n  - proxy: 123.234.53.22\n  - timeout: 50\n  + timeout: 20\n  + verbose: true\n}"
	if got != expected {
		t.Fatalf("unexpected diff\nexpected:\n%s\nactual:\n%s", expected, got)
	}
}

func TestFileDataReturnsErrorForMissingFile(t *testing.T) {
	path1 := filepath.Join("..", "..", "testdata", "fixture", "missing.json")
	path2 := filepath.Join("..", "..", "testdata", "fixture", "file2.json")

	_, err := FileData(path1, path2)
	if err == nil {
		t.Fatal("expected error for missing file")
	}
	if !strings.Contains(err.Error(), "failed to get file") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestFileDataReturnsErrorForInvalidJSON(t *testing.T) {
	dir := t.TempDir()
	badJSONPath := filepath.Join(dir, "bad.json")
	if err := os.WriteFile(badJSONPath, []byte("{invalid json"), 0o600); err != nil {
		t.Fatalf("failed to create invalid json file: %v", err)
	}
	goodJSONPath := filepath.Join("..", "..", "testdata", "fixture", "file2.json")

	_, err := FileData(badJSONPath, goodJSONPath)
	if err == nil {
		t.Fatal("expected error for invalid json")
	}
	if !strings.Contains(err.Error(), "failed to unmarshal file") {
		t.Fatalf("unexpected error: %v", err)
	}
}
