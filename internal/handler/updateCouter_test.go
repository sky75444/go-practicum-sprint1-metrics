package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sky75444/go-practicum-sprint1-metrics/internal/repository/memstorage"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/service/updatemetrics"
	"github.com/stretchr/testify/assert"
)

func TestCounterHandle(t *testing.T) {
	type want struct {
		code int
	}
	tests := []struct {
		name   string
		reqURL string
		want   want
	}{
		{
			name:   "correct counter",
			reqURL: "http://localhost:8080/update/counter/counter1/123/",
			want: want{
				code: 200,
			},
		},
		{
			name:   "invalid counter value",
			reqURL: "http://localhost:8080/update/counter/counter1/none/",
			want: want{
				code: 400,
			},
		},
		{
			name:   "empty counter value",
			reqURL: "http://localhost:8080/update/counter/counter1/",
			want: want{
				code: 404,
			},
		},
		{
			name:   "empty counter name",
			reqURL: "http://localhost:8080/update/counter//123",
			want: want{
				code: 404,
			},
		},
	}

	ch := NewUpdateCounterHandler(updatemetrics.NewUpdateMetrics(memstorage.NewMemStorage()))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, tt.reqURL, nil)
			req.Header.Add("Content-Type", "text/plain")
			w := httptest.NewRecorder()
			h := http.Handler(http.StripPrefix("/update/counter/", ch.CounterHandle()))

			h.ServeHTTP(w, req)

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.want.code, res.StatusCode)
		})
	}
}
