package model

import (
	"math/rand/v2"
	"runtime"
)

type MetricCollection struct {
	GaugeMetrics map[string]uint64
	CountMetrics map[string]uint64
}

func NewMetricCollector() *MetricCollection {
	return &MetricCollection{
		GaugeMetrics: make(map[string]uint64),
		CountMetrics: make(map[string]uint64),
	}
}

func (m *MetricCollection) Collect() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	m.GaugeMetrics["Alloc"] = memStats.Alloc
	m.GaugeMetrics["BuckHashSys"] = memStats.BuckHashSys
	m.GaugeMetrics["Frees"] = memStats.Frees
	m.GaugeMetrics["GCCPUFraction"] = uint64(memStats.GCCPUFraction)
	m.GaugeMetrics["GCSys"] = memStats.GCSys
	m.GaugeMetrics["HeapAlloc"] = memStats.HeapAlloc
	m.GaugeMetrics["HeapIdle"] = memStats.HeapIdle
	m.GaugeMetrics["HeapInuse"] = memStats.HeapInuse
	m.GaugeMetrics["HeapObjects"] = memStats.HeapObjects
	m.GaugeMetrics["HeapReleased"] = memStats.HeapReleased
	m.GaugeMetrics["HeapSys"] = memStats.HeapSys
	m.GaugeMetrics["LastGC"] = memStats.LastGC
	m.GaugeMetrics["Lookups"] = memStats.Lookups
	m.GaugeMetrics["MCacheInuse"] = memStats.MCacheInuse
	m.GaugeMetrics["MCacheSys"] = memStats.MCacheSys
	m.GaugeMetrics["MSpanInuse"] = memStats.MSpanInuse
	m.GaugeMetrics["MSpanSys"] = memStats.MSpanSys
	m.GaugeMetrics["Mallocs"] = memStats.Mallocs
	m.GaugeMetrics["NextGC"] = memStats.NextGC
	m.GaugeMetrics["NumForcedGC"] = uint64(memStats.NumForcedGC)
	m.GaugeMetrics["NumGC"] = uint64(memStats.NumGC)
	m.GaugeMetrics["OtherSys"] = memStats.OtherSys
	m.GaugeMetrics["PauseTotalNs"] = memStats.PauseTotalNs
	m.GaugeMetrics["StackInuse"] = memStats.StackInuse
	m.GaugeMetrics["StackSys"] = memStats.StackSys
	m.GaugeMetrics["Sys"] = memStats.Sys
	m.GaugeMetrics["TotalAlloc"] = memStats.TotalAlloc
	m.GaugeMetrics["RandomValue"] = rand.Uint64()

	m.CountMetrics["PollCount"]++
}

func (m *MetricCollection) Clear() {
	for k := range m.CountMetrics {
		m.CountMetrics[k] = 0
	}

	for k := range m.GaugeMetrics {
		m.CountMetrics[k] = 0
	}
}
