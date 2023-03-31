package releasenote

import (
	"errors"
	"os"
	"testing"

	"releaseros/internal/config"
	"releaseros/internal/context"

	"github.com/stretchr/testify/assert"
)

type gitTagFinderMock struct {
	latestTag        string
	latestTagError   error
	previousTag      string
	previousTagError error
}

func (mock gitTagFinderMock) LatestTag(ctx *context.Context) (string, error) {
	return mock.latestTag, mock.latestTagError
}

func (mock gitTagFinderMock) PreviousTag(ctx *context.Context, latestTag string) (string, error) {
	return mock.previousTag, mock.previousTagError
}

type gitLogFinderMock struct {
	logs string
	err  error
}

func (mock gitLogFinderMock) LogTo(ctx *context.Context, latestTag string) (string, error) {
	return mock.logs, mock.err
}

func (mock gitLogFinderMock) Log(ctx *context.Context, previousTag, latestTag string) (string, error) {
	return mock.logs, mock.err
}

func DefaultConfig(t *testing.T) config.Config {
	t.Helper()

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	file := dir + "/../config/config.yaml"
	if assert.FileExists(t, file) == false {
		t.Fatal(errors.New("config file not found"))
	}

	config, err := config.LoadFromFilePath(file)
	if err != nil {
		t.Fatal(err)
	}

	return config
}

func TestGenerate(t *testing.T) {
	data := []struct {
		name             string
		expected         string
		latestTag        string
		previousTag      string
		previousTagError error
		logs             string
		config           config.Config
	}{
		{
			name: "default",
			expected: `## Release Note

### Features

* 12345679 feat: commit 1
* 12345676 feat: commit 4
* 12345675 feat: commit 5

### Fixes

* 12345677 fix: a commit 3
* 12345678 fix: commit 2

### Others

* 12345674 chore: commit 6

**Full Changelog**: https://CHANGEME/-/compare/v0.0.1...v1.0.0

`,

			latestTag:   "v1.0.0",
			previousTag: "v0.0.1",
			logs: `12345679 feat: commit 1
12345678 fix: commit 2
12345677 fix: a commit 3
12345676 feat: commit 4
12345672 Merge branch main into develop
12345675 feat: commit 5
12345674 chore: commit 6
12345673 test: commit 7
`,
			config: DefaultConfig(t),
		},
		{
			name: "default without footer",
			expected: `## Release Note

### Features

* 12345679 feat: commit 1
* 12345676 feat: commit 4
* 12345675 feat: commit 5

### Fixes

* 12345677 fix: a commit 3
* 12345678 fix: commit 2

### Others

* 12345674 chore: commit 6

`,

			latestTag:   "v1.0.0",
			previousTag: "v0.0.1",
			logs: `12345679 feat: commit 1
12345678 fix: commit 2
12345677 fix: a commit 3
12345676 feat: commit 4
12345672 Merge branch main into develop
12345675 feat: commit 5
12345674 chore: commit 6
12345673 test: commit 7
`,
			config: config.Config{
				Sort: "asc",
				Filters: config.Filters{
					Exclude: []string{
						"^test:",
						"^ci:",
						"merge conflict",
						"Merge pull request",
						"Merge remote-tracking branch",
						"Merge branch",
					},
				},
				Categories: []config.Category{
					{
						Title:  "Features",
						Regexp: `^.*?feat(\([[:word:]]+\))??!?:.+$`,
						Weight: 1,
					},
					{
						Title:  "Fixes",
						Regexp: `^.*?fix(\([[:word:]]+\))??!?:.+$`,
						Weight: 2,
					},
					{
						Title:  "Documentation",
						Regexp: `^.*?docs(\([[:word:]]+\))??!?:.+$`,
						Weight: 3,
					},
					{
						Title:  "Others",
						Weight: 9999,
					},
				},
			},
		},
		{
			name: "no categories",
			expected: `## Release Note

* 12345674 chore: commit 6
* 12345679 feat: commit 1
* 12345676 feat: commit 4
* 12345675 feat: commit 5
* 12345677 fix: a commit 3
* 12345678 fix: commit 2

**Full Changelog**: https://gitweb.repo/compare/v0.0.1...v1.0.0

`,

			latestTag:   "v1.0.0",
			previousTag: "v0.0.1",
			logs: `12345679 feat: commit 1
12345676 feat: commit 4
12345675 feat: commit 5
12345678 fix: commit 2
12345677 fix: a commit 3
12345672 Merge branch main into develop
12345674 chore: commit 6
12345673 test: commit 7
`,
			config: config.Config{
				Sort: "asc",
				Filters: config.Filters{
					Exclude: []string{
						"^test:",
						"^ci:",
						"merge conflict",
						"Merge pull request",
						"Merge remote-tracking branch",
						"Merge branch",
					},
				},
				Categories: []config.Category{},
				Footer:     "**Full Changelog**: https://gitweb.repo/compare/{{ .PreviousTag }}...{{ .LatestTag }}\n",
			},
		},
		{
			name: "no categories without footer",
			expected: `## Release Note

* 12345674 chore: commit 6
* 12345679 feat: commit 1
* 12345676 feat: commit 4
* 12345675 feat: commit 5
* 12345677 fix: a commit 3
* 12345678 fix: commit 2

`,

			latestTag:   "v1.0.0",
			previousTag: "v0.0.1",
			logs: `12345679 feat: commit 1
12345676 feat: commit 4
12345675 feat: commit 5
12345678 fix: commit 2
12345677 fix: a commit 3
12345672 Merge branch main into develop
12345674 chore: commit 6
12345673 test: commit 7
`,
			config: config.Config{
				Sort: "asc",
				Filters: config.Filters{
					Exclude: []string{
						"^test:",
						"^ci:",
						"merge conflict",
						"Merge pull request",
						"Merge remote-tracking branch",
						"Merge branch",
					},
				},
				Categories: []config.Category{},
			},
		},
		{
			name: "sort asc",
			expected: `## Release Note

* 12345679 a: commit
* 12345676 b: commit
* 12345675 c: commit

`,

			latestTag:   "v1.0.0",
			previousTag: "v0.0.1",
			logs: `12345675 c: commit
12345679 a: commit
12345676 b: commit
`,
			config: config.Config{
				Sort:       "asc",
				Filters:    config.Filters{},
				Categories: []config.Category{},
			},
		},
		{
			name: "sort desc",
			expected: `## Release Note

* 12345675 c: commit
* 12345676 b: commit
* 12345679 a: commit

`,

			latestTag:   "v1.0.0",
			previousTag: "v0.0.1",
			logs: `12345675 c: commit
12345679 a: commit
12345676 b: commit
`,
			config: config.Config{
				Sort:       "desc",
				Filters:    config.Filters{},
				Categories: []config.Category{},
			},
		},
		{
			name: "sort default",
			expected: `## Release Note

* 12345675 c: commit
* 12345679 a: commit
* 12345676 b: commit

`,

			latestTag:   "v1.0.0",
			previousTag: "v0.0.1",
			logs: `12345675 c: commit
12345679 a: commit
12345676 b: commit
`,
			config: config.Config{
				Sort:       "",
				Filters:    config.Filters{},
				Categories: []config.Category{},
			},
		},
		{
			name: "categories breaking change",
			expected: `## Release Note

### BREAKING CHANGES / Features

* 98765431 feat!: is a breaking change
* 98765432 feat(foo)!: is a breaking change

### Features

* 12345679 feat: commit 1
* 12345676 feat: commit 4
* 12345675 feat: commit 5

### Fixes

* 12345677 fix: a commit 3
* 12345678 fix: commit 2

### Others

* 12345674 chore: commit 6

`,

			latestTag:   "v1.0.0",
			previousTag: "v0.0.1",
			logs: `12345679 feat: commit 1
12345678 fix: commit 2
12345677 fix: a commit 3
98765431 feat!: is a breaking change
12345676 feat: commit 4
12345672 Merge branch main into develop
12345675 feat: commit 5
98765432 feat(foo)!: is a breaking change
12345674 chore: commit 6
12345673 test: commit 7
`,
			config: config.Config{
				Sort: "asc",
				Filters: config.Filters{
					Exclude: []string{
						"^test:",
						"^ci:",
						"merge conflict",
						"Merge pull request",
						"Merge remote-tracking branch",
						"Merge branch",
					},
				},
				Categories: []config.Category{
					{
						Title:  "BREAKING CHANGES / Features",
						Regexp: `^.*?feat(\([[:word:]]+\))??!+:.+$`,
						Weight: 1,
					},
					{
						Title:  "Features",
						Regexp: `^.*?feat(\([[:word:]]+\))??!{0}:.+$`,
						Weight: 2,
					},
					{
						Title:  "Fixes",
						Regexp: `^.*?fix(\([[:word:]]+\))??!?:.+$`,
						Weight: 3,
					},
					{
						Title:  "Documentation",
						Regexp: `^.*?docs(\([[:word:]]+\))??!?:.+$`,
						Weight: 4,
					},
					{
						Title:  "Others",
						Weight: 9999,
					},
				},
			},
		},
		{
			name:     "initial release message",
			expected: "Initial release",

			latestTag:        "v1.0.0",
			previousTag:      "",
			previousTagError: errors.New("previous tag not found"),
			logs: `12345679 feat: commit 1
12345678 fix: commit 2
12345677 fix: a commit 3
98765431 feat!: is a breaking change
12345676 feat: commit 4
12345672 Merge branch main into develop
12345675 feat: commit 5
98765432 feat(foo)!: is a breaking change
12345674 chore: commit 6
12345673 test: commit 7
`,
			config: config.Config{
				InitialReleaseMessage: "Initial release",
			},
		},
	}

	for _, item := range data {
		t.Run(item.name, func(t *testing.T) {
			releaseNoteGenerator := Generator{
				gitTagFinder: gitTagFinderMock{
					latestTag:        item.latestTag,
					previousTag:      item.previousTag,
					previousTagError: item.previousTagError,
				},
				gitLogFinder: gitLogFinderMock{
					logs: item.logs,
				},
			}
			actual, err := releaseNoteGenerator.Generate(context.New(item.config))
			assert.NoError(t, err)
			assert.Equal(t, item.expected, actual)
		})
	}
}

func TestError(t *testing.T) {
	data := []struct {
		name             string
		latestTag        string
		latestTagError   error
		previousTag      string
		previousTagError error
		logs             string
		logsError        error
		config           config.Config
	}{
		{
			name:           "latest tag error",
			latestTag:      "",
			latestTagError: errors.New("latest tag error"),
			config:         config.Config{},
		},
		{
			name:      "log error",
			latestTag: "v1.0.0",
			logsError: errors.New("log error"),
			config:    config.Config{},
		},
		{
			name:      "record filter error",
			latestTag: "v1.0.0",
			logs:      "12345679 feat: commit 1",
			config: config.Config{
				Filters: config.Filters{
					Exclude: []string{
						// Negative lookahead are not supported so it should error.
						"(?!.)",
					},
				},
			},
		},
	}

	for _, item := range data {
		t.Run(item.name, func(t *testing.T) {
			releaseNoteGenerator := Generator{
				gitTagFinder: gitTagFinderMock{
					latestTag:        item.latestTag,
					latestTagError:   item.latestTagError,
					previousTag:      item.previousTag,
					previousTagError: item.previousTagError,
				},
				gitLogFinder: gitLogFinderMock{
					logs: item.logs,
					err:  item.logsError,
				},
			}
			_, err := releaseNoteGenerator.Generate(context.New(item.config))
			assert.Error(t, err)
		})
	}
}
