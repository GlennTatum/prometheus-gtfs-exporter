package exporter

import (
	"log"
	"net/http"
	"sort"

	"github.com/GlennTatum/prometheus-gtfs-exporter/mta"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Exporter struct {
}

type delta_point struct {
	x1, x2 int
}

func Collect_departure_frequency(c *mta.Client, id string) int {
	t := c.StopDepartures(id)
	series := make([]int, 0)
	for _, data := range t {
		series = append(series, int(data.Departure_Time.Unix()))
	}
	if len(series)%2 != 0 { // make the series an even series as we are comparing the delta of pairs
		series = series[:len(series)-1]
	}
	// (0, 1) (1, 2) (2, 3) len = 4, 3 pairs
	// (0, 1) (1, 2) (2, 3) (3, 4) (4, 5) len = 6, 5 pairs
	genPairs := func(s []int) []delta_point {
		pairs := make([]delta_point, 0, len(s))
		x1 := 0
		x2 := 1
		for i := 0; i < len(s)-1; i++ {
			pairs = append(pairs, delta_point{x1, x2})
			x1++
			x2++
		}
		return pairs
	}

	sort.Ints(series) // times are in increasing order
	log.Println("Time Series Group ", series)

	points := genPairs(series)

	deltas := make([]int, len(points))
	for _, p := range points {
		deltas = append(deltas, series[p.x2]-series[p.x1])
	}
	sum := 0
	for _, v := range deltas {
		sum += v
	}
	avg := sum / len(deltas)
	log.Println("Departure Average Timing Frequency: ", avg)
	return avg

}

// TODO Setup exporter to describe the metrics being collected
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	labels := prometheus.Labels{}
	labels["TEST"] = "TEST_LABEL"
	desc := prometheus.NewDesc("test", "HELP: testing", []string{"test1"}, labels)
	ch <- desc
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {

	// TODO Swap to a Series of MTA Subway Routes rather than individual
	client := mta.NewClient()
	client.Get(mta.ACE)
	station := "A02S"

	avg := Collect_departure_frequency(client, station)

	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gtfs_avg_departure_frequency",
	})
	gauge.Set(float64(avg))

	ch <- gauge

}

func NewExporter() *Exporter {
	return &Exporter{}
}

func Exec() {

	exporter := NewExporter()

	prometheus.MustRegister(exporter) // Register the exporter to default collector

	log.Println("Prometheus GTFS Exporter Running!")
	log.Fatal(http.ListenAndServe("0.0.0.0:9091", promhttp.Handler())) // Star collections on DefaultGatherer
}
