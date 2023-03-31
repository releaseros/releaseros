package config

type Config struct {
	InitialReleaseMessage string     `yaml:"initial_release_message,omitempty" json:"initial_release_message,omitempty"`
	Sort                  string     `yaml:"sort,omitempty" json:"sort,omitempty" jsonschema:"enum=asc,enum=desc,enum=,default="`
	Filters               Filters    `yaml:"filters,omitempty" json:"filters,omitempty"`
	Categories            []Category `yaml:"categories,omitempty" json:"categories,omitempty"`
	Footer                string     `yaml:"footer,omitempty" json:"footer,omitempty"`
}

type Filters struct {
	Exclude []string `yaml:"exclude,omitempty" json:"exclude,omitempty"`
}

type Category struct {
	Title  string `yaml:"title,omitempty" json:"title,omitempty"`
	Regexp string `yaml:"regexp,omitempty" json:"regexp,omitempty"`
	Weight int    `yaml:"weight,omitempty" json:"weight,omitempty"`
}
