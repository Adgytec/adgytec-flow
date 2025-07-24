package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/Adgytec/adgytec-flow/config/server"
)

func main() {
	rootCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	httpServer := server.CreateHttpServer("8080")
	go func() {
		log.Println("Server starting on :8080.")
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
