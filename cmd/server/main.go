package main

import (
	"fmt"
	"net/http"

	"github.com/sky75444/go-practicum-sprint1-metrics.git/internal/handler"
	"github.com/sky75444/go-practicum-sprint1-metrics.git/internal/models"
)

func main() {
	http.HandleFunc("/update/", handler.UpdateMetricsHandler)

	fmt.Println("Server started at http://localhost:8080")

	var storage = models.NewMemStorage()

	if err := http.ListenAndServe(`:8080`, nil); err != nil {
		panic(err)
	}
}
