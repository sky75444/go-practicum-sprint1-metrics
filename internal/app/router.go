package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/handler"
)

type router struct {
	R chi.Router
}

func NewRouter(handlers *handlers) *router {
	return &router{
		R: handler.NewChiMux(
			handlers.errorHandler,
			handlers.counterHandler,
			handlers.gaugeHandler,
			handlers.getAllHandler,
		),
	}
}

func (r *router) Start() {
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r.R))
}
