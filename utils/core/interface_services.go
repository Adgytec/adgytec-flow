package core

import (
	"github.com/go-chi/chi/v5"
)

// ServiceInit is used by app before creating IServiceMux or IServicePC to ensure related data initialization
type ServiceInit interface {
	InitService() error
}

// ServiceMux is used by router mux to init service http rest endpoints
type ServiceMux interface {
	BasePath() string
	Router() *chi.Mux
}

// ServicePC is used for inter-service communication
// each service will have their own PC(procedural call) interface
// this interface is just to tell there are 3 things each service must provide
type ServicePC any
