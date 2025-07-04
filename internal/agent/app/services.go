package app

import (
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/service"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/service/metriccollectoragent"
)

type services struct {
	MetricCollectorAgent service.MetricCollector
}

func NewServices(repos *repositories) *services {
	return &services{
		MetricCollectorAgent: metriccollectoragent.NewMetricCollectorAgent(repos.MetricStorage),
	}
}
