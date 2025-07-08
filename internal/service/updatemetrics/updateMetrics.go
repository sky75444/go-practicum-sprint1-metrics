package updatemetrics

import (
	"fmt"

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

func (u *updateMetrics) GetCounter(metricName string) (counterValue string, err error) {
	if metricName == "" {
		return "", fmt.Errorf("counter name is empty")
	}

	v, err := u.repo.GetCounter(metricName)
	if err != nil {
		return "", err
	}
	return v, nil
}

func (u *updateMetrics) GetGauge(metricName string) (gaugeValue string, err error) {
	if metricName == "" {
		return "", fmt.Errorf("gauge name is empty")
	}

	v, err := u.repo.GetGauge(metricName)
	if err != nil {
		return "", err
	}
	return v, nil
}

func (u *updateMetrics) GetAll() (string, error) {
	metrics, err := u.repo.GetAll()
	if err != nil {
		return "", err
	}

	return metrics, nil
}
