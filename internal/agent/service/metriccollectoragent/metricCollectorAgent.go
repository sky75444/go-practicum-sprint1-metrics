package metriccollectoragent

import (
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/agent/alogger"
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
	defer alogger.AZLog.Sync()
	sl := alogger.AZLog.Sugar()

	i := 0
	for {
		if i != 0 && i%mca.pollInterval == 0 {
			mca.mc.Collect()
			sl.Debugw("metrics collected")
		}

		if i != 0 && i%mca.reportInterval == 0 {
			if err := mca.repo.StoreMetrics(mca.mc, c); err != nil {
				sl.Errorw("internal error", alogger.ZError(err))
				return err
			}

			sl.Debugw("metrics successfuly reported")

			mca.mc.Clear()
			i = 0
		}

		time.Sleep(time.Duration(sleepTime) * time.Second)
		i++
	}
}
