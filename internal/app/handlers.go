package app

import (
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/handler"
)

type handlers struct {
	counterHandler *handler.UpdateCounterHandler
	gaugeHandler   *handler.UpdateGaugeHandler
	errorHandler   *handler.ErrorHandler
	getHandler     *handler.GetHandler
	updateHandler  *handler.UpdateHandler
	valueHandler   *handler.ValueHandler
	healthHandler  *handler.HealthHandler
}

func NewHandlers(services *services) *handlers {
	return &handlers{
		errorHandler:   handler.NewErrorHandler(),
		counterHandler: handler.NewUpdateCounterHandler(services.UpdateMetricsService),
		gaugeHandler:   handler.NewUpdateGaugeHandler(services.UpdateMetricsService),
		getHandler:     handler.NewGetHandler(services.UpdateMetricsService),
		updateHandler:  handler.NewUpdateHandler(services.UpdateMetricsService),
		valueHandler:   handler.NewValueHandler(services.UpdateMetricsService),
		healthHandler:  handler.NewHealthHandler(),
	}
}
