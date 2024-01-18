package metrics

import (
	"net/http"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/cpu"
)

func NewMetricsServer() *Metrics {
	httpRequestCount := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)

	httpRequestDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Histogram of HTTP request durations",
		},
		[]string{"method", "endpoint", "status"},
	)

	goGoroutines := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "goroutines",
			Help: "Number of goroutines currently running",
		},
	)

	memoryUsage := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "my_app_memory_usage",
			Help: "Memory usage statistics.",
		},
		[]string{"type"}, // "alloc", "total_alloc", "heap_alloc".
	)

	httpErrorCount := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_errors_total",
			Help: "Total number of HTTP errors",
		},
		[]string{"method", "endpoint"},
	)

	dbQueryDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "db_query_duration_seconds",
			Help: "Histogram of database query durations",
		},
		[]string{"query_type"},
	)

	cacheHits := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "cache_hits_total",
			Help: "Total number of cache hits",
		},
		[]string{"cache_type"},
	)

	cacheMisses := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "cache_misses_total",
			Help: "Total number of cache misses",
		},
		[]string{"cache_type"},
	)

	websocketConnections := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "websocket_connections_total",
			Help: "Total number of WebSocket connections",
		},
	)

	dbConnectionPoolUsage := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "db_connection_pool_usage",
			Help: "Current usage of the database connection pool",
		},
	)

	cpuUsage := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "my_app_cpu_usage",
			Help: "CPU usage percentage.",
		},
		[]string{"mode"}, // "user" and "system" CPU usage
	)

	prometheus.MustRegister(dbConnectionPoolUsage, websocketConnections, cacheMisses, cacheHits,
		dbQueryDuration, httpErrorCount, memoryUsage, goGoroutines, httpRequestCount, httpRequestDuration, cpuUsage)

	metricsData := &Metrics{HttpRequestCount: httpRequestCount, HttpErrorCount: httpErrorCount,
		HttpRequestDuration: httpRequestDuration, GoGoroutines: goGoroutines, GoMemoryUsage: memoryUsage,
		CpuUsage: cpuUsage, DBConnectionPoolUsage: dbConnectionPoolUsage, DBQueryDuration: dbQueryDuration,
		CacheMisses: cacheMisses, CacheHits: cacheHits, WebsocketConnections: websocketConnections}

	go metricsData.collectSystemMetrics()

	return metricsData
}

// MetricsHandler exposes the metrics endpoint for Prometheus scraping.
func MetricsHandler() http.Handler {
	return promhttp.Handler()
}

func (m *Metrics) collectSystemMetrics() {
	var stats runtime.MemStats

	for {
		_ = make([]byte, 1<<20)

		runtime.ReadMemStats(&stats)

		m.GoMemoryUsage.WithLabelValues("alloc").Set(float64(stats.Alloc))
		m.GoMemoryUsage.WithLabelValues("total_alloc").Set(float64(stats.TotalAlloc))
		m.GoMemoryUsage.WithLabelValues("heap_alloc").Set(float64(stats.HeapAlloc))

		numGoroutines := runtime.NumGoroutine()
		m.GoGoroutines.Set(float64(numGoroutines))

		cpuPercentages, err := cpu.Percent(time.Second, false)
		if err == nil && len(cpuPercentages) > 0 {
			userCPU := cpuPercentages[0]
			m.CpuUsage.WithLabelValues("user").Set(userCPU)
		}

		// Sleep for a short duration to observe changes
		time.Sleep(time.Second)
	}
}
