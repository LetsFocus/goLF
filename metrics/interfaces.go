package metrics

type CounterMetrics interface {
	Increment(labelValues ...string)
	Decrement(labelValues ...string)
	Add(count float64, labelValues ...string)
}

type HistogramMetrics interface {
	Observe(time float64, labelValues ...string)
}

type GaugeMetrics interface {
	Increment(labelValues ...string)
	Decrement(labelValues ...string)
	Subtract(value float64, labels ...string)
	Set(value float64, labels ...string)
	Add(value float64, labels ...string)
}
