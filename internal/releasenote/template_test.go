package releasenote

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvOrFallbackWhenEnvVarNotFound(t *testing.T) {
	assert.Equal(t, "foo", EnvOr("DOES_NOT_EXIST", "foo"))
}

func TestEnvOr(t *testing.T) {
	_ = os.Setenv("TEST_RELEASEROS_BAR", "bar")
	assert.Equal(t, "bar", EnvOr("TEST_RELEASEROS_BAR", "foo"))
	_ = os.Unsetenv("TEST_RELEASEROS_FOO")
}

func TestEnvOrError(t *testing.T) {
	_ = os.Setenv("TEST_RELEASEROS_BAZ", "baz")
	value, err := EnvOrError("TEST_RELEASEROS_BAZ")
	assert.NoError(t, err)
	assert.Equal(t, "baz", value)
	_ = os.Unsetenv("TEST_RELEASEROS_BAZ")

	_, err = EnvOrError("TEST_RELEASEROS_FOO")
	assert.Error(t, err)
}
