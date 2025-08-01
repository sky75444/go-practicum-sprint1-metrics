package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/logger"
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
		defer logger.ZLog.Sync()
		sl := logger.ZLog.Sugar()

		counterName := strings.ToLower(chi.URLParam(r, "counterName"))
		if counterName == "" {
			sl.Errorw("counter name is missing")
			http.Error(w, "counter name is missing", http.StatusNotFound)
			return
		}

		counterValueStr := strings.ToLower(chi.URLParam(r, "counterValue"))
		if counterValueStr == "" {
			sl.Errorw("counter value is missing")
			http.Error(w, "counter value is missing", http.StatusNotFound)
			return
		}

		value, err := strconv.ParseInt(counterValueStr, 10, 64)
		if err != nil {
			sl.Errorw("invalid counter value")
			http.Error(w, "invalid counter value", http.StatusBadRequest)
			return
		}

		if err := c.updateMetricsService.UpdateCounter(counterName, value); err != nil {
			sl.Errorw("internal server error", logger.ZError(err))
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		sl.Debugw("counter metric updated - ", counterName)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Metric updated"))
	})
}
