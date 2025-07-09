package metriccollectoragent

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/model"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/repository"
)

type metricCollectorAgent struct {
	pollInterval   int
	reportInterval int
	mc             *model.MetricCollection
	repo           repository.MetricRepo
}

func NewMetricCollectorAgent(pollInterval, reportInterval int, repo repository.MetricRepo) *metricCollectorAgent {
	return &metricCollectorAgent{
		repo:           repo,
		mc:             model.NewMetricCollector(),
		pollInterval:   pollInterval,
		reportInterval: reportInterval,
	}
}

func (mca *metricCollectorAgent) EndlessCollectMetrics(c *resty.Client) error {
	i := 0
	for {
		if i == mca.reportInterval {
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

		if i%mca.pollInterval == 0 {
			mca.mc.Collect()
		}
		i += mca.pollInterval
		time.Sleep(time.Duration(mca.pollInterval) * time.Second)
	}
}
