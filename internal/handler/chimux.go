package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewChiMux(
	errorHandler *ErrorHandler,
	counterHandler *UpdateCounterHandler,
	gaugeHandler *UpdateGaugeHandler,
	getHander *GetHandler,
) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.AllowContentType("text/plain"))

	r.Get("/", getHander.GetAll())
	r.Post("/", errorHandler.BadRequest)

	nfh := r.NotFoundHandler()

	r.Route("/value", func(r chi.Router) {
		r.Route("/{metricType}/{metricName}", func(r chi.Router) {
			r.Get("/", getHander.GetMetric())
		})
	})

	r.Route("/update", func(r chi.Router) {
		r.NotFound(errorHandler.BadRequest)

		r.Route("/counter/{counterName}", func(r chi.Router) {
			r.Post("/", nfh)
			r.Get("/", nfh)
			r.Route("/{counterValue}", func(r chi.Router) {
				r.Post("/", counterHandler.CounterHandle())
			})

		})

		r.Route("/gauge/{gaugeName}", func(r chi.Router) {
			r.Post("/", nfh)
			r.Get("/", nfh)
			r.Route("/{gaugeValue}", func(r chi.Router) {
				r.Post("/", gaugeHandler.GaugeHandle())
			})

		})
	})

	return r
}
