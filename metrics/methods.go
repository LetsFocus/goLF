package metrics

func (c *CounterMetric) Increment(labels ...string) {
	if c.CounterVec != nil {
		c.CounterVec.WithLabelValues(labels...).Inc()

		return
	}

	c.Counter.Inc()
}

func (c *CounterMetric) Decrement(labels ...string) {
	if c.CounterVec != nil {
		c.CounterVec.WithLabelValues(labels...).Desc()

		return
	}

	c.Counter.Desc()
}

func (c *CounterMetric) Add(count float64, labels ...string) {
	if c.CounterVec != nil {
		c.CounterVec.WithLabelValues(labels...).Add(count)

		return
	}

	c.Counter.Add(count)
}

func (h *HistogramMetric) Observe(time float64, labels ...string) {
	if h.HistogramVec != nil {
		h.HistogramVec.WithLabelValues(labels...).Observe(time)

		return
	}

	h.Histogram.Observe(time)
}

func (g *GaugeMetric) Set(value float64, labels ...string) {
	if g.GaugeVec != nil {
		g.GaugeVec.WithLabelValues(labels...).Set(value)

		return
	}

	g.Gauge.Set(value)
}

func (g *GaugeMetric) Increment(labels ...string) {
	if g.GaugeVec != nil {
		g.GaugeVec.WithLabelValues(labels...).Inc()

		return
	}

	g.Gauge.Inc()
}

func (g *GaugeMetric) Decrement(labels ...string) {
	if g.GaugeVec != nil {
		g.GaugeVec.WithLabelValues(labels...).Desc()

		return
	}

	g.Gauge.Desc()
}

func (g *GaugeMetric) Subtract(value float64, labels ...string) {
	if g.GaugeVec != nil {
		g.GaugeVec.WithLabelValues(labels...).Sub(value)

		return
	}

	g.Gauge.Sub(value)
}

func (g *GaugeMetric) Add(value float64, labels ...string) {
	if g.GaugeVec != nil {
		g.GaugeVec.WithLabelValues(labels...).Add(value)

		return
	}

	g.Gauge.Add(value)
}
