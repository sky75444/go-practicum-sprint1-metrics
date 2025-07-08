package app

import (
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/handler"
)

type handlers struct {
	counterHandler *handler.UpdateCounterHandler
	gaugeHandler   *handler.UpdateGaugeHandler
	errorHandler   *handler.ErrorHandler
	getHandler     *handler.GetHandler
}

func NewHandlers(services *services) *handlers {
	return &handlers{
		counterHandler: handler.NewUpdateCounterHandler(services.UpdateMetricsService),
		gaugeHandler:   handler.NewUpdateGaugeHandler(services.UpdateMetricsService),
		errorHandler:   handler.NewErrorHandler(),
		getHandler:     handler.NewGetHandler(services.UpdateMetricsService),
	}
}
