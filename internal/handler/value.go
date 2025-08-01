package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sky75444/go-practicum-sprint1-metrics/internal/logger"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/models"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/service"
)

type ValueHandler struct {
	updateMetricsService service.UpdateMetricsService
}

func NewValueHandler(umService service.UpdateMetricsService) *ValueHandler {
	return &ValueHandler{
		updateMetricsService: umService,
	}
}

func (u *ValueHandler) ValueHandle() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer logger.ZLog.Sync()
		sl := logger.ZLog.Sugar()

		if r.Header.Get("Content-Type") != "application/json" {
			sl.Errorw("Content-Type must be application/json")
			http.Error(w, "Content-Type must be application/json", http.StatusNotFound)
			return
		}

		var m models.Metrics

		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(&m); err != nil {
			sl.Errorw("cannot decode request JSON body", logger.ZError(err))
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		if m.MType == models.Counter {
			cVal, err := u.updateMetricsService.GetCounter(m.ID)
			if err != nil {
				sl.Errorw(fmt.Sprintf("metric not found - %s", m.ID), logger.ZError(err))
				http.Error(w, "metric not found", http.StatusNotFound)
				return
			}

			m.Delta = &cVal

			sl.Debugw("%s - %s", m.ID, cVal)
		}

		if m.MType == models.Gauge {
			gVal, err := u.updateMetricsService.GetGauge(m.ID)
			if err != nil {
				sl.Errorw(fmt.Sprintf("metric not found - %s", m.ID), logger.ZError(err))
				http.Error(w, "metric not found", http.StatusNotFound)
				return
			}

			m.Value = &gVal

			sl.Debugw("%s - %s", m.ID, gVal)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		enc := json.NewEncoder(w)
		if err := enc.Encode(m); err != nil {
			sl.Errorw("error encoding response", logger.ZError(err))
			return
		}
	})
}
