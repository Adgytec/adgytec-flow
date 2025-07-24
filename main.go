package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Adgytec/adgytec-flow/config/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file")
	}

	port := os.Getenv("port")

	rootCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	httpServer := server.CreateHttpServer(port)
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
