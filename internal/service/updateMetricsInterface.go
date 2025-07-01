package service

type UpdateMetricsService interface {
	UpdateGauge(metricName string, metricValue float64) error
	UpdateCounter(metricName string, metricValue int64) error
}
