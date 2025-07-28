package main

import (
	"fmt"

	"github.com/sky75444/go-practicum-sprint1-metrics/cmd/agent/agentflags"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/agentconfig"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/app"
)

func main() {
	f := agentflags.NewParsedFlags()
	c := agentconfig.NewConfig(
		f.GetMemServerAddr(),
		f.GetPollInterval(),
		f.GetReportInterval())

	di := app.NewDI()
	di.Init(c)

	fmt.Printf("Start endless collect metrics on: %s\n", c.MemStorageServerAddr)
	fmt.Printf("pollInterval: %d\n", c.PollInterval)
	fmt.Printf("reportInterval: %d\n", c.ReportInterval)
	if err := di.Services.MetricCollectorAgent.EndlessCollectMetrics(di.Client); err != nil {
		panic(err)
	}
}
