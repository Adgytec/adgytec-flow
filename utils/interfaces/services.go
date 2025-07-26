package interfaces

import (
	"github.com/go-chi/chi/v5"
)

// IService is used by router mux to init service http rest endpoints
type IService interface {
	ServiceName() string
	BasePath() string
	Router() *chi.Mux
}

// IServiceInit is used by app before creating IService or IServicePC to ensure related data initialization
type IServiceInit interface {
	InitService() error
}

// IServicePC is used for inter-service communication
// each service will have their own PC(procedural call) interface
// this interface is just to tell there are 3 things each service must provide
type IServicePC interface{}
