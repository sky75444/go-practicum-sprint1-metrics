package app

import (
	"github.com/go-resty/resty/v2"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/agentconfig"
)

type DI struct {
	Repositories *repositories
	Services     *services
	Client       *resty.Client
}

func NewDI() *DI {
	return &DI{}
}

func (d *DI) Init(config *agentconfig.Config) {
	d.Repositories = NewRepositories(config.MemStorageServerAddr)
	d.Services = NewServices(d.Repositories, config)
	d.Client = resty.New()
}
