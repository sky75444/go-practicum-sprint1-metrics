package app

type DI struct {
	Repositories *repositories
	Services     *services
	Handlers     *handlers

	Router *router
}

func NewDI() *DI {
	return &DI{}
}

func (d *DI) Init() {
	d.Repositories = NewRepositories()
	d.Services = NewServices(d.Repositories)
	d.Handlers = NewHandlers(d.Services)
	d.Router = NewRouter(d.Handlers)
}
