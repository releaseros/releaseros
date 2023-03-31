package git

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func Exec(ctx context.Context, args ...string) (string, error) {
	extraArgs := []string{
		"-c", "log.showSignature=false",
	}
	args = append(extraArgs, args...)
	cmd := exec.CommandContext(ctx, "git", args...)

	stdout := bytes.Buffer{}
	stderr := bytes.Buffer{}

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", errors.New(stderr.String())
	}

	return stdout.String(), nil
}

func Clean(output string, err error) (string, error) {
	output = strings.ReplaceAll(strings.Split(output, "\n")[0], "'", "")
	if err != nil {
		err = errors.New(strings.TrimSuffix(err.Error(), "\n"))
	}
	return output, err
}

func DescribeTag(ctx context.Context, tag string) (string, error) {
	return Clean(Exec(ctx, "describe", "--tags", "--abbrev=0", tag))
}

func LatestTag(ctx context.Context) (string, error) {
	return Clean(DescribeTag(ctx, "HEAD"))
}

func PreviousTag(ctx context.Context, tag string) (string, error) {
	return Clean(DescribeTag(ctx, fmt.Sprintf("tags/%s^", tag)))
}

func Log(ctx context.Context, from, to string) (string, error) {
	return Exec(ctx, "log", "--pretty=oneline", "--abbrev-commit", "--no-decorate", "--no-color", fmt.Sprintf("tags/%s..tags/%s", from, to))
}

func LogTo(ctx context.Context, to string) (string, error) {
	return Exec(ctx, "log", "--pretty=oneline", "--abbrev-commit", "--no-decorate", "--no-color", fmt.Sprintf("tags/%s", to))
}
