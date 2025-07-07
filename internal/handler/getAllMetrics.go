package handler

import (
	"log"
	"net/http"

	"github.com/sky75444/go-practicum-sprint1-metrics/internal/service"
)

type GetAllHandler struct {
	updateMetricsService service.UpdateMetricsService
}

func NewGetAllHandler(umService service.UpdateMetricsService) *GetAllHandler {
	return &GetAllHandler{
		updateMetricsService: umService,
	}
}

func (c *GetAllHandler) GetAll() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metricsList, err := c.updateMetricsService.GetAll()

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.Println("generated all metrics list")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(metricsList))
	})
}
