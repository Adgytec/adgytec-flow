package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

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
		log.Warn().
			Err(parseErr).
			Str("log_level_provided", logLevelStr).
			Msg("invalid log level provided, defaulting to 'info'")
		logLevel = zerolog.InfoLevel // default
	}
	zerolog.SetGlobalLevel(logLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	var output io.Writer = os.Stderr
	if os.Getenv("ENV") == "development" {
		output = zerolog.ConsoleWriter{
			Out: os.Stderr,
			FieldsExclude: []string{
				zerolog.TimestampFieldName,
				"remote_ip",
				"user_agent",
				"git_revision",
				"go_version",
			},
		}
	}
	log.Logger = log.Output(output)

	port := os.Getenv("PORT")

	rootCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	httpServer, serverErr := server.NewHttpServer(port)
	if serverErr != nil {
		log.Fatal().
			Err(serverErr).
			Str("action", "new http server").
			Send()
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
			Str("action", "http server shutdown").
			Send()
	} else {
		log.Info().
			Msg("server shutdown gracefully")
	}
}
