package app

import (
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/serverconfig"
)

type DI struct {
	Repositories *repositories
	Services     *services
	Handlers     *handlers

	Router *router
}

func NewDI() *DI {
	return &DI{}
}

func (d *DI) Init(config *serverconfig.Config) {
	d.Repositories = NewRepositories()
	d.Services = NewServices(d.Repositories)
	d.Handlers = NewHandlers(d.Services)

	d.Router = NewRouter(config.RunAddr, d.Handlers)
}
