package testing

import (
	"context"
	"testing"

	"releaseros/internal/git"

	"github.com/stretchr/testify/require"
)

func GitInit(tb testing.TB) {
	tb.Helper()
	out, err := gitExec("init", "-b", "main")
	require.NoError(tb, err)
	require.Contains(tb, out, "Initialized empty Git repository")
}

func GitCommit(tb testing.TB, msg string) {
	tb.Helper()
	out, err := gitExec("commit", "--allow-empty", "-m", msg)
	require.NoError(tb, err)
	require.Contains(tb, out, msg)
}

func GitAnnotatedTag(tb testing.TB, tag, message string) {
	tb.Helper()
	out, err := gitExec("tag", "-a", tag, "-m", message)
	require.NoError(tb, err)
	require.Empty(tb, out)
}

func GitMerge(tb testing.TB, name string) {
	tb.Helper()
	_, err := gitExec("merge", name)
	require.NoError(tb, err)
}

func GitSwitch(tb testing.TB, name string) {
	tb.Helper()
	out, err := gitExec("switch", name)
	require.NoError(tb, err)
	require.Empty(tb, out)
}

func GitSwitchDetach(tb testing.TB, name string) {
	tb.Helper()
	out, err := gitExec("switch", "--detach", name)
	require.NoError(tb, err)
	require.Empty(tb, out)
}

func GitSwitchAndCreateBranch(tb testing.TB, name string) {
	tb.Helper()
	out, err := gitExec("switch", "-c", name)
	require.NoError(tb, err)
	require.Empty(tb, out)
}

func gitExec(args ...string) (string, error) {
	allArgs := []string{
		"-c", "user.name='releaseros'",
		"-c", "user.email='test@releaseros'",
		"-c", "commit.gpgSign=false",
		"-c", "tag.gpgSign=false",
		"-c", "log.showSignature=false",
	}
	allArgs = append(allArgs, args...)
	return git.Exec(context.Background(), allArgs...)
}
