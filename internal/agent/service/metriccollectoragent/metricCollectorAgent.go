package metriccollectoragent

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/model"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/repository"
)

const (
	sleepTime = 1
)

type metricCollectorAgent struct {
	pollInterval   int
	reportInterval int
	mc             model.MetricCollection
	repo           repository.MetricRepo
}

func NewMetricCollectorAgent(pollInterval, reportInterval int, repo repository.MetricRepo) *metricCollectorAgent {
	return &metricCollectorAgent{
		repo: repo,
		mc: model.MetricCollection{
			GaugeMetrics: make(map[string]uint64),
			CountMetrics: make(map[string]uint64),
		},
		pollInterval:   pollInterval,
		reportInterval: reportInterval,
	}
}

func (mca *metricCollectorAgent) EndlessCollectMetrics(c *resty.Client) error {
	i := 0
	for {
		if i != 0 && i%mca.pollInterval == 0 {
			mca.mc.Collect()
		}

		if i != 0 && i%mca.reportInterval == 0 {
			if err := mca.repo.StoreGaugeMetrics(mca.mc, c); err != nil {
				fmt.Println(err)
				return err
			}
			if err := mca.repo.StoreCounterMetrics(mca.mc, c); err != nil {
				fmt.Println(err)
				return err
			}
			mca.mc.Clear()
			i = 0
		}

		time.Sleep(time.Duration(sleepTime) * time.Second)
		i++
	}
}
