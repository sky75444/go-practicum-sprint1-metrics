package metricstorage

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/model"
)

const (
	MetricGaugeStorageEndpoint   = "update/gauge"
	MetricCounterStorageEndpoint = "update/counter"
	MetricStorageEndpoint        = "update/"
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

		// fmt.Println("1")
		// reqBody, err := json.Marshal(m)
		// if err != nil {
		// 	fmt.Println("1e")
		// 	return err
		// }

		fmt.Println("2")
		req, err := createUpdateReqWithBody(ms.serverAddr, m, c)
		if err != nil {
			fmt.Println("2e")
			return err
		}

		fmt.Println("3")
		if err := send(req); err != nil {
			fmt.Println(m)
			// fmt.Println(string(reqBody))

			fmt.Println("3e")
			return err
		}
	}

	for mn, mv := range m.GaugeMetrics {
		var m model.Metrics

		m.ID = mn
		m.MType = model.Gauge
		d := float64(mv)
		m.Value = &d

		// fmt.Println("11")
		// reqBody, err := json.Marshal(m)
		// if err != nil {
		// 	fmt.Println("11e")
		// 	return err
		// }

		fmt.Println("22")
		req, err := createUpdateReqWithBody(ms.serverAddr, m, c)
		if err != nil {
			fmt.Println("22e")
			return err
		}

		fmt.Println("33")
		if err := send(req); err != nil {
			fmt.Println("33e")
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

	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req.SetBody(&reqBody)

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
	fmt.Println(req.Method)
	fmt.Println(req.Header.Get("Content-Type"))
	fmt.Println(req.URL)

	r, err := req.Send()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if r.StatusCode() != http.StatusOK {
		fmt.Println(r.Status())
		return fmt.Errorf("%s", r.Status())
	}

	return nil
}
