package metricstorage

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/model"
)

const (
	MetricGaugeStorageEndpoint = "update/guage"
	MetricServerHost           = "http://localhost"
	MetricServerPort           = "8080"
)

type metricStorage struct{}

func NewMetricStorage() *metricStorage {
	return &metricStorage{}
}

func (ms *metricStorage) StoreGaugeMetrics(m model.MetricCollection, c *http.Client) error {
	metricStorageURL := fmt.Sprintf("%s:%s/%s", MetricServerHost, MetricServerPort, MetricGaugeStorageEndpoint)

	for k, v := range m.GaugeMetrics {
		endpoint := fmt.Sprintf("%s/%s/%d/", metricStorageURL, k, v)

		req, err := http.NewRequest(http.MethodPost, endpoint, strings.NewReader(""))
		if err != nil {
			return err
		}

		req.Header.Add("Content-Type", "text/plain")

		if _, err := c.Do(req); err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

func (ms *metricStorage) StoreCounterMetrics(m model.MetricCollection, c *http.Client) error {
	metricStorageURL := fmt.Sprintf("%s:%s/%s", MetricServerHost, MetricServerPort, MetricGaugeStorageEndpoint)

	for k, v := range m.CountMetrics {
		endpoint := fmt.Sprintf("%s/%s/%d/", metricStorageURL, k, v)

		req, err := http.NewRequest(http.MethodPost, endpoint, strings.NewReader(""))
		if err != nil {
			return err
		}

		req.Header.Add("Content-Type", "text/plain")

		if _, err := c.Do(req); err != nil {
			return err
		}
	}

	return nil
}
