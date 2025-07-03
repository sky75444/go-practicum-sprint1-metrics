package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/sky75444/go-practicum-sprint1-metrics/internal/service"
)

type UpdateGaugeHandler struct {
	updateMetricsService service.UpdateMetricsService
}

func NewUpdateGaugeHandler(umService service.UpdateMetricsService) *UpdateGaugeHandler {
	return &UpdateGaugeHandler{
		updateMetricsService: umService,
	}
}

func (g *UpdateGaugeHandler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if r.Header.Get("Content-Type") != "text/plain" {
			http.Error(w, "Content-Type not allowed", http.StatusMethodNotAllowed)
			return
		}

		metricValueStr := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
		urlNameValue := r.URL.Path[14:]
		metricName := urlNameValue[:strings.LastIndex(urlNameValue, "/")]

		if metricName == "" {
			http.Error(w, "metric name is required", http.StatusNotFound)
			return
		}

		if metricValueStr == "" {
			http.Error(w, "metric value is required", http.StatusNotFound)
			return
		}

		value, err := strconv.ParseFloat(metricValueStr, 64)
		if err != nil {
			http.Error(w, "invalid gauge value", http.StatusBadRequest)
			return
		}

		if err := g.updateMetricsService.UpdateGauge(metricName, value); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Metric updated")
	}
}
