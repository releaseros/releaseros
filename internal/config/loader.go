package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type DefaultConfigFileNotFoundError struct{}

func (e *DefaultConfigFileNotFoundError) Error() string {
	return "Default configuration files not found."
}

func LoadDefaultConfig() (Config, error) {
	workingDirectory, err := os.Getwd()
	if err != nil {
		return Config{}, err
	}

	defaultFilenames := []string{
		DefaultFilename,
		".releaseros.yml",
		"releaseros.yaml",
		"releaseros.yml",
	}
	for _, filename := range defaultFilenames {
		filepath := workingDirectory + "/" + filename
		if _, err := os.Stat(filepath); err != nil {
			continue
		}
		return LoadFromFilePath(filepath)
	}

	return Config{}, &DefaultConfigFileNotFoundError{}
}

func LoadFromFilePath(filepath string) (Config, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	return LoadFromReader(file)
}

func LoadFromReader(fd io.Reader) (Config, error) {
	var config Config
	content, err := io.ReadAll(fd)
	if err != nil {
		return config, err
	}
	if err := yaml.UnmarshalStrict(content, &config); err != nil {
		return config, err
	}

	return config, nil
}
