package releasenote

import (
	"fmt"
	"os"
	"text/template"
)

var FunctionsForTemplate = template.FuncMap{
	"env":        os.Getenv,
	"envOr":      EnvOr,
	"envOrError": EnvOrError,
}

func EnvOr(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

func EnvOrError(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf("could not find ENV %q", key)
	}
	return value, nil
}
