package main

import "github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/app"

func main() {
	di := app.NewDI()
	di.Init()

	if err := di.Services.MetricCollectorAgent.EndlessCollectMetrics(di.Client); err != nil {
		panic(err)
	}
}
