package service

type MetricCollector interface {
	EndlessCollectMetrics() error
}
