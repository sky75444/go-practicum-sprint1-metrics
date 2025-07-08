package metriccollectoragent

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/model"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/repository"
)

const (
	PollInterval = 2
	SendInterval = 10
)

type metricCollectorAgent struct {
	mc   *model.MetricCollection
	repo repository.MetricRepo
}

func NewMetricCollectorAgent(repo repository.MetricRepo) *metricCollectorAgent {
	return &metricCollectorAgent{
		repo: repo,
		mc:   model.NewMetricCollector(),
	}
}

func (mca *metricCollectorAgent) EndlessCollectMetrics(c *resty.Client) error {
	i := 0
	for {
		if i == SendInterval {
			if err := mca.repo.StoreGaugeMetrics(*mca.mc, c); err != nil {
				fmt.Println(err)
				return err
			}
			if err := mca.repo.StoreCounterMetrics(*mca.mc, c); err != nil {
				fmt.Println(err)
				return err
			}
			mca.mc.Clear()
			i = 0
		}

		if i%2 == 0 {
			mca.mc.Collect()
		}
		i += PollInterval
		time.Sleep(PollInterval * time.Second)
	}
}
