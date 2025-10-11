package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/Adgytec/adgytec-flow/config/server"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Warn().
			Err(envErr).
			Msg("failed to load .env")
	}

	// add logger details
	logLevelStr := strings.ToLower(os.Getenv("LOG_LEVEL"))
	logLevel, parseErr := zerolog.ParseLevel(logLevelStr)
	if parseErr != nil {
		logLevel = zerolog.InfoLevel // default
	}
	zerolog.SetGlobalLevel(logLevel)

	if os.Getenv("ENV") == "development" {
		zerolog.TimeFieldFormat = time.RFC3339
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC3339,
		})
	} else {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	}

	port := os.Getenv("PORT")

	rootCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	httpServer, serverErr := server.NewHttpServer(port)
	if serverErr != nil {
		log.Fatal().
			Err(serverErr).
			Msg("error creating new http server")
	}

	go func() {
		log.Info().
			Str("port", port).
			Msg("server started listening")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panic().
				Err(err).
				Send()
		}
	}()
	<-rootCtx.Done()
	stop()

	if err := httpServer.Shutdown(); err != nil {
		log.Error().
			Err(err).
			Msg("graceful shutdown failed")
	} else {
		log.Info().
			Msg("server shutdown gracefully")
	}
}
