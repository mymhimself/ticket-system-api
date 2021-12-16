package config

import (
	"encoding/json"
	"os"
)

func ReadJSON(path string) (*Config, error) {
	var config ConfigurationType
	configFile, err := os.Open(path)
	defer configFile.Close()
	if err != nil {
		return nil, err
	}

	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&config); err != nil {
		return nil, err
	}

	if config.IsProd {
		return &config.Prod, nil
	} else {
		return &config.Dev, nil
	}
}
