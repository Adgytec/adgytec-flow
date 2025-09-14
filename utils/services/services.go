package services

import (
	"github.com/go-chi/chi/v5"
)

// Init is used by app before creating IServiceMux or IServicePC to ensure related data initialization
type Init interface {
	InitService() error
}

// Mux is used by router mux to init service http rest endpoints
type Mux interface {
	BasePath() string
	Router() *chi.Mux
}

// PC is used for inter-service communication
// each service will have their own PC(procedural call) interface
// this interface is just to tell there are 3 things each service must provide
type PC any

// Cron is used by services to perform actions on regular intervals
type Cron interface {
	Trigger()
}
