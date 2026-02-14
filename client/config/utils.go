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
		return nil, errors.Join(errCannotUnmarshalConfigFile, err)
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

func saveConf(conf *ConfigOptions) {
	f, err := os.OpenFile(configFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	yaml.NewEncoder(f).Encode(conf)
}
