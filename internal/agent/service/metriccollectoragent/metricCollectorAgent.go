package metriccollectoragent

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/model"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/repository"
)

const (
	PollInterval = 2 * time.Second
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

func (mca *metricCollectorAgent) EndlessCollectMetrics() error {
	client := &http.Client{}

	i := 0
	for {
		if i == 10 {
			if err := mca.repo.StoreGaugeMetrics(*mca.mc, client); err != nil {
				fmt.Println(err)
				return err
			}
			if err := mca.repo.StoreCounterMetrics(*mca.mc, client); err != nil {
				fmt.Println(err)
				return err
			}
			i = 0
		}

		mca.mc.Collect()
		i += 2
		time.Sleep(PollInterval)
	}
}
