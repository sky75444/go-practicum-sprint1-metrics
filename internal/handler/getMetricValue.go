package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/logger"
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
		defer logger.ZLog.Sync()
		sl := logger.ZLog.Sugar()

		metricType := strings.ToLower(chi.URLParam(r, "metricType"))
		if metricType == "" {
			sl.Errorw("metric type is missing")
			http.Error(w, "metric type is missing", http.StatusNotFound)
			return
		}

		metricName := strings.ToLower(chi.URLParam(r, "metricName"))
		if metricName == "" {
			sl.Errorw("metric name is missing")
			http.Error(w, "metric name is missing", http.StatusNotFound)
			return
		}

		var metricValStr string

		if metricType == "counter" {
			counterValue, err := gh.updateMetricsService.GetCounter(metricName)
			if err != nil {
				sl.Errorw("metric not found - ", metricName)
				http.Error(w, "metric not found", http.StatusNotFound)
				return
			}

			metricValStr = fmt.Sprintf("%d", counterValue)
		} else {
			gaugeValue, err := gh.updateMetricsService.GetGauge(metricName)
			if err != nil {
				sl.Errorw("metric not found - ", metricName)
				http.Error(w, "metric not found", http.StatusNotFound)
				return
			}

			metricValStr = strconv.FormatFloat(gaugeValue, 'f', -1, 64)
		}

		sl.Debugw(fmt.Sprintf("%s - %s", metricName, metricValStr))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(metricValStr))
	})
}

func (gh *GetHandler) GetAll() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer logger.ZLog.Sync()
		sl := logger.ZLog.Sugar()

		metricsList, err := gh.updateMetricsService.GetAll()
		if err != nil {
			sl.Errorw("internal server error", logger.ZError(err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		sl.Debugw("generated all metrics list")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(metricsList))
	})
}
