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

func (g *UpdateGaugeHandler) GaugeHandle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if r.Header.Get("Content-Type") != "text/plain" {
			http.Error(w, "Content-Type not allowed", http.StatusMethodNotAllowed)
			return
		}

		correctPath := r.URL.Path
		if len(r.URL.Path) == strings.LastIndex(r.URL.Path, "/")+1 {
			correctPath = r.URL.Path[:strings.LastIndex(r.URL.Path, "/")]
		}

		mNameValue := r.URL.Path[14:strings.LastIndex(r.URL.Path, "/")]
		if strings.LastIndex(mNameValue, "/") < 0 {
			http.Error(w, "metric name/value is required", http.StatusNotFound)
			return
		}

		metricName := correctPath[14:strings.LastIndex(correctPath, "/")]

		metricValueStr := correctPath[strings.LastIndex(correctPath, "/")+1:]
		value, err := strconv.ParseFloat(metricValueStr, 64)
		if err != nil {
			http.Error(w, "invalid gauge value", http.StatusBadRequest)
			return
		}

		if err := g.updateMetricsService.UpdateGauge(metricName, value); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		fmt.Println("Gauge metric updated - " + metricName)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Metric updated"))
	}
}
