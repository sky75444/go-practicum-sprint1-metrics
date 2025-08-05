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
			reqURL: "/update/counter/counter1/123/",
			want: want{
				code: 200,
			},
		},
		{
			name:   "invalid counter value",
			reqURL: "/update/counter/counter1/none/",
			want: want{
				code: 400,
			},
		},
		{
			name:   "empty counter value",
			reqURL: "/update/counter/counter1/",
			want: want{
				code: 404,
			},
		},
		{
			name:   "empty counter name",
			reqURL: "/update/counter//123",
			want: want{
				code: 404,
			},
		},
		{
			name:   "unknown method",
			reqURL: "/update/unknown/counter1/123",
			want: want{
				code: 400,
			},
		},
	}

	needRestore := false
	fname := "./savedData.json"
	storeInterval := 10

	mem, err := memstorage.NewMemStorage(fname, needRestore, storeInterval)
	if err != nil {
		panic(err)
	}

	um := updatemetrics.NewUpdateMetrics(mem)

	cu := NewUpdateCounterHandler(um)
	gu := NewUpdateGaugeHandler(um)
	gh := NewGetHandler(um)
	uh := NewUpdateHandler(um)
	vh := NewValueHandler(um)
	hh := NewHealthHandler()
	eh := NewErrorHandler()

	srv := httptest.NewServer(NewChiMux(eh, cu, gu, gh, uh, vh, hh))
	defer srv.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, srv.URL+tt.reqURL, nil)
			require.NoError(t, err, "error making HTTP request")

			req.Header.Add("Content-Type", "text/plain")

			resp, err := srv.Client().Do(req)
			require.NoError(t, err, "error making HTTP request")
			defer resp.Body.Close()

			assert.Equal(t, tt.want.code, resp.StatusCode)
		})
	}
}
