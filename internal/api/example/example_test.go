package example_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/yusupovanton/rate-limited-api/internal/api/example"
	"github.com/yusupovanton/rate-limited-api/internal/api/example/mocks"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const dummyIP = "127.0.0.1"

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}

type HandlerSuite struct {
	suite.Suite

	logger  *slog.Logger
	rlMock  *mocks.RateLimiter
	handler *example.Handler
}

func (s *HandlerSuite) SetupTest() {
	s.rlMock = mocks.NewRateLimiter(s.T())
	s.logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	s.handler = example.NewHandler(s.rlMock, s.logger)
}

func (s *HandlerSuite) TestHandleRequestAllowed() {
	req, _ := http.NewRequest(http.MethodGet, "/example", nil)
	req.RemoteAddr = dummyIP

	s.rlMock.EXPECT().IsAllowed(dummyIP).Return(true).Once()

	recorder := httptest.NewRecorder()

	handlerFunc := s.handler.Handle()
	handlerFunc.ServeHTTP(recorder, req)

	s.Require().Equal(http.StatusOK, recorder.Code)
}

func (s *HandlerSuite) TestHandleRequestDenied() {
	req, _ := http.NewRequest(http.MethodGet, "/example", nil)
	req.RemoteAddr = dummyIP

	s.rlMock.EXPECT().IsAllowed(dummyIP).Return(true).Times(2)
	s.rlMock.EXPECT().IsAllowed(dummyIP).Return(false).Times(1)

	recorder := httptest.NewRecorder()

	handlerFunc := s.handler.Handle()

	for i := 0; i < 3; i++ {
		handlerFunc.ServeHTTP(recorder, req)
	}

	s.Require().Equal(http.StatusTooManyRequests, recorder.Code)
}

func (s *HandlerSuite) TestHandleMethodIsNotAllo() {
	req, _ := http.NewRequest(http.MethodPatch, "/example", nil)
	req.RemoteAddr = dummyIP

	recorder := httptest.NewRecorder()

	handlerFunc := s.handler.Handle()

	handlerFunc.ServeHTTP(recorder, req)

	s.Require().Equal(http.StatusMethodNotAllowed, recorder.Code)
}
