package app

import (
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/agentconfig"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/service"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/service/metriccollectoragent"
)

type services struct {
	MetricCollectorAgent service.MetricCollector
}

func NewServices(repos *repositories, config *agentconfig.Config) *services {
	return &services{
		MetricCollectorAgent: metriccollectoragent.NewMetricCollectorAgent(
			config.PollInterval,
			config.ReportInterval,
			repos.MetricStorage,
		),
	}
}
