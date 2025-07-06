package metricstorage

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/model"
)

const (
	MetricGaugeStorageEndpoint   = "update/gauge"
	MetricCounterStorageEndpoint = "update/counter"
	MetricServerHost             = "http://localhost"
	MetricServerPort             = "8080"
)

type metricStorage struct{}

func NewMetricStorage() *metricStorage {
	return &metricStorage{}
}

func (ms *metricStorage) StoreGaugeMetrics(m model.MetricCollection, c *http.Client) error {
	for k, v := range m.GaugeMetrics {
		req, err := createReq(k, MetricGaugeStorageEndpoint, v)
		if err != nil {
			return err
		}

		req.Header.Add("Content-Type", "text/plain")

		r, err := c.Do(req)
		if err != nil {
			return err
		}

		defer r.Body.Close()

		if r.StatusCode != http.StatusOK {
			return fmt.Errorf("%s", r.Status)
		}
	}

	return nil
}

func (ms *metricStorage) StoreCounterMetrics(m model.MetricCollection, c *http.Client) error {
	for k, v := range m.CountMetrics {
		req, err := createReq(k, MetricCounterStorageEndpoint, v)
		if err != nil {
			return err
		}

		r, err := c.Do(req)
		if err != nil {
			return err
		}

		defer r.Body.Close()

		if r.StatusCode != http.StatusOK {
			return fmt.Errorf("%s", r.Status)
		}

	}

	return nil
}

func createReq(memName, memTypeEndpoint string, memValue uint64) (*http.Request, error) {
	metricStorageURL := fmt.Sprintf("%s:%s/%s", MetricServerHost, MetricServerPort, memTypeEndpoint)
	endpoint := fmt.Sprintf("%s/%s/%d/", metricStorageURL, memName, memValue)

	req, err := http.NewRequest(http.MethodPost, endpoint, strings.NewReader(""))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "text/plain")

	return req, nil
}
