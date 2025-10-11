package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Adgytec/adgytec-flow/config/server"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file")
	}

	// add logger details
	logLevelStr := strings.ToLower(os.Getenv("LOG_LEVEL"))
	logLevel, err := zerolog.ParseLevel(logLevelStr)
	if err != nil {
		logLevel = zerolog.InfoLevel // default
	}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(logLevel)

	port := os.Getenv("PORT")

	rootCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	httpServer, serverErr := server.NewHttpServer(port)
	if serverErr != nil {
		log.Fatalf("Error creating new http server. Cause: %v", serverErr)
	}

	go func() {
		log.Printf("Server starting on port %s.", port)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	<-rootCtx.Done()
	stop()

	if err := httpServer.Shutdown(); err != nil {
		log.Printf("Graceful shutdown failed: %v", err)
	} else {
		log.Println("Server shut down gracefully.")
	}
}
