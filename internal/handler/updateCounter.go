package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
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

func (c *UpdateCounterHandler) CounterHandle() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		counterName := strings.ToLower(chi.URLParam(r, "counterName"))
		if counterName == "" {
			http.Error(w, "counter name is missing", http.StatusNotFound)
			return
		}

		counterValueStr := strings.ToLower(chi.URLParam(r, "counterValue"))
		if counterValueStr == "" {
			http.Error(w, "counter value is missing", http.StatusNotFound)
			return
		}

		value, err := strconv.ParseInt(counterValueStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid counter value", http.StatusBadRequest)
			return
		}

		if err := c.updateMetricsService.UpdateCounter(counterName, value); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		log.Println("counter metric updated - " + counterName)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Metric updated"))
	})
}

// func (c *UpdateCounterHandler) GetCounterHandle() http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		counterName := strings.ToLower(chi.URLParam(r, "counterName"))
// 		if counterName == "" {
// 			http.Error(w, "counter name is missing", http.StatusNotFound)
// 			return
// 		}

// 		counterValue, err := c.updateMetricsService.GetCounter(counterName)
// 		if err != nil {
// 			log.Println("metric not found" + " - " + counterName)
// 			http.Error(w, "metric not found", http.StatusNotFound)
// 			return
// 		}

// 		log.Printf("%s - %s", counterName, counterValue)
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(counterValue))
// 	})
// }
