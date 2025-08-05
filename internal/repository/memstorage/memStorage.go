package memstorage

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/sky75444/go-practicum-sprint1-metrics/internal/logger"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/models"
)

type memStorage struct {
	mu            sync.RWMutex
	gauges        map[string]float64
	counters      map[string]int64
	fname         string
	storeInterval int
}

func NewMemStorage(fname string, needRestore bool, storeInterval int) (*memStorage, error) {
	mem := memStorage{
		gauges:        make(map[string]float64),
		counters:      make(map[string]int64),
		fname:         fname,
		storeInterval: storeInterval,
	}

	if !needRestore {
		return &mem, nil
	}

	if err := mem.initMetricsMapsFromFileData(); err != nil {
		return nil, err
	}

	return &mem, nil
}

func (m *memStorage) initMetricsMapsFromFileData() error {
	mm, err := m.loadMetricsFromFile()
	if err != nil {
		return err
	}

	for _, metric := range *mm {
		switch metric.MType {
		case models.Gauge:
			m.gauges[metric.ID] = *metric.Value
		default:
			m.counters[metric.ID] = *metric.Delta
		}
	}

	return nil
}

func (m *memStorage) UpdateGauge(name string, value float64) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.gauges[name] = float64(value)

	if m.storeInterval == 0 {
		go m.SaveDataToFile()
	}

	return nil
}

func (m *memStorage) UpdateCounter(name string, value int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.counters[name] += int64(value)

	if m.storeInterval == 0 {
		go m.SaveDataToFile()
	}
	return nil
}

func (m *memStorage) GetCounter(name string) (int64, error) {
	if name == "" {
		return 0, fmt.Errorf("counter name is empty")
	}

	v, exist := m.counters[name]
	if !exist {
		return 0, fmt.Errorf("metric not found")
	}

	return v, nil
}

func (m *memStorage) GetGauge(name string) (float64, error) {
	if name == "" {
		return 0, fmt.Errorf("gauge name is empty")
	}

	v, exist := m.gauges[name]
	if !exist {
		return 0, fmt.Errorf("metric not found")
	}

	return v, nil
}

func (m *memStorage) GetAll() (string, error) {
	var metrics []string

	for k, v := range m.counters {
		metrics = append(metrics, fmt.Sprintf("%s - %d", k, v))
	}
	for k, v := range m.gauges {
		metrics = append(metrics, fmt.Sprintf("%s - %.3f", k, v))
	}

	sort.Strings(metrics)

	return strings.Join(metrics, "\n"), nil
}

func (m *memStorage) StoreMetricsToFile(ctx context.Context) error {
	if m.storeInterval == 0 {
		return nil
	}

	defer logger.ZLog.Sync()
	sl := logger.ZLog.Sugar()

	defer m.SaveDataToFile()

	i := 0
	for {

		select {
		case <-ctx.Done():
			return nil
		default:
			if i == m.storeInterval {
				if err := m.SaveDataToFile(); err != nil {
					return err
				}

				sl.Debugw("metrics stored to file")
				i = 0
			}

			time.Sleep(time.Duration(1) * time.Second)
			i++
		}
	}
}

func (m *memStorage) SaveDataToFile() error {
	jData, err := json.Marshal(m.convertToModelMetrics())
	if err != nil {
		return err
	}

	formattedJ := formatJsonString(jData)
	os.WriteFile(m.fname, []byte(formattedJ), 0666)

	return nil
}

func formatJsonString(jsonData []byte) []byte {
	jString := string(jsonData)
	jString = strings.ReplaceAll(jString, "},{", "},\n{") // добавляем перенос строки между объектами
	jString = strings.ReplaceAll(jString, "{", "\t{")     // добавляем табуляцию строки перед объектом
	jString = strings.ReplaceAll(jString, "[", "[\n")
	jString = strings.ReplaceAll(jString, "]", "\n]")
	return []byte(jString)
}

func (m *memStorage) convertToModelMetrics() []models.Metrics {
	mm := []models.Metrics{}
	for mn, mv := range m.gauges {
		mmOne := models.Metrics{
			ID:    mn,
			MType: models.Gauge,
			Value: &mv,
		}

		mm = append(mm, mmOne)
	}
	for mn, mv := range m.counters {
		mmOne := models.Metrics{
			ID:    mn,
			MType: models.Counter,
			Delta: &mv,
		}

		mm = append(mm, mmOne)
	}

	return mm
}

func (m *memStorage) loadMetricsFromFile() (*[]models.Metrics, error) {
	data, err := os.ReadFile(m.fname)
	if err != nil {
		return nil, err
	}

	mm := []models.Metrics{}
	if len(data) == 0 {
		return &mm, nil
	}

	if err := json.Unmarshal(data, &mm); err != nil {
		return nil, err
	}

	return &mm, nil
}
