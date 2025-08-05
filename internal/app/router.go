package app

import (
	"net/http"

	"github.com/sky75444/go-practicum-sprint1-metrics/internal/handler"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/logger"

	"github.com/go-chi/chi"
)

type router struct {
	R       chi.Router
	Srv     *http.Server
	RunAddr string
}

func NewRouter(runAddr string, handlers *handlers) *router {
	r := router{
		R: handler.NewChiMux(
			handlers.errorHandler,
			handlers.counterHandler,
			handlers.gaugeHandler,
			handlers.getHandler,
			handlers.updateHandler,
			handlers.valueHandler,
			handlers.healthHandler,
		),
		RunAddr: runAddr,
	}

	r.Srv = &http.Server{
		Addr:    r.RunAddr,
		Handler: r.R,
	}

	return &r
}

func (r *router) Start() {
	defer logger.ZLog.Sync()
	sl := logger.ZLog.Sugar()

	if err := r.Srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		sl.Fatalw(err.Error(), "event", "start server")
	}
}
