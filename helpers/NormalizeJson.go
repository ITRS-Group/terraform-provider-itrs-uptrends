package helpers

import (
	"encoding/json"
	"sort"
)

// Recursively sort map keys for stable JSON output
func sortKeys(v interface{}) interface{} {
	switch vv := v.(type) {
	case map[string]interface{}:
		keys := make([]string, 0, len(vv))
		for k := range vv {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		sorted := make(map[string]interface{}, len(vv))
		for _, k := range keys {
			sorted[k] = sortKeys(vv[k])
		}
		return sorted
	case []interface{}:
		for i, u := range vv {
			vv[i] = sortKeys(u)
		}
	}
	return v
}

func NormalizeJSON(input string) (string, error) {
	var v interface{}
	if err := json.Unmarshal([]byte(input), &v); err != nil {
		return "", err
	}
	// Sort keys for stable output
	v = sortKeys(v)
	out, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(out), nil
}
