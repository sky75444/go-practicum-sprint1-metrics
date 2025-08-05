package service

type UpdateMetricsService interface {
	UpdateGauge(metricName string, metricValue float64) error
	UpdateCounter(metricName string, metricValue int64) error
	GetCounter(metricName string) (counterValue int64, err error)
	GetGauge(metricName string) (gaugeValue float64, err error)
	GetAll() (string, error)
	// EndlessStoreMetricsToFileAsync(errChan chan error, ctx context.Context)
	EndlessStoreMetricsToFile() error
	SaveDataToFile() error
}
