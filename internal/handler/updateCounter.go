package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/sky75444/go-practicum-sprint1-metrics/internal/service"
)

type UpdateCounterHandler struct {
	updateMetricsService service.UpdateMetricsService
}

func NewUpdateCounterHandler(umService service.UpdateMetricsService) *UpdateCounterHandler {
	return &UpdateCounterHandler{
		updateMetricsService: umService,
	}
}

func (c *UpdateCounterHandler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			// http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if r.Header.Get("Content-Type") != "text/plain" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			// http.Error(w, "Content-Type not allowed", http.StatusMethodNotAllowed)
			return
		}

		metricValueStr := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
		urlNameValue := r.URL.Path[16:]
		metricName := urlNameValue[:strings.LastIndex(urlNameValue, "/")]

		if metricName == "" {
			w.WriteHeader(http.StatusNotFound)
			// http.Error(w, "metric name is required", http.StatusNotFound)
			return
		}

		value, err := strconv.ParseInt(metricValueStr, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			// http.Error(w, "invalid counter value", http.StatusBadRequest)
			return
		}

		if err := c.updateMetricsService.UpdateCounter(metricName, value); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Metric updated"))
	}
}
