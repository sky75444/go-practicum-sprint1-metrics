package service

import "net/http"

type MetricCollector interface {
	EndlessCollectMetrics(c *http.Client) error
}
