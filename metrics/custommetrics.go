package metrics

import "github.com/prometheus/client_golang/prometheus"

func NewCounterCustomMetrics(name, description string, labels ...string) *CounterMetric {
	var counterMetrics CounterMetric

	if len(labels) == 0 {
		counterMetrics.Counter = prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: name,
				Help: description,
			},
		)
	}

	if len(labels) > 0 {
		counterMetrics.CounterVec = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: name,
				Help: description,
			},
			labels,
		)
	}

	return &counterMetrics
}

func NewGaugeCustomMetrics(name, description string, labels ...string) *GaugeMetric {
	var gaugeMetrics GaugeMetric

	if len(labels) == 0 {
		gaugeMetrics.Gauge = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: name,
				Help: description,
			},
		)
	}

	if len(labels) > 0 {
		gaugeMetrics.GaugeVec = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: name,
				Help: description,
			},
			labels,
		)
	}

	return &gaugeMetrics
}

func NewHistogramCustomMetrics(name, description string, labels ...string) *HistogramMetric {
	var HistogramMetrics HistogramMetric

	if len(labels) == 0 {
		HistogramMetrics.Histogram = prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Name:    name,
				Help:    description,
				Buckets: prometheus.DefBuckets,
			},
		)
	}

	if len(labels) > 0 {
		HistogramMetrics.HistogramVec = prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    name,
				Help:    description,
				Buckets: prometheus.DefBuckets,
			},
			labels,
		)
	}

	return &HistogramMetrics
}
