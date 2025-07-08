package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/service"
)

type GetHandler struct {
	updateMetricsService service.UpdateMetricsService
}

func NewGetHandler(umService service.UpdateMetricsService) *GetHandler {
	return &GetHandler{
		updateMetricsService: umService,
	}
}

func (gh *GetHandler) GetMetric() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metricType := strings.ToLower(chi.URLParam(r, "metricType"))
		if metricType == "" {
			http.Error(w, "counter name is missing", http.StatusNotFound)
			return
		}

		metricName := strings.ToLower(chi.URLParam(r, "metricName"))
		if metricName == "" {
			http.Error(w, "counter name is missing", http.StatusNotFound)
			return
		}

		var metricValStr string

		if metricType == "counter" {
			counterValue, err := gh.updateMetricsService.GetCounter(metricName)
			if err != nil {
				log.Println("metric not found" + " - " + metricName)
				http.Error(w, "metric not found", http.StatusNotFound)
				return
			}

			metricValStr = counterValue
		} else {
			gaugeValue, err := gh.updateMetricsService.GetGauge(metricName)
			if err != nil {
				log.Println("metric not found" + " - " + metricName)
				http.Error(w, "metric not found", http.StatusNotFound)
				return
			}

			metricValStr = gaugeValue
		}

		log.Printf("%s - %s", metricName, metricValStr)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(metricValStr))
	})
}

func (gh *GetHandler) GetAll() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metricsList, err := gh.updateMetricsService.GetAll()

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.Println("generated all metrics list")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(metricsList))
	})
}
