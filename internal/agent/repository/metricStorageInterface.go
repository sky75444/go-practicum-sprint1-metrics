package repository

import (
	"github.com/go-resty/resty/v2"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/model"
)

type MetricRepo interface {
	StoreMetrics(m model.MetricCollection, c *resty.Client) error
	ServerHealthCheck(c *resty.Client) (bool, error)
}
