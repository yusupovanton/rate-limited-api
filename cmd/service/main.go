package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"

	"github.com/yusupovanton/rate-limited-api/internal/config"
	"github.com/yusupovanton/rate-limited-api/internal/di"
)

const (
	successExitCode = 0
	failExitCode    = 1
)

func main() {
	os.Exit(run())
}

func run() (exitCode int) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env.example file", "err", err)
		return failExitCode
	}

	cfg := config.MustNew()

	container := di.NewContainer(ctx, cfg, logger)

	mux := http.NewServeMux()
	mux.HandleFunc("/example", container.GetExampleHandler().Handle())

	server := &http.Server{
		Addr:    cfg.Port.Address,
		Handler: mux,
	}

	go func() {
		logger.Info("Server starting...", "port", server.Addr)
		if err = server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.With("err", err).Error("Failed to start server")
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = server.Shutdown(shutdownCtx); err != nil {
		logger.With("err", err).Error("Server forced to shutdown:")
		return failExitCode
	}

	logger.Info("Server stopped gracefully")
	return successExitCode
}
