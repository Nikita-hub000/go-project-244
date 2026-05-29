package code

import (
	"path/filepath"
	"testing"
)

func TestGenDiffJSON(t *testing.T) {
	path1 := filepath.Join("testdata", "fixture", "file1.json")
	path2 := filepath.Join("testdata", "fixture", "file2.json")

	expected := "{\n  - follow: false\n    host: hexlet.io\n  - proxy: 123.234.53.22\n  - timeout: 50\n  + timeout: 20\n  + verbose: true\n}"

	actual, err := GenDiff(path1, path2, "stylish")
	if err != nil {
		t.Fatalf("GenDiff returned error: %v", err)
	}

	if actual != expected {
		t.Fatalf("unexpected diff\nexpected:\n%s\nactual:\n%s", expected, actual)
	}
}
