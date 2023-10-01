package releasenote

import (
	"testing"

	"releaseros/internal/config"

	"github.com/stretchr/testify/assert"
)

func TestItGeneratesAFooter(t *testing.T) {
	footer := Footer{
		LatestTag:   "v1.0.0",
		PreviousTag: "v0.0.1",
	}

	actual, err := footer.Generate(config.Config{
		Footer: "**Full Changelog**: https://gitweb.repo/compare/{{ .PreviousTag }}...{{ .LatestTag }}\n",
	})
	assert.NoError(t, err)
	assert.Exactly(t, "**Full Changelog**: https://gitweb.repo/compare/v0.0.1...v1.0.0\n", actual)
}
