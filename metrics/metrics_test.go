package metrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMetricsServer(t *testing.T) {
	metrics := NewMetricsServer()

	assert.NotNil(t, metrics.HttpRequestCount)
	assert.NotNil(t, metrics.HttpRequestDuration)
	assert.NotNil(t, metrics.GoGoroutines)
	assert.NotNil(t, metrics.GoMemoryUsage)
	assert.NotNil(t, metrics.HttpErrorCount)
	assert.NotNil(t, metrics.DBQueryDuration)
	assert.NotNil(t, metrics.CacheHits)
	assert.NotNil(t, metrics.CacheMisses)
	assert.NotNil(t, metrics.WebsocketConnections)
	assert.NotNil(t, metrics.DBConnectionPoolUsage)
	assert.NotNil(t, metrics.CpuUsage)
}
