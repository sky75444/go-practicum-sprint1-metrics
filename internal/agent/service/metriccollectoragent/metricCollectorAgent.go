package metriccollectoragent

import (
	"fmt"
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

	j := 0
	for {

		if j == 10 {
			sl.Errorw("internal error", alogger.ZError(fmt.Errorf("server not allowed")))
			return fmt.Errorf("server not allowed")
		}

		sl.Infow("health checking server")
		if ok, _ := mca.repo.ServerHealthCheck(c); ok {
			sl.Infow("server is healthy")
			break
		}

		time.Sleep(time.Duration(2) * time.Second)

		j++
	}

	errChan := make(chan error)
	defer close(errChan)

	go func() {
		i := 0
		for {
			if i != 0 && i%mca.pollInterval == 0 {
				mca.mc.Collect()
				sl.Debugw("metrics collected")
			}

			if i != 0 && i%mca.reportInterval == 0 {
				if err := mca.repo.StoreMetrics(mca.mc, c); err != nil {
					sl.Errorw("internal error", alogger.ZError(err))
					errChan <- err
					return
				}

				sl.Debugw("metrics successfuly reported")

				mca.mc.Clear()
				i = 0
			}

			time.Sleep(time.Duration(sleepTime) * time.Second)
			i++
		}
	}()

	for err := range errChan {
		return err
	}

	return nil
}
