package service

import (
	"github.com/go-resty/resty/v2"
)

type MetricCollector interface {
	EndlessCollectMetrics(c *resty.Client) error
}
