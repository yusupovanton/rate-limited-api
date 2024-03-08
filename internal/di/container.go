package di

import (
	"context"
	"log/slog"

	"github.com/yusupovanton/rate-limited-api/internal/api/example"
	"github.com/yusupovanton/rate-limited-api/internal/config"
	"github.com/yusupovanton/rate-limited-api/internal/service/ratelimiter"
)

type container struct {
	ctx    context.Context
	cfg    *config.Config
	logger *slog.Logger

	exampleHandler *example.Handler

	rateLimiter *ratelimiter.RateLimiter
}

func NewContainer(ctx context.Context, cfg *config.Config, logger *slog.Logger) *container {
	return &container{ctx: ctx, cfg: cfg, logger: logger}
}

func (c *container) GetExampleHandler() *example.Handler {
	return get(&c.exampleHandler, func() *example.Handler {
		return example.NewHandler(
			c.getRateLimiter(),
			c.logger,
		)
	})
}

func (c *container) getRateLimiter() *ratelimiter.RateLimiter {
	return get(&c.rateLimiter, func() *ratelimiter.RateLimiter {
		return ratelimiter.NewRateLimiter(
			c.logger,
			c.cfg.RateLimit.Limit,
			c.cfg.TimeFrame.Frame,
		)
	})
}

func get[T comparable](obj *T, builder func() T) T {
	if *obj != *new(T) {
		return *obj
	}
	*obj = builder()
	return *obj
}
