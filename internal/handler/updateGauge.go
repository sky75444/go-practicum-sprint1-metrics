package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gaugeName := strings.ToLower(chi.URLParam(r, "gaugeName"))
		if gaugeName == "" {
			http.Error(w, "gauge name is missing", http.StatusNotFound)
			return
		}

		gaugeValueStr := strings.ToLower(chi.URLParam(r, "gaugeValue"))
		if gaugeValueStr == "" {
			http.Error(w, "gauge value is missing", http.StatusNotFound)
			return
		}

		value, err := strconv.ParseFloat(gaugeValueStr, 64)
		if err != nil {
			http.Error(w, "invalid gauge value", http.StatusBadRequest)
			return
		}

		if err := g.updateMetricsService.UpdateGauge(gaugeName, value); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		log.Println("gauge metric updated - " + gaugeName)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Metric updated"))
	})
}

// func (g *UpdateGaugeHandler) GetGaugeHandle() http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		gaugeName := strings.ToLower(chi.URLParam(r, "gaugeName"))
// 		if gaugeName == "" {
// 			http.Error(w, "gauge name is missing", http.StatusNotFound)
// 			return
// 		}

// 		gaugeValue, err := g.updateMetricsService.GetGauge(gaugeName)
// 		if err != nil {
// 			log.Println("metric not found" + " - " + gaugeName)
// 			http.Error(w, "metric not found", http.StatusNotFound)
// 			return
// 		}

// 		log.Printf("%s - %s", gaugeName, gaugeValue)
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(gaugeValue))
// 	})
// }
