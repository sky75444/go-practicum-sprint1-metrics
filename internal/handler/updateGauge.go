package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	models "github.com/sky75444/go-practicum-sprint1-metrics.git/internal/model"
)

func UpdateGaugeHandler(m *models.MemStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// if r.Method != http.MethodPost {
		// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		// 	return
		// }

		// if r.Header.Get("Content-Type") != "text/plain" {
		// 	http.Error(w, "Content-Type not allowed", http.StatusMethodNotAllowed)
		// 	return
		// }

		metricValueStr := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
		metricName := r.URL.Path[14:strings.LastIndex(r.URL.Path, "/")]

		if metricName == "" {
			http.Error(w, "metric name is required", http.StatusNotFound)
			return
		}

		value, err := strconv.ParseFloat(metricValueStr, 64)
		if err != nil {
			http.Error(w, "invalid gauge value", http.StatusBadRequest)
			return
		}

		m.UpdateGauge(metricName, value)

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Metric updated")
	}
}
