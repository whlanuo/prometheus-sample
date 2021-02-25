package main

import (
	"math"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

// metrics type: counter
var sampleCounter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "sample_counter",
	Help: "A sample counter metric"})

var sampleCounterVec = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "sample_counter_vec",
	Help: "A sample counter metric"},
	[]string{"label_name"})

// metrics type: gauge
var sampleGauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "sample_gauge",
	Help: "A sample gauge metric"})

var sampleGaugeVec = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "sample_gauge_vec",
	Help: "A sample gauge metric"},
	[]string{"label_name"})

// metrics type: histogram
var sampleHistogram = promauto.NewHistogram(prometheus.HistogramOpts{
	Name: "sample_histogram",
	Help: "A sample histogram metric",
	Buckets: []float64{0, 0.5, 1, 1.5, 2}})

var sampleHistogramVec = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "sample_histogram_vec",
	Help: "A sample histogram metric",
	Buckets: []float64{0, 0.5, 1, 1.5, 2}},
	[]string{"label_name"})

// metrics type: summary
var sampleSummary = promauto.NewSummary(prometheus.SummaryOpts{
	Name: "sample_summary",
	Help: "A sample summary metric",
	Objectives: map[float64]float64{0: 0, 0.5: 0, 0.75: 0, 0.9: 0, 0.99: 0, 1: 0}})

var sampleSummaryVec = promauto.NewSummaryVec(prometheus.SummaryOpts{
	Name: "sample_summary_vec",
	Help: "A sample summary metric",
	Objectives: map[float64]float64{0: 0, 0.5: 0, 0.75: 0, 0.9: 0, 0.99: 0, 1: 0}},
	[]string{"label_name"})

func main() {
	go func() {
		var x float64
		var step float64
		for {
			x = step * math.Pi
			step += 0.1

			sinValue := 1 + math.Sin(x)
			cosValue := 1 + math.Cos(x)
			log.Infof("sin_value: %v, cos_value: %v", sinValue, cosValue)

			// counter
			sampleCounter.Add(sinValue)
			sampleCounterVec.WithLabelValues("sin").Add(sinValue)
			sampleCounterVec.WithLabelValues("cos").Add(cosValue)

			// gauge
			sampleGauge.Set(sinValue)
			sampleGaugeVec.WithLabelValues("sin").Set(sinValue)
			sampleGaugeVec.WithLabelValues("cos").Set(cosValue)

			// histogram
			sampleHistogram.Observe(sinValue)
			sampleHistogramVec.WithLabelValues("sin").Observe(sinValue)
			sampleHistogramVec.WithLabelValues("cos").Observe(cosValue)

			// summary
			sampleSummary.Observe(sinValue)
			sampleSummaryVec.WithLabelValues("sin").Observe(sinValue)
			sampleSummaryVec.WithLabelValues("cos").Observe(cosValue)

			time.Sleep(10 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2000", nil)
}
