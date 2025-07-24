package server

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

type IServer interface {
	ListenAndServe() error
	Shutdown() error
}

type httpServer struct {
	server *http.Server
	// stopOngoingGracefully context.CancelFunc
}

func (s *httpServer) ListenAndServe() error {
	log.Println("Server now listening")
	return s.server.ListenAndServe()
}

func (s *httpServer) Shutdown() error {
	log.Println("shutting down server")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := s.server.Shutdown(shutdownCtx)
	// if s.stopOngoingGracefully != nil {
	// 	s.stopOngoingGracefully()
	// }

	return err
}

func CreateHttpServer(port string) IServer {
	// TODO: will get mux from router package later for now simple mux
	handle := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "working")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handle)

	var protocols http.Protocols
	protocols.SetUnencryptedHTTP2(true)
	// ongoingCtx, stopOngoingGracefully := context.WithCancel(context.Background())
	appServer := http.Server{
		Addr:              ":" + port,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		// BaseContext: func(_ net.Listener) context.Context {
		// 	return ongoingCtx
		// },
		Handler:   mux,
		Protocols: &protocols,
	}

	return &httpServer{
		server: &appServer,
		// stopOngoingGracefully: stopOngoingGracefully,
	}
}
