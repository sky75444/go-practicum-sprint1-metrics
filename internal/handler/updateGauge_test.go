package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sky75444/go-practicum-sprint1-metrics/internal/repository/memstorage"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/service/updatemetrics"
	"github.com/stretchr/testify/assert"
)

func TestGaugeHandle(t *testing.T) {
	type want struct {
		code int
	}
	tests := []struct {
		name   string
		reqURL string
		want   want
	}{
		{
			name:   "correct gauge",
			reqURL: "http://localhost:8080/update/gauge/gauge1/123/",
			want: want{
				code: 200,
			},
		},
		{
			name:   "invalid gauge value",
			reqURL: "http://localhost:8080/update/gauge/gauge1/none/",
			want: want{
				code: 400,
			},
		},
		{
			name:   "empty gauge value",
			reqURL: "http://localhost:8080/update/gauge/gauge1/",
			want: want{
				code: 404,
			},
		},
		{
			name:   "empty gauge name",
			reqURL: "http://localhost:8080/update/gauge//123",
			want: want{
				code: 404,
			},
		},
	}

	gh := NewUpdateGaugeHandler(updatemetrics.NewUpdateMetrics(memstorage.NewMemStorage()))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, tt.reqURL, nil)
			req.Header.Add("Content-Type", "text/plain")
			w := httptest.NewRecorder()
			h := http.HandlerFunc(gh.GaugeHandle())

			h(w, req)

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.want.code, res.StatusCode)
		})
	}
}
