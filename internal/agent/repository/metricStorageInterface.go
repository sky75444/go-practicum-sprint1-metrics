package repository

import (
	"net/http"

	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/model"
)

type MetricRepo interface {
	StoreGaugeMetrics(m model.MetricCollection, c *http.Client) error
	StoreCounterMetrics(m model.MetricCollection, c *http.Client) error
}
