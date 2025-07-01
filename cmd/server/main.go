package main

import (
	"fmt"
	"net/http"

	"github.com/sky75444/go-practicum-sprint1-metrics.git/cmd/server/di"
)

func main() {
	d := di.NewDI()
	d.Init()

	fmt.Println("Server started at http://localhost:8080")

	if err := http.ListenAndServe(`:8080`, d.Router.Mux); err != nil {
		panic(err)
	}
}
