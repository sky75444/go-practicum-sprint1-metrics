package app

import (
	"net/http"

	"github.com/sky75444/go-practicum-sprint1-metrics/internal/handler"
)

type router struct {
	Mux *http.ServeMux
}

func NewRouter(handlers *handlers) *router {
	mux := http.NewServeMux()
	mux.HandleFunc("/update/counter/", handlers.counterHandler.CounterHandle())
	mux.HandleFunc("/update/gauge/", handlers.gaugeHandler.GaugeHandle())
	mux.HandleFunc("/", handler.ErrorHandler)

	return &router{
		Mux: mux,
	}
}
