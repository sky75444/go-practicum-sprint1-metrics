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
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if r.Header.Get("Content-Type") != "text/plain" {
			http.Error(w, "Content-Type not allowed", http.StatusMethodNotAllowed)
			return
		}

		if len(r.URL.Path) == strings.LastIndex(r.URL.Path, "/")+1 {
			http.Error(w, "metric value is required", http.StatusNotFound)
			return
		}

		metricValueStr := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
		urlNameValue := r.URL.Path[16:]
		metricName := urlNameValue[:strings.LastIndex(urlNameValue, "/")]

		if metricName == "" {
			http.Error(w, "metric name is required", http.StatusNotFound)
			return
		}

		value, err := strconv.ParseInt(metricValueStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid counter value", http.StatusBadRequest)
			return
		}

		if err := c.updateMetricsService.UpdateCounter(metricName, value); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Metric updated"))
	}
}
