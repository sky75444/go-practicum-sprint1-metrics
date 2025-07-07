package app

import (
	"github.com/go-resty/resty/v2"
)

type DI struct {
	Repositories *repositories
	Services     *services
	Client       *resty.Client
}

func NewDI() *DI {
	return &DI{}
}

func (d *DI) Init() {
	d.Repositories = NewRepositories()
	d.Services = NewServices(d.Repositories)
	d.Client = resty.New()
}
