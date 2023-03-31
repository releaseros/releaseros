package config

import _ "embed"

var DefaultFilename = ".releaseros.yaml"

//go:embed config.yaml
var DefaultConfig []byte
