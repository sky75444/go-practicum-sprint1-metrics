package app

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

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

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		sl.Debugw("EndlessStoreMetricsToFile ig go()")

		if err := d.Services.UpdateMetricsService.EndlessStoreMetricsToFile(ctx); err != nil {
			sl.Fatalw("error while store metrics", logger.ZError(err))
			panic(err)
		}
	}()

	go func() {
		d.Router.Start()
	}()

	<-sigs
	cancel()
	wg.Wait()
	sl.Infow("Shutting down server...")

	if err := d.Services.UpdateMetricsService.SaveDataToFile(); err != nil {
		sl.Fatalw("error saving metrics to file", logger.ZError(err))
		panic(err)
	}

	sl.Infow("data saved")

	if err := d.Router.Srv.Shutdown(ctx); err != nil {
		sl.Errorw("server forced shutdown", logger.ZError(err))
	}

	sl.Infow("server stopped")
}
