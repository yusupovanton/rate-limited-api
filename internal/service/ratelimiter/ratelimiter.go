package ratelimiter

import (
	"fmt"
	"log/slog"
	"sync"
	"time"
)

type RateLimiter struct {
	logger *slog.Logger

	requests map[string]int64
	mu       sync.Mutex
	limit    int64
	window   time.Duration
}

func NewRateLimiter(logger *slog.Logger, limit int64, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string]int64),
		limit:    limit,
		window:   window,
		logger:   logger,
	}
}

func (rl *RateLimiter) IsAllowed(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if _, ok := rl.requests[ip]; !ok {
		go rl.resetAfter(ip, rl.window)
	}
	rl.requests[ip]++

	rl.logger.Info(fmt.Sprintf("this is request no. %d", rl.requests[ip]))

	return rl.requests[ip] <= rl.limit
}

func (rl *RateLimiter) resetAfter(ip string, duration time.Duration) {
	time.Sleep(duration)

	rl.mu.Lock()

	defer rl.mu.Unlock()
	delete(rl.requests, ip)
}
