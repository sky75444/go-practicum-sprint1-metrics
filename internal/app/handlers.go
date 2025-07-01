package app

import (
	"github.com/sky75444/go-practicum-sprint1-metrics.git/internal/handler"
)

type handlers struct {
	counterHandler *handler.UpdateCounterHandler
	gaugeHandler   *handler.UpdateGaugeHandler
}

func NewHandlers(services *services) *handlers {
	return &handlers{
		counterHandler: handler.NewUpdateCounterHandler(services.UpdateMetricsService),
		gaugeHandler:   handler.NewUpdateGaugeHandler(services.UpdateMetricsService),
	}
}
