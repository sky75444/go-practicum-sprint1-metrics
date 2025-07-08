package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sky75444/go-practicum-sprint1-metrics/internal/repository/memstorage"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/service/updatemetrics"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			reqURL: "/update/gauge/gauge1/123/",
			want: want{
				code: 200,
			},
		},
		{
			name:   "invalid gauge value",
			reqURL: "/update/gauge/gauge1/none/",
			want: want{
				code: 400,
			},
		},
		{
			name:   "empty gauge value",
			reqURL: "/update/gauge/gauge1/",
			want: want{
				code: 404,
			},
		},
		{
			name:   "empty gauge name",
			reqURL: "/update/gauge//123",
			want: want{
				code: 404,
			},
		},
		{
			name:   "unknown method",
			reqURL: "/update/unknown/gauge1/123/",
			want: want{
				code: 404,
			},
		},
	}

	ch := NewUpdateCounterHandler(updatemetrics.NewUpdateMetrics(memstorage.NewMemStorage()))
	gh := NewUpdateGaugeHandler(updatemetrics.NewUpdateMetrics(memstorage.NewMemStorage()))
	ga := NewGetAllHandler(updatemetrics.NewUpdateMetrics(memstorage.NewMemStorage()))
	eh := NewErrorHandler()

	srv := httptest.NewServer(NewChiMux(eh, ch, gh, ga))
	defer srv.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, srv.URL+tt.reqURL, nil)
			require.NoError(t, err, "error making HTTP request")

			resp, err := srv.Client().Do(req)
			require.NoError(t, err, "error making HTTP request")
			defer resp.Body.Close()

			assert.Equal(t, tt.want.code, resp.StatusCode)
		})
	}
}
