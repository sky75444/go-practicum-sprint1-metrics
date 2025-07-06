package handler

import (
	"fmt"
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

func (c *UpdateCounterHandler) CounterHandle() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if r.Header.Get("Content-Type") != "text/plain" {
			http.Error(w, "Content-Type not allowed", http.StatusMethodNotAllowed)
			return
		}

		correctPath := r.URL.String()
		if len(r.URL.Path) == strings.LastIndex(r.URL.String(), "/")+1 {
			correctPath = r.URL.String()[:strings.LastIndex(r.URL.String(), "/")]
		}

		if strings.LastIndex(correctPath, "/") < 0 || len(correctPath) == strings.LastIndex(correctPath, "/")+1 {
			http.Error(w, "metric name/value is required", http.StatusNotFound)
			return
		}

		metricName := correctPath[:strings.LastIndex(correctPath, "/")]

		metricValueStr := correctPath[strings.LastIndex(correctPath, "/")+1:]
		value, err := strconv.ParseInt(metricValueStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid counter value", http.StatusBadRequest)
			return
		}

		if err := c.updateMetricsService.UpdateCounter(metricName, value); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		fmt.Println("Counter metric updated - " + metricName)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Metric updated"))
	})
}
