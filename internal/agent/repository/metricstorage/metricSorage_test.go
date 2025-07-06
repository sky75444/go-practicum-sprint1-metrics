package metricstorage

import (
	"net/http"
	"testing"

	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateReq(t *testing.T) {
	type want struct {
		reqMethod   string
		reqUrl      string
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
				reqUrl:      "http://localhost:8080/update/gauge/gauge1/123/",
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
				reqUrl:      "http://localhost:8080/update/counter/counter1/123/",
				contentType: "text/plain",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.memCollection.GaugeMetrics {
				req, err := createReq(k, MetricGaugeStorageEndpoint, v)

				assert.NoError(t, err)
				assert.Equal(t, tt.want.reqUrl, req.URL.String())
				assert.Equal(t, tt.want.reqMethod, req.Method)
				assert.Equal(t, tt.want.contentType, req.Header.Get("Content-Type"))
			}
			for k, v := range tt.memCollection.CountMetrics {
				req, err := createReq(k, MetricCounterStorageEndpoint, v)

				assert.NoError(t, err)
				assert.Equal(t, tt.want.reqUrl, req.URL.String())
				assert.Equal(t, tt.want.reqMethod, req.Method)
				assert.Equal(t, tt.want.contentType, req.Header.Get("Content-Type"))
			}
		})
	}
}
