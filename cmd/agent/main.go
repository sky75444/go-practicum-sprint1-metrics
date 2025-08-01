package main

import (
	"github.com/sky75444/go-practicum-sprint1-metrics/cmd/agent/agentflags"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/agentconfig"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/alogger"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/app"
)

func main() {
	f := agentflags.NewParsedFlags()
	c := agentconfig.NewConfig(
		f.GetMemServerAddr(),
		f.GetLogLevel(),
		f.GetPollInterval(),
		f.GetReportInterval())

	di := app.NewDI()
	di.Init(c)

	if err := alogger.Init(c.LogLevel); err != nil {
		panic(err)
	}

	defer alogger.AZLog.Sync()
	sl := alogger.AZLog.Sugar()

	sl.Infow(
		"Starting agent",
		"server addr", c.MemStorageServerAddr,
		"pollInterval", c.PollInterval,
		"reportInterval", c.ReportInterval,
	)

	if err := di.Services.MetricCollectorAgent.EndlessCollectMetrics(di.Client); err != nil {
		sl.Fatalw(err.Error(), "event", "starting agent")
	}
}
