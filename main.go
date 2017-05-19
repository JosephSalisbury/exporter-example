package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	Count float64 = 0
)

type Exporter struct {
	count *prometheus.Desc
}

func NewExporter() *Exporter {
	return &Exporter{
		count: prometheus.NewDesc(
			prometheus.BuildFQName("exporter_example", "", "count"),
			"Total number of requests",
			nil,
			nil,
		),
	}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.count
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	Count++
	ch <- prometheus.MustNewConstMetric(e.count, prometheus.CounterValue, Count)
}

func main() {
	prometheus.MustRegister(NewExporter())

	http.Handle("/metrics", prometheus.Handler())
	http.ListenAndServe("localhost:8000", nil)
}
