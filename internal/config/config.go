package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	RateLimit *RateLimit // RateLimit - number of requests allowed for one user per time frame
	TimeFrame *TimeFrame // TimeFrame - duration after which the rate limit resets
	Port      *Port      // Port      - http server port
}

type RateLimit struct {
	Limit int64
}

type TimeFrame struct {
	Frame time.Duration
}

type Port struct {
	Address string
}

func MustNew() *Config {
	return &Config{
		RateLimit: &RateLimit{
			Limit: mustGetEnvInt("LIMIT_PER_USER"),
		},
		TimeFrame: &TimeFrame{
			Frame: mustGetEnvDuration("RESET_TIME_FRAME"),
		},
		Port: &Port{
			Address: fmt.Sprintf(":%d", mustGetEnvInt("PORT")),
		},
	}
}

func mustGetEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Sprintf("env variable %s must be set", key))
	}
	return v
}

func mustGetEnvInt(key string) int64 {
	s := mustGetEnv(key)

	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("'%v' value is not a number", key))
	}

	return v
}

func mustGetEnvDuration(key string) time.Duration {
	s := mustGetEnv(key)

	v, err := time.ParseDuration(s)
	if err != nil {
		panic(fmt.Sprintf("'%v' value is not a duration", key))
	}

	return v
}
