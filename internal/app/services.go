package app

import (
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/service"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/service/updatemetrics"
)

type services struct {
	UpdateMetricsService service.UpdateMetricsService
}

func NewServices(repos *repositories) *services {
	return &services{
		UpdateMetricsService: updatemetrics.NewUpdateMetrics(repos.MemStorage),
	}
}
