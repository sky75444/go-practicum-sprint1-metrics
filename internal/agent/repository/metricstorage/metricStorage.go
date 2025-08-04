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

func (ms *metricStorage) StoreMetrics(m model.MetricCollection, c *resty.Client) error {
	for mn, mv := range m.CountMetrics {
		req, err := createUpdateReqWithBody(ms.serverAddr, mn, model.Counter, mv, c)
		if err != nil {
			return err
		}

		if err := send(req); err != nil {
			return err
		}
	}

	for mn, mv := range m.GaugeMetrics {
		req, err := createUpdateReqWithBody(ms.serverAddr, mn, model.Gauge, mv, c)
		if err != nil {
			return err
		}

		if err := send(req); err != nil {
			return err
		}
	}

	return nil
}

func (ms *metricStorage) ServerHealthCheck(c *resty.Client) (bool, error) {
	req := c.R()
	req.Method = http.MethodGet
	req.URL = craftURL(ms.serverAddr, HelthEndpoint)

	if err := send(req); err != nil {
		return false, err
	}

	return true, nil
}

func createUpdateReqWithBody(
	serverAddr, metricName, metricType string,
	metricVal uint64,
	c *resty.Client,
) (*resty.Request, error) {
	var m model.Metrics

	m.ID = metricName

	switch {
	case metricType == model.Gauge:
		m.MType = model.Gauge
		d := float64(metricVal)
		m.Value = &d
	default:
		m.MType = model.Counter
		d := int64(metricVal)
		m.Delta = &d
	}

	req := c.R()
	req.Method = http.MethodPost
	req.SetHeader("Content-Type", "application/json")
	req.URL = craftURL(serverAddr, MetricStorageEndpoint)
	// req.SetBody(m)
	req.Body = m

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

func craftURL(srvAddr, edpoint string) string {
	if len(srvAddr) == 5 {
		//Если длина 5, это значит что хост не указан. А для агента важно знать хост
		srvAddr = fmt.Sprintf("http://localhost%s", srvAddr)
	}

	edpointPath := fmt.Sprintf("%s/%s", srvAddr, edpoint)
	if edpointPath[:4] != "http" {
		edpointPath = fmt.Sprintf("http://%s", edpointPath)
	}

	return edpointPath
}
