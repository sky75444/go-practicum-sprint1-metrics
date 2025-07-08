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

	nfh := r.NotFoundHandler()

	r.Get("/", getHander.GetAll())
	r.Post("/", errorHandler.BadRequest)

	r.Route("/value", func(r chi.Router) {
		r.Get("/{metricType}/{metricName}/", getHander.GetMetric())
	})

	r.Route("/update", func(r chi.Router) {
		r.Post("/", nfh)
		r.Get("/", nfh)
		r.NotFound(errorHandler.BadRequest)

		r.Route("/counter", func(r chi.Router) {
			r.Post("/", nfh)
			r.Get("/", nfh)
			r.Route("/{counterName}", func(r chi.Router) {
				r.Get("/", nfh)
				r.Post("/", nfh)
				r.Route("/{counterValue}", func(r chi.Router) {
					r.Post("/", counterHandler.CounterHandle())
				})

			})
		})

		r.Route("/gauge", func(r chi.Router) {
			r.Post("/", nfh)
			r.Get("/", nfh)
			r.Route("/{gaugeName}", func(r chi.Router) {
				r.Get("/", nfh)
				r.Post("/", nfh)
				r.Route("/{gaugeValue}", func(r chi.Router) {
					r.Post("/", gaugeHandler.GaugeHandle())
				})
			})
		})
	})

	return r
}
