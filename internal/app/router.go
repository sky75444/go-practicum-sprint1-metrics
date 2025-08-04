package app

import (
	"net/http"

	"github.com/sky75444/go-practicum-sprint1-metrics/internal/handler"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/logger"

	"github.com/go-chi/chi"
)

type router struct {
	R       chi.Router
	RunAddr string
}

func NewRouter(runAddr string, handlers *handlers) *router {
	return &router{
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
}

func (r *router) Start() {
	defer logger.ZLog.Sync()
	sl := logger.ZLog.Sugar()

	sl.Infow(
		"Starting server",
		"addr", r.RunAddr,
	)

	if err := http.ListenAndServe(r.RunAddr, r.R); err != nil {
		sl.Fatalw(err.Error(), "event", "start server")
	}
}
