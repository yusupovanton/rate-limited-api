package ratelimiter_test

import (
	"context"
	cryptoRand "crypto/rand"
	"fmt"
	"github.com/stretchr/testify/suite"
	"log/slog"
	"math/big"
	"os"
	"testing"
	"time"

	rl "github.com/yusupovanton/rate-limited-api/internal/service/ratelimiter"
)

const (
	defaultLimit     = 2
	defaultTimeFrame = time.Second * 1
)

func TestRateLimiterService(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

type ServiceSuite struct {
	suite.Suite

	ctx    context.Context
	logger *slog.Logger
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()
	s.logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func (s *ServiceSuite) TestService() {
	service := rl.NewRateLimiter(s.logger, defaultLimit, defaultTimeFrame)

	testcases := []struct {
		name              string
		ip                string
		callsCount        int
		expectedLastValue bool
	}{
		{
			"single_request",
			generateRandomIP(),
			1,
			true,
		},
		{
			"more_than_limit",
			generateRandomIP(),
			3,
			false,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			var gotLastValue bool

			for i := 0; i <= tc.callsCount; i++ {
				gotLastValue = service.IsAllowed(tc.ip)
			}

			if tc.expectedLastValue {
				s.Require().True(gotLastValue)
			} else {
				s.Require().False(gotLastValue)
			}
		})
	}
}

func generateRandomIP() string {
	return fmt.Sprintf("%d.%d.%d.%d",
		randomInt64(256), // Generate a secure random number between 0 and 255
		randomInt64(256),
		randomInt64(256),
		randomInt64(256))
}

func randomInt64(maxNum int64) int64 {
	nBig, err := cryptoRand.Int(cryptoRand.Reader, big.NewInt(maxNum))
	if err != nil {
		panic(err)
	}

	return nBig.Int64()
}
