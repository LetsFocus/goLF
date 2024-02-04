package metrics

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

func TestNewCounterCustomMetrics(t *testing.T) {
	testCases := []struct {
		name        string
		description string
		labels      []string
		expected    CounterMetric
	}{
		{
			name:        "test_counter",
			description: "Test Counter",
			labels:      nil,
			expected: CounterMetric{
				Counter:    prometheus.NewCounter(prometheus.CounterOpts{Name: "test_counter", Help: "Test Counter"}),
				CounterVec: nil,
			},
		},
		{
			name:        "test_counter_with_labels",
			description: "Test Counter With Labels",
			labels:      []string{"label1", "label2"},
			expected: CounterMetric{
				Counter: nil,
				CounterVec: prometheus.NewCounterVec(prometheus.CounterOpts{Name: "test_counter_with_labels",
					Help: "Test Counter With Labels"}, []string{"label1", "label2"}),
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := NewCounterCustomMetrics(testCase.name, testCase.description, testCase.labels...)

			assert.NotNil(t, testCase.expected, result)
		})
	}
}

func TestNewGaugeCustomMetrics(t *testing.T) {
	testCases := []struct {
		name        string
		description string
		labels      []string
		expected    GaugeMetric
	}{
		{
			name:        "test_gauge",
			description: "Test Gauge",
			labels:      nil,
			expected: GaugeMetric{
				Gauge:    prometheus.NewGauge(prometheus.GaugeOpts{Name: "test_gauge", Help: "Test Gauge"}),
				GaugeVec: nil,
			},
		},
		{
			name:        "test_gauge_with_labels",
			description: "Test Gauge With Labels",
			labels:      []string{"label1", "label2"},
			expected: GaugeMetric{
				Gauge: nil,
				GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "test_gauge_with_labels",
					Help: "Test Gauge With Labels"}, []string{"label1", "label2"}),
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := NewGaugeCustomMetrics(testCase.name, testCase.description, testCase.labels...)

			assert.NotNil(t, testCase.expected, result)
		})
	}
}

func TestNewHistogramCustomMetrics(t *testing.T) {
	testCases := []struct {
		name        string
		description string
		labels      []string
		expected    HistogramMetric
	}{
		{
			name:        "test_histogram",
			description: "Test Histogram",
			labels:      nil,
			expected: HistogramMetric{
				Histogram: prometheus.NewHistogram(prometheus.HistogramOpts{Name: "test_histogram",
					Help: "Test Histogram", Buckets: prometheus.DefBuckets}),
				HistogramVec: nil,
			},
		},
		{
			name:        "test_histogram_with_labels",
			description: "Test Histogram With Labels",
			labels:      []string{"label1", "label2"},
			expected: HistogramMetric{
				Histogram: nil,
				HistogramVec: prometheus.NewHistogramVec(prometheus.HistogramOpts{Name: "test_histogram_with_labels",
					Help: "Test Histogram With Labels", Buckets: prometheus.DefBuckets}, []string{"label1", "label2"}),
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := NewHistogramCustomMetrics(testCase.name, testCase.description, testCase.labels...)

			assert.NotNil(t, testCase.expected, result)
		})
	}
}
