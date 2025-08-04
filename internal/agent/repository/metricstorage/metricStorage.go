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
	MetricStorageEndpoint        = "update/"
	HelthEndpoint                = "health/"
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

func (ms *metricStorage) StoreMetrics(m model.MetricCollection, c *resty.Client) error {
	for mn, mv := range m.CountMetrics {
		var m model.Metrics

		m.ID = mn
		m.MType = model.Counter
		d := int64(mv)
		m.Delta = &d

		req, err := createUpdateReqWithBody(ms.serverAddr, m, c)
		if err != nil {
			return err
		}

		if err := send(req); err != nil {
			return err
		}
	}

	for mn, mv := range m.GaugeMetrics {
		var m model.Metrics

		m.ID = mn
		m.MType = model.Gauge
		d := float64(mv)
		m.Value = &d

		req, err := createUpdateReqWithBody(ms.serverAddr, m, c)
		if err != nil {
			return err
		}

		if err := send(req); err != nil {
			return err
		}
	}

	return nil
}

func createUpdateReqWithBody(serverAddr string, body model.Metrics, c *resty.Client) (*resty.Request, error) {
	if len(serverAddr) == 5 {
		//Если длина 5, это значит что хост не указан. А для агента важно знать хост
		serverAddr = fmt.Sprintf("http://localhost%s", serverAddr)
	}

	metricStorageURL := fmt.Sprintf("%s/%s", serverAddr, MetricStorageEndpoint)
	if metricStorageURL[:4] != "http" {
		metricStorageURL = fmt.Sprintf("http://%s", metricStorageURL)
	}

	req := c.R()
	req.Method = http.MethodPost
	req.SetHeader("Content-Type", "application/json")
	req.URL = metricStorageURL
	req.SetBody(&body)

	return req, nil
}

func createReq(serverAddr, memName, memTypeEndpoint string, memValue uint64, c *resty.Client) (*resty.Request, error) {
	if len(serverAddr) == 5 {
		//Если длина 5, это значит что хост не указан. А для агента важно знать хост
		serverAddr = fmt.Sprintf("http://localhost%s", serverAddr)
	}

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

func send(req *resty.Request) error {
	r, err := req.Send()
	if err != nil {
		return err
	}

	if r.StatusCode() != http.StatusOK {
		return fmt.Errorf("%s", r.Status())
	}

	return nil
}

func (ms *metricStorage) ServerHealthCheck(c *resty.Client) (bool, error) {
	if len(ms.serverAddr) == 5 {
		//Если длина 5, это значит что хост не указан. А для агента важно знать хост
		ms.serverAddr = fmt.Sprintf("http://localhost%s", ms.serverAddr)
	}

	HelthEndpointURL := fmt.Sprintf("%s/%s", ms.serverAddr, HelthEndpoint)
	if HelthEndpointURL[:4] != "http" {
		HelthEndpointURL = fmt.Sprintf("http://%s", HelthEndpointURL)
	}

	req := c.R()
	req.Method = http.MethodGet
	req.URL = HelthEndpointURL

	r, err := req.Send()
	if err != nil {
		return false, err
	}

	if r.StatusCode() != http.StatusOK {
		return false, fmt.Errorf("%s", r.Status())
	}

	return true, nil
}
