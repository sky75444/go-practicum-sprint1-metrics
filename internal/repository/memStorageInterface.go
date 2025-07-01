package repository

type MemStorage interface {
	UpdateGauge(name string, value float64) error
	UpdateCounter(name string, value int64) error
}
