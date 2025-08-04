package metricstorage

import (
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateReq(t *testing.T) {
	type body struct {
		ID    string
		MType string
		Delta int64
		Value float64
	}
	type want struct {
		reqMethod   string
		reqURL      string
		contentType string
		body        body
	}
	tests := []struct {
		name          string
		req           string
		memCollection model.MetricCollection
		want          want
	}{
		{
			name: "correct gauge",
			memCollection: model.MetricCollection{
				GaugeMetrics: map[string]uint64{"gauge1": 123},
			},
			want: want{
				reqMethod:   http.MethodPost,
				reqURL:      "http://localhost:8080/update/",
				contentType: "application/json",
				body: body{
					ID:    "gauge1",
					MType: "gauge",
					Value: float64(123),
					Delta: 0,
				},
			},
		},
		{
			name: "correct counter",
			memCollection: model.MetricCollection{
				CountMetrics: map[string]uint64{"counter1": 123},
			},
			want: want{
				reqMethod:   http.MethodPost,
				reqURL:      "http://localhost:8080/update/",
				contentType: "application/json",
				body: body{
					ID:    "counter1",
					MType: "counter",
					Delta: int64(123),
					Value: 0,
				},
			},
		},
	}

	client := resty.New()
	serverAddr := "http://localhost:8080"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for mn, mv := range tt.memCollection.GaugeMetrics {
				var err error
				req, err := createUpdateReqWithBody(serverAddr, mn, model.Gauge, mv, client)
				assert.NoError(t, err)

				m := req.Body.(model.Metrics)
				if m.Delta == nil {
					d := int64(0)
					m.Delta = &d
				}
				if m.Value == nil {
					d := float64(0)
					m.Value = &d
				}

				assert.Equal(t, tt.want.reqURL, req.URL)
				assert.Equal(t, tt.want.reqMethod, req.Method)
				assert.Equal(t, tt.want.contentType, req.Header.Get("Content-Type"))
				assert.Equal(t, tt.want.body.ID, m.ID)
				assert.Equal(t, tt.want.body.MType, m.MType)
				assert.Equal(t, tt.want.body.Delta, *m.Delta)
				assert.Equal(t, tt.want.body.Value, *m.Value)
			}
			for mn, mv := range tt.memCollection.CountMetrics {
				var err error
				req, err := createUpdateReqWithBody(serverAddr, mn, model.Counter, mv, client)

				assert.NoError(t, err)

				m := req.Body.(model.Metrics)
				if m.Delta == nil {
					d := int64(0)
					m.Delta = &d
				}
				if m.Value == nil {
					d := float64(0)
					m.Value = &d
				}

				assert.NoError(t, err)
				assert.Equal(t, tt.want.reqURL, req.URL)
				assert.Equal(t, tt.want.reqMethod, req.Method)
				assert.Equal(t, tt.want.contentType, req.Header.Get("Content-Type"))
				assert.Equal(t, tt.want.body.ID, m.ID)
				assert.Equal(t, tt.want.body.MType, m.MType)
				assert.Equal(t, tt.want.body.Delta, *m.Delta)
				assert.Equal(t, tt.want.body.Value, *m.Value)
			}
		})
	}
}
