package testing

import (
	"os"
	stdtesting "testing"

	"github.com/stretchr/testify/assert"
)

func Mkcd(tb stdtesting.TB) string {
	tb.Helper()

	folder := tb.TempDir()

	workdir, err := os.Getwd()
	assert.NoError(tb, err)

	tb.Cleanup(func() {
		assert.NoError(tb, os.Chdir(workdir))
	})
	assert.NoError(tb, os.Chdir(folder))

	return folder
}
