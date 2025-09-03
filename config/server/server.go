package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Adgytec/adgytec-flow/config/app"
	app_init "github.com/Adgytec/adgytec-flow/config/init"
	"github.com/Adgytec/adgytec-flow/config/router"
)

type IServer interface {
	ListenAndServe() error
	Shutdown() error
}

type httpServer struct {
	server *http.Server
	app    app.IApp
}

func (s *httpServer) ListenAndServe() error {
	log.Println("Server now listening")
	return s.server.ListenAndServe()
}

func (s *httpServer) Shutdown() error {
	log.Println("shutting down server")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	s.app.Shutdown()
	err := s.server.Shutdown(shutdownCtx)

	return err
}

func NewHttpServer(port string) IServer {
	appConfig := app.NewApp()
	app_init.EnsureServicesInitialization(appConfig)
	mux := router.NewApplicationRouter(appConfig)

	var protocols http.Protocols
	protocols.SetUnencryptedHTTP2(true)

	appServer := http.Server{
		Addr:              ":" + port,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		Handler:           mux,
		Protocols:         &protocols,
	}

	return &httpServer{
		server: &appServer,
		app:    appConfig,
	}
}
