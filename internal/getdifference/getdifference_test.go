package getdifference

import "testing"

func TestGetDifference(t *testing.T) {
	file1 := map[string]any{
		"host":    "hexlet.io",
		"timeout": float64(50),
		"proxy":   "123.234.53.22",
		"follow":  false,
	}
	file2 := map[string]any{
		"timeout": float64(20),
		"verbose": true,
		"host":    "hexlet.io",
	}

	expected := "{\n  - follow: false\n    host: hexlet.io\n  - proxy: 123.234.53.22\n  - timeout: 50\n  + timeout: 20\n  + verbose: true\n}"

	actual := GetDifference(file1, file2)
	if actual != expected {
		t.Fatalf("unexpected diff\nexpected:\n%s\nactual:\n%s", expected, actual)
	}
}
