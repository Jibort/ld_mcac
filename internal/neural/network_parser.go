package neural

import (
	"encoding/json"
	"fmt"
	"os"
)

func ParseConfig(filePath string) (*NetworkConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %v", err)
	}
	defer file.Close()

	var config NetworkConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("error decoding config file: %v", err)
	}

	if config.Layers <= 0 || len(config.NeuronsPerLayer) != config.Layers {
		return nil, fmt.Errorf("invalid configuration: mismatch between layers and neurons per layer")
	}

	return &config, nil
}
