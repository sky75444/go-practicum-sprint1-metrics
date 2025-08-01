package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/logger"
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
		defer logger.ZLog.Sync()
		sl := logger.ZLog.Sugar()

		gaugeName := strings.ToLower(chi.URLParam(r, "gaugeName"))
		if gaugeName == "" {
			sl.Errorw("gauge name is missing")
			http.Error(w, "gauge name is missing", http.StatusNotFound)
			return
		}

		gaugeValueStr := strings.ToLower(chi.URLParam(r, "gaugeValue"))
		if gaugeValueStr == "" {
			sl.Errorw("gauge value is missing")
			http.Error(w, "gauge value is missing", http.StatusNotFound)
			return
		}

		value, err := strconv.ParseFloat(gaugeValueStr, 64)
		if err != nil {
			sl.Errorw("invalid gauge value")
			http.Error(w, "invalid gauge value", http.StatusBadRequest)
			return
		}

		if err := g.updateMetricsService.UpdateGauge(gaugeName, value); err != nil {
			sl.Errorw("invalid counter value", logger.ZError(err))
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		sl.Debugw("gauge metric updated - ", gaugeName)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Metric updated"))
	})
}
