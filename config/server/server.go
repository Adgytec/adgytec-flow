package server

import (
	"context"
	"net/http"
	"time"

	"github.com/Adgytec/adgytec-flow/config/app"
	"github.com/Adgytec/adgytec-flow/config/appinit"
	"github.com/Adgytec/adgytec-flow/config/router"
	"github.com/rs/zerolog/log"
)

type Server interface {
	ListenAndServe() error
	Shutdown() error
}

type httpServer struct {
	server *http.Server
	app    app.App
}

func (s *httpServer) ListenAndServe() error {
	log.Info().Msg("server started listening")
	return s.server.ListenAndServe()
}

func (s *httpServer) Shutdown() error {
	log.Info().Msg("server shutting down")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	s.app.Shutdown(shutdownCtx)
	err := s.server.Shutdown(shutdownCtx)

	return err
}

func NewHttpServer(port string) (Server, error) {
	appConfig, appConfigErr := app.NewApp()
	if appConfigErr != nil {
		return nil, appConfigErr
	}

	appInitErr := appinit.EnsureServicesInitialization(appConfig)
	if appInitErr != nil {
		return nil, appInitErr
	}

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

	// note: currently multiple services are not implemented, it is intentional
	// current development focus is in creating robust application skeleton for core services communication and methods identification
	// than the implementation will be added
	// Logger will also be changed before moving into production
	return &httpServer{
		server: &appServer,
		app:    appConfig,
	}, nil
}
