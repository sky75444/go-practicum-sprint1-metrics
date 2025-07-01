package main

import (
	"fmt"
	"net/http"

	"github.com/sky75444/go-practicum-sprint1-metrics/internal/app"
)

func main() {
	d := app.NewDI()
	d.Init()

	fmt.Println("Server started at http://localhost:8080")

	if err := http.ListenAndServe(`:8080`, d.Router.Mux); err != nil {
		panic(err)
	}
}
