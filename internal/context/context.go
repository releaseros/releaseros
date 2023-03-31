package context

import (
	stdctx "context"

	"releaseros/internal/config"
)

type Context struct {
	stdctx.Context
	Config config.Config
}

func New(config config.Config) *Context {
	return &Context{
		Context: stdctx.Background(),
		Config:  config,
	}
}
