package repository

type MemStorage interface {
	UpdateGauge(name string, value float64) error
	UpdateCounter(name string, value int64) error
	GetGauge(name string) (float64, error)
	GetCounter(name string) (int64, error)
	GetAll() (string, error)
	// StoreMetricsToFile(errChan chan error)
	StoreMetricsToFile() error
	SaveDataToFile() error
}
