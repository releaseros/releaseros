package command

import (
	stdtesting "testing"

	"releaseros/internal/config"
	"releaseros/internal/testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *stdtesting.T) {
	testing.Mkcd(t)

	cases := []struct {
		expectedFile string
		args         []string
	}{
		{expectedFile: config.DefaultFilename, args: []string{"releaseros", "init"}},
		{expectedFile: "banana.yml", args: []string{"releaseros", "init", "--config", "banana.yml"}},
	}

	for _, c := range cases {
		assert.NoError(t, testingApp.Run(c.args))
		assert.FileExists(t, c.expectedFile)
	}
}
