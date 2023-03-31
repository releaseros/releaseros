package command

import (
	"bytes"
	stdtesting "testing"

	"releaseros/internal/testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate_ItErrorsWhithoutConfig(t *stdtesting.T) {
	testing.Mkcd(t)

	err := testingApp.Run([]string{"releaseros", "generate"})
	assert.Error(t, err)
	assert.ErrorContains(t, err, "Default configuration files not found.")
}

func TestGenerate_ItErrorsWhithoutAGitRepository(t *stdtesting.T) {
	testing.Mkcd(t)

	assert.NoError(t, testingApp.Run([]string{"releaseros", "init"}))

	err := testingApp.Run([]string{"releaseros", "generate"})
	assert.Error(t, err)
	assert.ErrorContains(t, err, "fatal: not a git repository")
}

func TestGenerate_ItErrorsWhithoutAtLeastOneTag(t *stdtesting.T) {
	testing.Mkcd(t)

	assert.NoError(t, testingApp.Run([]string{"releaseros", "init"}))

	testing.GitInit(t)
	testing.GitCommit(t, "initial commit")

	err := testingApp.Run([]string{"releaseros", "generate"})
	assert.Error(t, err)
	assert.ErrorContains(t, err, "fatal: No names found, cannot describe anything.")
}

func TestGenerate(t *stdtesting.T) {
	testing.Mkcd(t)

	testing.GitInit(t)
	testing.GitCommit(t, "initial commit")
	testing.GitCommit(t, "feat: whatthecommit")
	testing.GitAnnotatedTag(t, "v1.0.0", "initial release")

	var outputBuffer bytes.Buffer
	testingApp.Writer = &outputBuffer

	assert.NoError(t, testingApp.Run([]string{"releaseros", "init"}))
	assert.NoError(t, testingApp.Run([]string{"releaseros", "generate"}))
	assert.Exactly(t, "Initial Release\n", outputBuffer.String())
	outputBuffer.Reset()

	testing.GitCommit(t, "feat: what happens in vegas stays in vegas")
	testing.GitCommit(t, "fix: it's secret!")
	testing.GitCommit(t, "fix: never before had a small typo like this one caused so much damage")
	testing.GitCommit(t, "feat: commit 2")
	testing.GitCommit(t, "Merge branch main into develop")
	testing.GitCommit(t, "feat: does it work? maybe. will I check? no")
	testing.GitCommit(t, "chore: don't ask me")
	testing.GitCommit(t, "test: I cannot believe that it took this long to write a test for this")

	testing.GitAnnotatedTag(t, "v1.1.0", "second release")

	assert.NoError(t, testingApp.Run([]string{"releaseros", "generate"}))
	assert.Regexp(t, `^## Release Note

### Features

\* ([a-z0-9]){7} feat: commit 2
\* ([a-z0-9]){7} feat: does it work\? maybe\. will I check\? no
\* ([a-z0-9]){7} feat: what happens in vegas stays in vegas

### Fixes

\* ([a-z0-9]){7} fix: it's secret!
\* ([a-z0-9]){7} fix: never before had a small typo like this one caused so much damage

### Others

\* ([a-z0-9]){7} chore: don't ask me

\*\*Full Changelog\*\*: https://CHANGEME/-/compare/v1\.0\.0\.\.\.v1\.1\.0

$`, outputBuffer.String())
	outputBuffer.Reset()

	testing.GitSwitchDetach(t, "v1.0.0")
	testing.GitSwitchAndCreateBranch(t, "release/v1.0.1")
	testing.GitCommit(t, "fix: this is a hotfix")
	testing.GitAnnotatedTag(t, "v1.0.1", "hotfix")
	testing.GitSwitch(t, "main")
	testing.GitMerge(t, "release/v1.0.1")

	testing.GitSwitchDetach(t, "v1.0.1")
	assert.NoError(t, testingApp.Run([]string{"releaseros", "generate"}))
	assert.Regexp(t, `^## Release Note

### Fixes

\* ([a-z0-9]){7} fix: this is a hotfix

\*\*Full Changelog\*\*: https://CHANGEME/-/compare/v1\.0\.0\.\.\.v1\.0\.1

$`, outputBuffer.String())
	outputBuffer.Reset()

	testing.GitSwitch(t, "main")
	testing.GitCommit(t, "feat: this is a feature")
	testing.GitCommit(t, "fix: this is a fix")
	testing.GitCommit(t, "chore: this is a chore")
	testing.GitCommit(t, "test: this is a test")
	testing.GitCommit(t, "docs: this is a docs")
	testing.GitCommit(t, "refactor: this is a refactor")
	testing.GitAnnotatedTag(t, "v1.2.0", "third release")

	assert.NoError(t, testingApp.Run([]string{"releaseros", "generate"}))
	assert.Regexp(t, `^## Release Note

### Features

\* ([a-z0-9]){7} feat: this is a feature

### Fixes

\* ([a-z0-9]){7} fix: this is a fix
\* ([a-z0-9]){7} fix: this is a hotfix

### Documentation

\* ([a-z0-9]){7} docs: this is a docs

### Others

\* ([a-z0-9]){7} chore: this is a chore
\* ([a-z0-9]){7} refactor: this is a refactor

\*\*Full Changelog\*\*: https://CHANGEME/-/compare/v1\.1\.0\.\.\.v1\.2\.0

$`, outputBuffer.String())
	outputBuffer.Reset()
}
