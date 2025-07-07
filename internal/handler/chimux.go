package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewChiMux(
	errorHandler *ErrorHandler,
	counterHandler *UpdateCounterHandler,
	gaugeHandler *UpdateGaugeHandler,
	getAllHander *GetAllHandler,
) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.AllowContentType("text/plain"))

	r.Get("/", getAllHander.GetAll())
	r.Route("/update", func(r chi.Router) {
		r.Post("/", errorHandler.BadRequest)
		r.Get("/", errorHandler.BadRequest)

		r.Route("/counter", func(r chi.Router) {
			r.Post("/", errorHandler.BadRequest)
			r.Get("/", errorHandler.BadRequest)
			r.Route("/{counterName}", func(r chi.Router) {
				r.Get("/", counterHandler.GetCounterHandle())
				r.Post("/", errorHandler.NotFound)
				r.Route("/{counterValue}", func(r chi.Router) {
					r.Post("/", counterHandler.CounterHandle())
				})

			})
		})

		r.Route("/gauge", func(r chi.Router) {
			r.Post("/", errorHandler.BadRequest)
			r.Get("/", errorHandler.BadRequest)
			r.Route("/{gaugeName}", func(r chi.Router) {
				r.Get("/", gaugeHandler.GetGaugeHandle())
				r.Post("/", errorHandler.NotFound)
				r.Route("/{gaugeValue}", func(r chi.Router) {
					r.Post("/", gaugeHandler.GaugeHandle())
				})

			})
		})
	})

	return r
}
