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
)

type metricStorage struct {
	serverAddr string
}

func NewMetricStorage(serverAddr string) *metricStorage {
	return &metricStorage{
		serverAddr: serverAddr,
	}
}

func (ms *metricStorage) StoreGaugeMetrics(m model.MetricCollection, c *resty.Client) error {
	for k, v := range m.GaugeMetrics {
		req, err := createReq(ms.serverAddr, k, MetricGaugeStorageEndpoint, v, c)
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
		req, err := createReq(ms.serverAddr, k, MetricCounterStorageEndpoint, v, c)
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

func createReq(serverAddr, memName, memTypeEndpoint string, memValue uint64, c *resty.Client) (*resty.Request, error) {
	metricStorageURL := fmt.Sprintf("%s/%s", serverAddr, memTypeEndpoint)
	endpoint := fmt.Sprintf("%s/%s/%d/", metricStorageURL, memName, memValue)

	if endpoint[:4] != "http" {
		endpoint = fmt.Sprintf("http://%s", endpoint)
	}

	req := c.R()
	req.Method = http.MethodPost
	req.Header.Add("Content-Type", "text/plain")
	req.URL = endpoint

	return req, nil
}
