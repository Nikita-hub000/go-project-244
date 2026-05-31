package parsers

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

func ParseJSON(data []byte) (map[string]any, error) {
	var decoded map[string]any
	if err := json.Unmarshal(data, &decoded); err != nil {
		return nil, fmt.Errorf("failed to parse json: %w", err)
	}

	return decoded, nil
}

func ParseYAML(data []byte) (map[string]any, error) {
	var decoded map[string]any
	if err := yaml.Unmarshal(data, &decoded); err != nil {
		return nil, fmt.Errorf("failed to parse yaml: %w", err)
	}

	return decoded, nil
}
