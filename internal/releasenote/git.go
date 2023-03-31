package releasenote

import (
	"releaseros/internal/context"
	"releaseros/internal/git"
)

type GitTagFinder interface {
	LatestTag(ctx *context.Context) (string, error)
	PreviousTag(ctx *context.Context, latestTag string) (string, error)
}

type GitLogFinder interface {
	LogTo(ctx *context.Context, latestTag string) (string, error)
	Log(ctx *context.Context, previousTag, latestTag string) (string, error)
}

type gitTagFinder struct{}

func (gitTagFinder gitTagFinder) LatestTag(ctx *context.Context) (string, error) {
	return git.LatestTag(ctx)
}

func (gitTagFinder gitTagFinder) PreviousTag(ctx *context.Context, latestTag string) (string, error) {
	return git.PreviousTag(ctx, latestTag)
}

type gitLogFinder struct{}

func (gitLogFinder gitLogFinder) LogTo(ctx *context.Context, latestTag string) (string, error) {
	return git.LogTo(ctx, latestTag)
}

func (gitLogFinder gitLogFinder) Log(ctx *context.Context, previousTag, latestTag string) (string, error) {
	return git.Log(ctx, previousTag, latestTag)
}
