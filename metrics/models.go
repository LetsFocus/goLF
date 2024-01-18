package metrics

import "github.com/prometheus/client_golang/prometheus"

type CounterMetric struct {
	Counter    prometheus.Counter
	CounterVec *prometheus.CounterVec
}

type GaugeMetric struct {
	Gauge    prometheus.Gauge
	GaugeVec *prometheus.GaugeVec
}

type HistogramMetric struct {
	Histogram    prometheus.Histogram
	HistogramVec *prometheus.HistogramVec
}

type Metrics struct {
	HttpRequestCount, HttpErrorCount, CacheHits, CacheMisses  *prometheus.CounterVec
	HttpRequestDuration, DBQueryDuration                      *prometheus.HistogramVec
	GoGoroutines, DBConnectionPoolUsage, WebsocketConnections prometheus.Gauge
	CpuUsage, GoMemoryUsage                                   *prometheus.GaugeVec
}
