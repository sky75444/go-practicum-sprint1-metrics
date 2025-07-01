package main

import (
	"fmt"
	"net/http"

	"github.com/sky75444/go-practicum-sprint1-metrics.git/internal/handler"
	models "github.com/sky75444/go-practicum-sprint1-metrics.git/internal/model"
)

func main() {
	var storage = models.NewMemStorage()

	http.HandleFunc("/update/counter/", handler.UpdateCounterHandler(storage))
	http.HandleFunc("/update/gauge/", handler.UpdateGaugeHandler(storage))
	http.HandleFunc("/", handler.ErrorHandler)

	fmt.Println("Server started at http://localhost:8080")

	if err := http.ListenAndServe(`:8080`, nil); err != nil {
		panic(err)
	}
}
