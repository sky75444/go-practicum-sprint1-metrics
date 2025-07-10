package metricstorage

import (
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateReq(t *testing.T) {
	type want struct {
		reqMethod   string
		reqURL      string
		contentType string
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
				reqURL:      "http://localhost:8080/update/gauge/gauge1/123/",
				contentType: "text/plain",
			},
		},
		{
			name: "correct counter",
			memCollection: model.MetricCollection{
				CountMetrics: map[string]uint64{"counter1": 123},
			},
			want: want{
				reqMethod:   http.MethodPost,
				reqURL:      "http://localhost:8080/update/counter/counter1/123/",
				contentType: "text/plain",
			},
		},
	}

	client := resty.New()
	serverAddr := "http://localhost:8080"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.memCollection.GaugeMetrics {
				req, err := createReq(serverAddr, k, GaugeEndpoint, v, client)

				assert.NoError(t, err)
				assert.Equal(t, tt.want.reqURL, req.URL)
				assert.Equal(t, tt.want.reqMethod, req.Method)
				assert.Equal(t, tt.want.contentType, req.Header.Get("Content-Type"))
			}
			for k, v := range tt.memCollection.CountMetrics {
				req, err := createReq(serverAddr, k, CounterEndpoint, v, client)

				assert.NoError(t, err)
				assert.Equal(t, tt.want.reqURL, req.URL)
				assert.Equal(t, tt.want.reqMethod, req.Method)
				assert.Equal(t, tt.want.contentType, req.Header.Get("Content-Type"))
			}
		})
	}
}
