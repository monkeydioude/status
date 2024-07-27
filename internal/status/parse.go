package status

import (
	"encoding/json"
	"os"
)

func Parse(path string) ([]Config, error) {
	config := make([]Config, 0)
	data, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}
	if err := json.Unmarshal(data, &config); err != nil {
		return config, err
	}
	return config, nil
}
