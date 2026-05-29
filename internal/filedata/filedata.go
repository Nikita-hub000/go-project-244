package filedata

import (
	"code/internal/getdifference"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func FileData(path1, path2 string) (string, error) {
	data1, err := getFile(path1)
	if err != nil {
		return "", fmt.Errorf("failed to get file %s: %w", path1, err)
	}
	data2, err := getFile(path2)
	if err != nil {
		return "", fmt.Errorf("failed to get file %s: %w", path2, err)
	}

	var encodedData1 map[string]any

	err = json.Unmarshal(data1, &encodedData1)

	if err != nil {
		return "", fmt.Errorf("failed to unmarshal file %s: %w", path1, err)
	}

	var encodedData2 map[string]any

	err = json.Unmarshal(data2, &encodedData2)

	if err != nil {
		return "", fmt.Errorf("failed to unmarshal file %s: %w", path2, err)
	}

	answer := getdifference.GetDifference(encodedData1, encodedData2)

	return answer, nil
}

func getFile(path string) ([]byte, error) {
	if filepath.IsAbs(path) {
		file, err := os.ReadFile(path)
		if err != nil {
			return []byte{}, fmt.Errorf("failed to read file %s: %w", path, err)
		}
		return file, nil
	}
	filePath, err := filepath.Abs(path)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to get absolute path for %s: %w", path, err)
	}
	file, err := os.ReadFile(filePath)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}
	return file, nil
}
