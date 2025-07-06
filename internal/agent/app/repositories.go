package app

import (
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/repository"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/repository/metricstorage"
)

type repositories struct {
	MetricStorage repository.MetricRepo
}

func NewRepositories() *repositories {
	return &repositories{
		MetricStorage: metricstorage.NewMetricStorage(),
	}
}
