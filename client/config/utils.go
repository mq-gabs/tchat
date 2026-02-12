package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

func loadConfig() (*ConfigOptions, error) {
	data, err := os.ReadFile(configFileName)
	if err != nil {
		return nil, errors.Join(errCannotReadConfigFile, err)
	}

	var options ConfigOptions
	err = yaml.Unmarshal(data, &options)
	if err != nil {
		return nil, errors.Join(errCannotUnmarshalConfigFile)
	}

	return &options, nil
}

func createFile() error {
	_, err := os.Create(configFileName)
	if err != nil {
		return errors.Join(errCannotCreateConfigFile, err)
	}

	return nil
}
