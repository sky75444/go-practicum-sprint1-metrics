package app

import (
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/logger"
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
	defer logger.ZLog.Sync()
	sl := logger.ZLog.Sugar()

	repos, err := NewRepositories(config)
	if err != nil {
		sl.Errorw("error init repos", logger.ZError(err))
		panic(err)
	}

	d.Repositories = repos
	d.Services = NewServices(d.Repositories)
	d.Handlers = NewHandlers(d.Services)

	d.Router = NewRouter(config.RunAddr, d.Handlers)
}

func (d *DI) Start() {
	defer logger.ZLog.Sync()
	sl := logger.ZLog.Sugar()

	go func() {
		sl.Debugw("EndlessStoreMetricsToFile ig go()")

		if err := d.Services.UpdateMetricsService.EndlessStoreMetricsToFile(); err != nil {
			sl.Fatalw("error while store metrics", logger.ZError(err))
			panic(err)
		}
	}()

	d.Router.Start()
}
