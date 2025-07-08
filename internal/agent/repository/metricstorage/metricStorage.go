package metricstorage

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
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

func (ms *metricStorage) StoreGaugeMetrics(m model.MetricCollection, c *resty.Client) error {
	for k, v := range m.GaugeMetrics {
		req, err := createReq(k, MetricGaugeStorageEndpoint, v, c)
		if err != nil {
			return err
		}

		r, err := req.Send()

		if err != nil {
			return err
		}

		if r.StatusCode() != http.StatusOK {
			return fmt.Errorf("%s", r.Status())
		}
	}

	return nil
}

func (ms *metricStorage) StoreCounterMetrics(m model.MetricCollection, c *resty.Client) error {
	for k, v := range m.CountMetrics {
		req, err := createReq(k, MetricCounterStorageEndpoint, v, c)
		if err != nil {
			return err
		}

		r, err := req.Send()

		if err != nil {
			return err
		}

		if r.StatusCode() != http.StatusOK {
			return fmt.Errorf("%s", r.Status())
		}
	}

	return nil
}

func createReq(memName, memTypeEndpoint string, memValue uint64, c *resty.Client) (*resty.Request, error) {
	metricStorageURL := fmt.Sprintf("%s:%s/%s", MetricServerHost, MetricServerPort, memTypeEndpoint)
	endpoint := fmt.Sprintf("%s/%s/%d/", metricStorageURL, memName, memValue)

	req := c.R()
	req.Method = http.MethodPost
	req.Header.Add("Content-Type", "text/plain")
	req.URL = endpoint

	return req, nil
}
