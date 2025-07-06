package app

import "net/http"

type DI struct {
	Repositories *repositories
	Services     *services
	Client       *http.Client
}

func NewDI() *DI {
	return &DI{}
}

func (d *DI) Init() {
	d.Repositories = NewRepositories()
	d.Services = NewServices(d.Repositories)
	d.Client = &http.Client{}
}
