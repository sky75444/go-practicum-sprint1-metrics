package repository

import (
	"github.com/go-resty/resty/v2"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/model"
)

type MetricRepo interface {
	StoreGaugeMetrics(m model.MetricCollection, c *resty.Client) error
	StoreCounterMetrics(m model.MetricCollection, c *resty.Client) error
}
