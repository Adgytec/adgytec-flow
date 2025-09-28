package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Adgytec/adgytec-flow/config/app"
	"github.com/Adgytec/adgytec-flow/config/appcron"
	"github.com/Adgytec/adgytec-flow/config/appinit"
	"github.com/Adgytec/adgytec-flow/config/router"
)

type Server interface {
	ListenAndServe() error
	Shutdown() error
}

type httpServer struct {
	server   *http.Server
	app      app.App
	cronStop context.CancelFunc
}

func (s *httpServer) ListenAndServe() error {
	log.Println("Server now listening")
	return s.server.ListenAndServe()
}

func (s *httpServer) Shutdown() error {
	log.Println("shutting down server")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if s.cronStop != nil {
		s.cronStop()
	}

	s.app.Shutdown(shutdownCtx)
	err := s.server.Shutdown(shutdownCtx)

	return err
}

func NewHttpServer(port string) (Server, error) {
	appConfig, appConfigErr := app.NewApp()
	if appConfigErr != nil {
		return nil, appConfigErr
	}

	appinit.EnsureServicesInitialization(appConfig)
	mux := router.NewApplicationRouter(appConfig)

	cronCtx, cronCancel := context.WithCancel(context.Background())
	go appcron.ServicesCronJobs(cronCtx, appConfig)

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

	// note: currently multiple services are not implemented, it is intentional
	// current development focus is in creating robust application skeleton for core services communication and methods identification
	// than the implementation will be added
	// Logger will also be changed before moving into production
	return &httpServer{
		server:   &appServer,
		app:      appConfig,
		cronStop: cronCancel,
	}, nil
}
