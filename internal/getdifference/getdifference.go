package getdifference

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

func GetDifference(file1, file2 map[string]any) string {
	var answer strings.Builder
	answer.WriteString("{")

	keys := collectSortedUniqueKeys(file1, file2)

	for _, v := range keys {
		value1, ok1 := file1[v]
		value2, ok2 := file2[v]
		if ok1 && !ok2 {
			fmt.Fprintf(&answer, "\n  - %s: %v", v, value1)
			continue
		}
		if !ok1 && ok2 {
			fmt.Fprintf(&answer, "\n  + %s: %v", v, value2)
			continue
		}
		if value2 != value1 {
			fmt.Fprintf(&answer, "\n  - %s: %v", v, value1)
			fmt.Fprintf(&answer, "\n  + %s: %v", v, value2)
			continue
		}
		fmt.Fprintf(&answer, "\n    %s: %v", v, value1)
	}
	return answer.String() + "\n}"
}

func collectSortedUniqueKeys(file1, file2 map[string]any) []string {
	keys1 := slices.Collect(maps.Keys(file1))
	keys2 := slices.Collect(maps.Keys(file2))
	allKeys := append(keys1, keys2...)
	unique := make(map[string]struct{}, len(allKeys))
	keys := make([]string, 0, len(allKeys))
	for _, k := range allKeys {
		if _, ok := unique[k]; ok {
			continue
		}
		unique[k] = struct{}{}
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}
