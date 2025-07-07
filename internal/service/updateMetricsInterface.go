package service

type UpdateMetricsService interface {
	UpdateGauge(metricName string, metricValue float64) error
	UpdateCounter(metricName string, metricValue int64) error
	GetCounter(metricName string) (counterValue string, err error)
	GetGauge(metricName string) (gaugeValue string, err error)
	GetAll() (string, error)
}
