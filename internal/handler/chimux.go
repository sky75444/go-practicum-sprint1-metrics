package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/logger"
)

func NewChiMux(
	errorHandler *ErrorHandler,
	counterHandler *UpdateCounterHandler,
	gaugeHandler *UpdateGaugeHandler,
	getHander *GetHandler,
	updateHandler *UpdateHandler,
	valueHandler *ValueHandler,
) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.AllowContentType("text/plain", "application/json"))

	r.Get("/", logger.WithLogging(getHander.GetAll()))
	r.Post("/", logger.WithLogging(errorHandler.BadRequest))

	nfh := r.NotFoundHandler()

	r.Route("/value", func(r chi.Router) {
		r.Post("/", logger.WithLogging(valueHandler.ValueHandle()))

		r.Route("/{metricType}/{metricName}", func(r chi.Router) {
			r.Get("/", logger.WithLogging(getHander.GetMetric()))
		})
	})

	r.Route("/update", func(r chi.Router) {
		r.Post("/", logger.WithLogging(updateHandler.UpdateHandle()))
		r.NotFound(logger.WithLogging(errorHandler.BadRequest))

		r.Route("/counter", func(r chi.Router) {
			r.Post("/", logger.WithLogging(nfh))
			r.Get("/", logger.WithLogging(nfh))
			r.Route("/{counterName}", func(r chi.Router) {
				r.Post("/", logger.WithLogging(nfh))
				r.Get("/", logger.WithLogging(nfh))
				r.Route("/{counterValue}", func(r chi.Router) {
					r.Post("/", logger.WithLogging(counterHandler.CounterHandle()))
				})
			})
		})

		r.Route("/gauge", func(r chi.Router) {
			r.Post("/", logger.WithLogging(nfh))
			r.Get("/", logger.WithLogging(nfh))
			r.Route("/{gaugeName}", func(r chi.Router) {
				r.Post("/", logger.WithLogging(nfh))
				r.Get("/", logger.WithLogging(nfh))
				r.Route("/{gaugeValue}", func(r chi.Router) {
					r.Post("/", logger.WithLogging(gaugeHandler.GaugeHandle()))
				})
			})
		})
	})

	return r
}
