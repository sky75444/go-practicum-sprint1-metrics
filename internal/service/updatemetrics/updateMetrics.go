package updatemetrics

import (
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/repository"
)

type updateMetrics struct {
	repo repository.MemStorage
}

func NewUpdateMetrics(repo repository.MemStorage) *updateMetrics {
	return &updateMetrics{
		repo: repo,
	}
}

func (u *updateMetrics) UpdateGauge(metricName string, metricValue float64) error {
	//some logic
	if err := u.repo.UpdateGauge(metricName, metricValue); err != nil {
		return err
	}

	return nil
}

func (u *updateMetrics) UpdateCounter(metricName string, metricValue int64) error {
	//some logic
	if err := u.repo.UpdateCounter(metricName, metricValue); err != nil {
		return err
	}

	return nil
}
