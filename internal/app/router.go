package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/handler"
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
		),
		RunAddr: runAddr,
	}
}

func (r *router) Start() {
	fmt.Printf("Server started at %s\n", r.RunAddr)
	log.Fatal(http.ListenAndServe(r.RunAddr, r.R))
}
