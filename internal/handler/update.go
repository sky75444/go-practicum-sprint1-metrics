package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
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

		fmt.Println("q1")
		if r.Header.Get("Content-Type") != "application/json" {
			sl.Errorw("Content-Type must be application/json")
			http.Error(w, "Content-Type must be application/json", http.StatusNotFound)
			return
		}

		fmt.Println("q2")
		var m models.Metrics

		var buf bytes.Buffer
		if _, err := buf.ReadFrom(r.Body); err != nil {
			sl.Errorw("unmarshall error", logger.ZError(err))
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		if err := json.Unmarshal(buf.Bytes(), &m); err != nil {
			sl.Errorw("unmarshall error", logger.ZError(err))
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		// dec := json.NewDecoder(r.Body)
		// if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		// 	sl.Errorw("cannot decode request JSON body", logger.ZError(err))
		// 	http.Error(w, "internal server error", http.StatusInternalServerError)
		// 	return
		// }

		fmt.Println("q4")
		var err error
		if strings.ToLower(m.MType) == models.Counter {
			fmt.Println("q4.1")
			sl.Debugw("UpdateCounter", m.ID, m.Delta)
			err = u.updateMetricsService.UpdateCounter(strings.ToLower(m.ID), *m.Delta)
		}
		if strings.ToLower(m.MType) == models.Gauge {
			fmt.Println("q4.2")
			sl.Debugw("UpdateGauge", m.ID, m.Value)
			err = u.updateMetricsService.UpdateGauge(strings.ToLower(m.ID), *m.Value)
		}

		fmt.Println("q5")
		if err != nil {
			sl.Errorw("internal server error while during update metric", logger.ZError(err))
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		fmt.Println("q6")
		// w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Metric updated"))
	})
}
