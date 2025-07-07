package memstorage

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

type memStorage struct {
	mu       sync.RWMutex
	gauges   map[string]float64
	counters map[string]int64
}

func NewMemStorage() *memStorage {
	return &memStorage{
		gauges:   make(map[string]float64),
		counters: make(map[string]int64),
	}
}

func (m *memStorage) UpdateGauge(name string, value float64) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.gauges[name] = value
	return nil
}

func (m *memStorage) UpdateCounter(name string, value int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.counters[name] += value
	return nil
}

func (m *memStorage) GetCounter(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("counter name is empty")
	}

	v, exist := m.counters[name]
	if !exist {
		return "", fmt.Errorf("metric not found")
	}

	return fmt.Sprintf("%d", v), nil
}

func (m *memStorage) GetGauge(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("gauge name is empty")
	}

	v, exist := m.gauges[name]
	if !exist {
		return "", fmt.Errorf("metric not found")
	}

	return fmt.Sprintf("%f", v), nil
}

func (m *memStorage) GetAll() (string, error) {
	var metrics []string

	for k, v := range m.counters {
		metrics = append(metrics, fmt.Sprintf("%s - %d", k, v))
	}
	for k, v := range m.gauges {
		metrics = append(metrics, fmt.Sprintf("%s - %f", k, v))
	}

	sort.Strings(metrics)

	return strings.Join(metrics, "\n"), nil
}
