package main

import (
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/app"
)

func main() {
	d := app.NewDI()
	d.Init()
	d.Router.Start()
}
