package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sky75444/go-practicum-sprint1-metrics/internal/logger"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/models"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/service"
)

type UpdateHandler struct {
	updateMetricsService service.UpdateMetricsService
}

func NewUpdateHandler(umService service.UpdateMetricsService) *UpdateHandler {
	return &UpdateHandler{
		updateMetricsService: umService,
	}
}

func (u *UpdateHandler) UpdateHandle() http.HandlerFunc {
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

		var err error
		if strings.ToLower(m.MType) == models.Counter {
			sl.Debugw("UpdateCounter", m.ID, m.Delta)
			err = u.updateMetricsService.UpdateCounter(strings.ToLower(m.ID), *m.Delta)
		}
		if strings.ToLower(m.MType) == models.Gauge {
			sl.Debugw("UpdateGauge", m.ID, m.Value)
			err = u.updateMetricsService.UpdateGauge(strings.ToLower(m.ID), *m.Value)
		}

		if err != nil {
			sl.Errorw("internal server error while during update metric", logger.ZError(err))
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Metric updated"))
	})
}
