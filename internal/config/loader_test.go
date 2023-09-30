package config

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestItLoadsTheConfigFromAYamlFilePath(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	file := dir + "/testdata/config.yaml"
	if assert.FileExists(t, file) == false {
		return
	}

	actual, err := LoadFromFilePath(file)
	if err != nil {
		t.Fatal(err)
	}

	assert.Exactly(t, ExpectedConfig(), actual)
}

func TestItErrorsWhenLoadingNonExistentConfigFile(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	file := dir + "/testdata/non-existent-config"
	assert.NoFileExists(t, file)

	config, err := LoadFromFilePath(file)
	assert.Error(t, err)
	assert.EqualError(t, err, fmt.Sprintf("open %s: no such file or directory", file))
	assert.Exactly(t, Config{}, config)
}

func TestItLoadsTheConfigFromReader(t *testing.T) {
	conf := `initial_release_message: "Initial Release\n"
sort: asc
filters:
  exclude:
    - '^test:'
    - '^ci:'
    - 'merge conflict'
    - Merge pull request
    - Merge remote-tracking branch
    - Merge branch
categories:
  - title: 'Features'
    regexp: '^.*?feat(\([[:word:]-]+\))??!?:.+$'
    weight: 1
  - title: 'Fixes'
    regexp: '^.*?fix(\([[:word:]-]+\))??!?:.+$'
    weight: 2
  - title: 'Documentation'
    regexp: ^.*?docs(\([[:word:]-]+\))??!?:.+$
    weight: 3
  - title: Others
    weight: 9999
footer: |
    **Full Changelog**: https://gitweb.repo/compare/{{ .PreviousTag }}...{{ .LatestTag }}
`
	buf := strings.NewReader(conf)
	fmt.Println(buf)
	actual, err := LoadFromReader(buf)

	assert.NoError(t, err)
	assert.Exactly(t, ExpectedConfig(), actual)
}

func TestDefaultSort(t *testing.T) {
	buf := strings.NewReader("")
	fmt.Println(buf)
	actual, err := LoadFromReader(buf)

	assert.NoError(t, err)
	assert.Exactly(t, Config{Sort: ""}, actual)
}

func TestItErrorsWhenLoadingInvalidYaml(t *testing.T) {
	conf := `
filters:
  exclude: '^docs:'
`
	buf := strings.NewReader(conf)
	fmt.Println(buf)
	_, err := LoadFromReader(buf)

	assert.Error(t, err)
	assert.EqualError(t, err, "yaml: unmarshal errors:\n  line 3: cannot unmarshal !!str `^docs:` into []string")
}

type errorReader struct{}

func (errorReader) Read(_ []byte) (n int, err error) {
	return 1, fmt.Errorf("error")
}

func TestItErrorsWhenLoadingErrorReader(t *testing.T) {
	_, err := LoadFromReader(errorReader{})
	require.Error(t, err)
}

func ExpectedConfig() Config {
	return Config{
		InitialReleaseMessage: "Initial Release\n",
		Sort:                  "asc",
		Filters: Filters{
			Exclude: []string{
				"^test:",
				"^ci:",
				"merge conflict",
				"Merge pull request",
				"Merge remote-tracking branch",
				"Merge branch",
			},
		},
		Categories: []Category{
			{
				Title:  "Features",
				Regexp: `^.*?feat(\([[:word:]-]+\))??!?:.+$`,
				Weight: 1,
			},
			{
				Title:  "Fixes",
				Regexp: `^.*?fix(\([[:word:]-]+\))??!?:.+$`,
				Weight: 2,
			},
			{
				Title:  "Documentation",
				Regexp: `^.*?docs(\([[:word:]-]+\))??!?:.+$`,
				Weight: 3,
			},
			{
				Title:  "Others",
				Weight: 9999,
			},
		},
		Footer: "**Full Changelog**: https://gitweb.repo/compare/{{ .PreviousTag }}...{{ .LatestTag }}\n",
	}
}
