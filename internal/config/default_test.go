package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultConfigIsLoadable(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	file := dir + "/config.yaml"
	if assert.FileExists(t, file) == false {
		return
	}

	actual, err := LoadFromFilePath(file)
	if err != nil {
		t.Fatal(err)
	}

	assert.Exactly(
		t,
		Config{
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
					Weight: 10,
				},
				{
					Title:  "Fixes",
					Regexp: `^.*?fix(\([[:word:]-]+\))??!?:.+$`,
					Weight: 20,
				},
				{
					Title:  "Documentation",
					Regexp: `^.*?docs(\([[:word:]-]+\))??!?:.+$`,
					Weight: 30,
				},
				{
					Title:  "Others",
					Weight: 9999,
				},
			},
			Footer: "**Full Changelog**: https://CHANGEME/-/compare/{{ .PreviousTag }}...{{ .LatestTag }}\n",
		},
		actual,
	)
}
